package relaxting

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	authzConstants "github.com/chef/automate/components/authz-service/constants/v2"
	reportingapi "github.com/chef/automate/components/compliance-service/api/reporting"
	"github.com/chef/automate/components/compliance-service/ingest/ingestic/mappings"
	"github.com/chef/automate/components/compliance-service/inspec"
	"github.com/chef/automate/components/compliance-service/reporting"
	"github.com/chef/automate/components/compliance-service/reporting/util"
	"github.com/chef/automate/components/compliance-service/utils"
	"github.com/chef/automate/lib/errorutils"
	"github.com/chef/automate/lib/stringutils"
)

const MaxScrollRecordSize = 10000

func (backend ES2Backend) getDocIdHits(esIndex string,
	searchSource *elastic.SearchSource) ([]*elastic.SearchHit, time.Duration, error) {
	defer util.TimeTrack(time.Now(), "getDocIdHits")
	var startT = time.Now()

	hits := make([]*elastic.SearchHit, 0)

	client, err := backend.ES2Client()
	if err != nil {
		return hits, time.Since(startT), errors.Wrap(err, "getDocIdsHits cannot connect to elasticsearch")
	}

	source, err := searchSource.Source()
	if err != nil {
		return hits, time.Since(startT), errors.Wrap(err, "getDocIdHits unable to get Source")
	}
	LogQueryPartMin(esIndex, source, "getDocIdHits query searchSource")

	scroll := client.Scroll().
		Index(esIndex).
		SearchSource(searchSource)

	for {
		results, err := scroll.Do(context.Background())
		//LogQueryPartMin(results, "getDocIdHits query results")
		if err == io.EOF {
			return hits, time.Since(startT), nil // all results retrieved
		}
		if err != nil {
			return hits, time.Since(startT), errors.Wrap(err, "getDocIdHits unable to get results")
		}
		if results.TotalHits() > 0 && len(results.Hits.Hits) > 0 {
			for _, hit := range results.Hits.Hits {
				hits = append(hits, hit)
			}
		}
	}
}

func (backend ES2Backend) getNodeReportIdsFromTimeseries(esIndex string,
	filters map[string][]string,
	latestOnly bool) ([]string, error) {
	nodeReport := make(map[string]string, 0)
	boolQuery := backend.getFiltersQuery(filters, latestOnly)

	// aggs
	aggs := elastic.NewTermsAggregation().Field("node_uuid").Size(reporting.ESize).
		SubAggregation("distinct", elastic.NewTopHitsAggregation().Size(1).
			FetchSource(false).
			Sort("end_time", false))

	client, err := backend.ES2Client()
	if err != nil {
		return []string{}, errors.Wrap(err, "getNodeReportIdsFromTimeseries cannot connect to elasticsearch")
	}

	searchSource := elastic.NewSearchSource().
		Query(boolQuery).
		Aggregation("nodes", aggs).
		Size(0)

	source, err := searchSource.Source()
	if err != nil {
		return []string{}, errors.Wrap(err, "getNodeReportIdsFromTimeseries unable to get Source")
	}
	LogQueryPartMin(esIndex, source, "getNodeReportIdsFromTimeseries query searchSource")

	searchResult, err := client.Search().
		Index(esIndex).
		FilterPath(
			"took",
			"hits.total",
			"aggregations.nodes.buckets.key",
			"aggregations.nodes.buckets.distinct.hits.hits._id").
		SearchSource(searchSource).
		Do(context.Background())

	if err != nil {
		logrus.Errorf("unable to getNodeReportIdsFromTimeseries %v", err)
		return []string{}, errors.Wrap(err, "getNodeReportIdsFromTimeseries unable to complete search")
	}

	if searchResult.TotalHits() == 0 || searchResult.Hits.TotalHits == 0 {
		logrus.Debugf("getNodeReportIdsFromTimeseries: No report ids for the given filters: %+v\n", filters)
		// no matching report IDs is not an error, just return an empty array
		return []string{}, nil
	}

	outermostAgg, _ := searchResult.Aggregations.Terms("nodes")
	if outermostAgg != nil {
		for _, nodeBucket := range outermostAgg.Buckets {
			topHits, _ := nodeBucket.Aggregations.TopHits("distinct")
			nodeID := fmt.Sprintf("%s", nodeBucket.Key)
			for _, hit := range topHits.Hits.Hits {
				nodeReport[nodeID] = hit.Id
			}
		}
	}

	reportIds := MapValues(nodeReport)

	logrus.Debugf("getNodeReportIdsFromTimeseries returning %d report ids in %d milliseconds\n",
		len(reportIds), searchResult.TookInMillis)

	return reportIds, nil
}

func (backend ES2Backend) GetReportIds(esIndex string, filters map[string][]string) ([]string, error) {
	reportIds, err := backend.getNodeReportIdsFromTimeseries(esIndex, filters, true)
	if err != nil {
		return []string{}, errors.Wrap(err, "GetReportIds unable to get node report ids")
	}

	return reportIds, nil
}

func (backend ES2Backend) filterIdsByControl(esIndex string, ids, controls []string) ([]string, error) {
	idsQuery := elastic.NewIdsQuery(mappings.DocType)
	idsQuery.Ids(ids...)
	termsQuery := elastic.NewTermsQuery("profiles.controls.id", stringArrayToInterfaceArray(controls)...)
	reportIdsAndControlIdQuery := elastic.NewBoolQuery()

	reportIdsNestedQuery := elastic.NewNestedQuery("profiles.controls", reportIdsAndControlIdQuery.Must(idsQuery,
		termsQuery))

	fsc := elastic.NewFetchSourceContext(false).
		Include("took", "hits.total", "hits.hits._id")

	//we use scrolling to accomplish sizes greater than 10000 we set it to 10k here which gets us 10k/scroll.
	searchSource := elastic.NewSearchSource().
		Query(reportIdsNestedQuery).
		FetchSourceContext(fsc).
		Size(MaxScrollRecordSize)

	source, err := searchSource.Source()
	if err != nil {
		return ids, errors.Wrap(err, "filterIdsByControl unable to get Source")
	}
	LogQueryPartMin(esIndex, source, "filterIdsByControl Control filter query searchSource")

	logrus.Debugf("esIndex: %s", esIndex)
	hits, _, err := backend.getDocIdHits(esIndex, searchSource)
	if err != nil {
		return ids, errors.Wrap(err, "filterIdsByControl unable to get doc ids")
	}

	ids = make([]string, 0)

	// building up the new ids array from the controls filter query
	if len(hits) > 0 {
		for _, hit := range hits {
			ids = append(ids, hit.Id)
		}
	}
	return ids, nil
}

// GetAllReports returns all reports in a given timeframe
// TODO: support timeframe and pagination
func (backend *ES2Backend) GetReports(from int32, size int32, filters map[string][]string,
	sortField string, sortAsc bool) ([]*reportingapi.Report, int64, error) {
	myName := "GetReports"

	depth, err := backend.NewDepth(filters, true, false)
	if err != nil {
		return nil, 0, errors.Wrap(err, fmt.Sprintf("%s unable to get depth level for report", myName))
	}

	queryInfo := depth.getQueryInfo()

	logrus.Debugf("%s will retrieve all nodes", myName)

	client, err := backend.ES2Client()

	if err != nil {
		return nil, 0, errors.Wrapf(err, "%s, cannot connect to Elasticsearch", myName)
	}

	fsc := elastic.NewFetchSourceContext(true).Include(
		"node_uuid",
		"node_name",
		"environment",
		"end_time",
		"ipaddress")

	if queryInfo.level == ReportLevel {
		fsc.Include(
			"status",
			"controls_sums")
	}

	searchSource := elastic.NewSearchSource().
		FetchSourceContext(fsc).
		Query(queryInfo.filtQuery).
		Sort(sortField, sortAsc).
		From(int(from)).
		Size(int(size))

	source, err := searchSource.Source()
	if err != nil {
		return nil, 0, errors.Wrapf(err, "%s unable to get Source", myName)
	}
	LogQueryPartMin(queryInfo.esIndex, source, fmt.Sprintf("%s query searchSource", myName))

	searchResult, err := client.Search().
		SearchSource(searchSource).
		Index(queryInfo.esIndex).
		FilterPath(
			"took",
			"hits.total",
			"hits.hits._index",
			"hits.hits._id",
			"hits.hits._source",
			"hits.hits.inner_hits").
		Do(context.Background())

	if err != nil {
		return nil, 0, errors.Wrapf(err, "%s unable to complete search", myName)
	}
	logrus.Debugf("GetAllReports got %d reports in %d milliseconds\n", searchResult.TotalHits(),
		searchResult.TookInMillis)
	reports := make([]*reportingapi.Report, 0)
	if searchResult.TotalHits() > 0 && searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			item := ESInSpecSummary{}
			if hit.Source != nil {
				err := json.Unmarshal(*hit.Source, &item)
				if err == nil {
					t := item.EndTime.Round(1 * time.Second)
					timestamp, _ := ptypes.TimestampProto(t)
					report := reportingapi.Report{
						Id:        hit.Id,
						NodeId:    item.NodeID,
						NodeName:  item.NodeName,
						EndTime:   timestamp,
						Ipaddress: item.IPAddress,
					}

					var controlSummary reporting.NodeControlSummary
					var status string

					if queryInfo.level == ReportLevel {
						controlSummary = item.ControlsSums
						status = item.Status
					} else if queryInfo.level == ProfileLevel || queryInfo.level == ControlLevel {
						controlSummary, status, err = getDeepControlsSums(hit, queryInfo)
						if err != nil {
							//todo - handle this
							logrus.Errorf("%s time error: %s", myName, err.Error())
						}
					}
					report.Controls = convertToRSControlSummary(controlSummary)
					report.Status = status

					reports = append(reports, &report)
				} else {
					logrus.Errorf("GetAllReports unmarshal error: %s", err.Error())
				}
			}
		}
		return reports, searchResult.TotalHits(), nil
	}

	logrus.Debug("Found no reports\n")
	return reports, 0, nil
}

// GetReport returns the information about a single report
func (backend *ES2Backend) GetReport(esIndex string, reportId string,
	filters map[string][]string) (*reportingapi.Report, error) {
	myName := "GetReport"
	var report *reportingapi.Report

	depth, err := backend.NewDepth(filters, true, false)
	if err != nil {
		return report, errors.Wrapf(err, "%s unable to get depth level for report", myName)
	}

	queryInfo := depth.getQueryInfo()

	//normally, we compute the esIndex when we create a new Depth obj.. here we overwrite it what what's been passed in.
	queryInfo.esIndex = esIndex

	//normally, we compute the boolQuery when we create a new Depth obj.. here we, instead call a variation of the full
	// query builder. we need to do this because this variation of filter query provides this func with deeper
	// information about the report being retrieved.
	queryInfo.filtQuery = backend.getFiltersQueryForDeepReport(reportId, filters)

	logrus.Debugf("%s will retrieve all nodes", myName)

	fsc := elastic.NewFetchSourceContext(true)

	if queryInfo.level != ReportLevel {
		fsc.Exclude("profiles")
	}
	logrus.Debugf("GetReport for reportid=%s, filters=%+v", reportId, filters)

	searchSource := elastic.NewSearchSource().
		FetchSourceContext(fsc).
		Query(queryInfo.filtQuery).
		Size(1)

	source, err := searchSource.Source()
	if err != nil {
		return report, errors.Wrapf(err, "%s unable to get Source", myName)
	}
	LogQueryPartMin(queryInfo.esIndex, source, fmt.Sprintf("%s query searchSource", myName))

	searchResult, err := queryInfo.client.Search().
		SearchSource(searchSource).
		Index(queryInfo.esIndex).
		FilterPath(
			"took",
			"hits.total",
			"hits.hits._id",
			"hits.hits._source",
			"hits.hits.inner_hits").
		Do(context.Background())

	if err != nil {
		return report, errors.Wrapf(err, "%s unable to complete search", myName)
	}

	logrus.Debugf("%s got %d reports in %d milliseconds\n", myName, searchResult.TotalHits(),
		searchResult.TookInMillis)

	// we should only receive one value
	if searchResult.TotalHits() > 0 && searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			esInSpecReport := ESInSpecReport{}
			if hit.Source != nil {
				err := json.Unmarshal(*hit.Source, &esInSpecReport)
				if err == nil {
					//TODO: FIX Unmarshal error(json: cannot unmarshal array into Go struct field
					// ESInSpecReportControl.results) and move the read all profiles section here
				} else {
					logrus.Errorf("%s unmarshal error: %s", myName, err.Error())
				}

				var esInspecProfiles []ESInSpecReportProfile //`json:"profiles"`
				var status string

				if queryInfo.level == ReportLevel {
					esInspecProfiles = esInSpecReport.Profiles
					status = esInSpecReport.Status
				} else if queryInfo.level == ProfileLevel || queryInfo.level == ControlLevel {
					esInspecProfiles, status, err = getDeepInspecProfiles(hit, queryInfo)
					if err != nil {
						//todo - handle this
						logrus.Errorf("%s time error: %s", myName, err.Error())
					}
				}
				esInSpecReport.Status = status

				// read all profiles
				profiles := make([]*reportingapi.Profile, 0)
				for _, esInSpecReportProfileMin := range esInspecProfiles {
					logrus.Debugf("Determine profile: %s", esInSpecReportProfileMin.Name)
					esInSpecProfile, err := backend.GetProfile(esInSpecReportProfileMin.SHA256)
					if err != nil {
						logrus.Errorf("%s - Could not get profile: %s", myName, err.Error())
					}

					reportProfile := inspec.Profile{}
					reportProfile.Name = esInSpecProfile.Name
					reportProfile.CopyrightEmail = esInSpecProfile.CopyrightEmail
					reportProfile.Copyright = esInSpecProfile.Copyright
					reportProfile.Title = esInSpecProfile.Title
					reportProfile.Summary = esInSpecProfile.Summary
					reportProfile.Version = esInSpecProfile.Version
					reportProfile.Sha256 = esInSpecProfile.Sha256
					reportProfile.Status = esInSpecProfile.Status
					reportProfile.SkipMessage = esInSpecProfile.SkipMessage
					reportProfile.Status = esInSpecReportProfileMin.Status
					reportProfile.SkipMessage = esInSpecReportProfileMin.SkipMessage
					reportProfile.Full = esInSpecReportProfileMin.Full

					dependsHash := make(map[string]*ESInSpecReportDepends, len(esInSpecReportProfileMin.Depends))
					for _, esInSpecProfileDependency := range esInSpecReportProfileMin.Depends {
						depCopy := esInSpecProfileDependency // to prevent https://github.com/golang/go/issues/20725
						dependsHash[esInSpecProfileDependency.Name] = &depCopy
					}
					// Picking the report specific dependency info(status, skip_message) and adding it to the static
					// info retrieved from comp-profiles
					for _, esInSpecProfileDep := range esInSpecProfile.Depends {
						if hash, ok := dependsHash[esInSpecProfileDep.Name]; ok {
							esInSpecProfileDep.Status = hash.Status
							esInSpecProfileDep.SkipMessage = hash.SkipMessage
						}
					}
					reportProfile.Dependencies = esInSpecProfile.Depends

					reportProfile.Controls = make([]inspec.Control, len(esInSpecReportProfileMin.Controls))
					// Creating a map of report control ids to avoid a n^2 complexity later on when we look for the
					// matching profile control id
					profileControlsMap := make(map[string]*reportingapi.Control, len(esInSpecProfile.Controls))
					for _, control := range esInSpecProfile.Controls {
						profileControlsMap[control.Id] = control
					}

					convertedControls := make([]*reportingapi.Control, 0)
					// Enrich min report controls with profile metadata
					for _, reportControlMin := range esInSpecReportProfileMin.Controls {
						// store controls to returned report
						convertedControl := convertControl(profileControlsMap, reportControlMin, filters)
						if convertedControl != nil {
							convertedControls = append(convertedControls, convertedControl)
						}
					}

					if len(convertedControls) > 1 {
						// Sort convertedControls by Id
						sort.Slice(convertedControls, func(i, j int) bool {
							return convertedControls[i].Id < convertedControls[j].Id
						})
					}

					// TODO: fix this (vj)
					// Name: control.Attribute
					convertedAttributes := []*reportingapi.Attribute{}

					convertedProfile := reportingapi.Profile{
						Name:           reportProfile.Name,
						Full:           reportProfile.Full,
						Title:          reportProfile.Title,
						Maintainer:     reportProfile.Maintainer,
						Copyright:      reportProfile.Copyright,
						CopyrightEmail: reportProfile.CopyrightEmail,
						License:        reportProfile.License,
						Summary:        reportProfile.Summary,
						Version:        reportProfile.Version,
						Supports:       convertMapSupportsToRSSupports(reportProfile.Supports),
						Depends:        reportProfile.Dependencies,
						Sha256:         reportProfile.Sha256,
						Groups:         reportProfile.Groups,
						Controls:       convertedControls,
						Attributes:     convertedAttributes,
						Status:         reportProfile.Status,
						SkipMessage:    reportProfile.SkipMessage,
					}
					profiles = append(profiles, &convertedProfile)
				}
				ipAddress := ""
				if esInSpecReport.IPAddress != nil {
					ipAddress = *esInSpecReport.IPAddress
				}
				timestamp, _ := ptypes.TimestampProto(esInSpecReport.EndTime)
				report = &reportingapi.Report{
					Id:          hit.Id,
					NodeId:      esInSpecReport.NodeID,
					NodeName:    esInSpecReport.NodeName,
					Environment: esInSpecReport.Environment,
					Status:      esInSpecReport.Status,
					EndTime:     timestamp,
					Version:     esInSpecReport.InSpecVersion,
					Profiles:    profiles,
					Ipaddress:   ipAddress,
					Fqdn:        esInSpecReport.FQDN,
				}
				report.Statistics = &reportingapi.Statistics{
					Duration: esInSpecReport.Statistics.Duration,
				}
				report.Platform = &reportingapi.Platform{
					Name:    esInSpecReport.Platform.Name,
					Release: esInSpecReport.Platform.Release,
					Full:    esInSpecReport.Platform.Full,
				}
			}
		}
		return report, nil
	}

	return report, errorutils.ProcessNotFound(nil, reportId)
}

func convertControl(profileControlsMap map[string]*reportingapi.Control, reportControlMin ESInSpecReportControl, filters map[string][]string) *reportingapi.Control {
	profileControl := profileControlsMap[reportControlMin.ID]

	if profileControl == nil {
		return nil
	}

	profileControl.Results = make([]*reportingapi.Result, 0)
	if reportControlMin.ID != profileControl.Id {
		return nil
	}

	minResults := make([]*reportingapi.Result, len(reportControlMin.Results))
	for i, result := range reportControlMin.Results {
		minResults[i] = &reportingapi.Result{
			Status:      result.Status,
			CodeDesc:    result.CodeDesc,
			RunTime:     result.RunTime,
			Message:     result.Message,
			SkipMessage: result.SkipMessage,
		}
	}

	convertedControl := reportingapi.Control{
		Id:             profileControl.Id,
		Code:           profileControl.Code,
		Desc:           profileControl.Desc,
		Impact:         profileControl.Impact,
		Title:          profileControl.Title,
		SourceLocation: profileControl.SourceLocation,
		Results:        minResults,
	}

	jsonTags := make(map[string]*reportingapi.TagValues, 0)
	for _, tag := range reportControlMin.StringTags {
		if len(tag.Values) == 0 {
			jsonTags[tag.Key] = &reportingapi.TagValues{Values: []string{"null"}}
		}
		vals := make([]string, 0)
		for _, val := range tag.Values {
			vals = append(vals, val)
			jsonTags[tag.Key] = &reportingapi.TagValues{Values: vals}
		}
	}

	if !doesControlTagMatchFilter(filters, jsonTags) {
		return nil
	}

	convertedControl.StringTags = jsonTags
	var jsonRefs []*reportingapi.Ref
	refs, _ := json.Marshal(profileControl.Refs)
	json.Unmarshal(refs, &jsonRefs) // nolint: errcheck
	convertedControl.Refs = jsonRefs
	return &convertedControl

}

func doesControlTagMatchFilter(filters map[string][]string,
	jsonTags map[string]*reportingapi.TagValues) bool {
	matchingTagRequired := false
	for filterKey, filterVals := range filters {
		if strings.HasPrefix(filterKey, "control_tag") {
			matchingTagRequired = true
			trimmed := strings.TrimPrefix(filterKey, "control_tag:")
			if tagVal, ok := jsonTags[trimmed]; ok {
				for _, val := range filterVals {
					if contains(tagVal.Values, val) || val == "null" {
						return true
					}
				}
			}

		}
	}
	if matchingTagRequired {
		return false
	}
	return true
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if strings.HasSuffix(x, n) {
			return true
		}
	}
	return false
}

func (backend *ES2Backend) GetControlListItems(ctx context.Context, filters map[string][]string,
	size int32) (*reportingapi.ControlItems, error) {
	myName := "GetControlListItems"

	contListItems := make([]*reportingapi.ControlItem, 0)

	client, err := backend.ES2Client()
	if err != nil {
		logrus.Errorf("Cannot connect to ElasticSearch: %s", err)
		return nil, err
	}

	// we need to use start_time or else we won't see all of the controls that were used outside of end_time day
	// we use start_time for the same reason we need it for trends
	esIndex, err := GetEsIndex(filters, false, false)
	if err != nil {
		return nil, errors.Wrap(err, myName)
	}

	//here, we set latestOnly to true.  We may need to set it to false if we want to search non lastest reports
	//for now, we don't search non-latest reports so don't do it.. it's slower for obvious reasons.
	filtQuery := backend.getFiltersQuery(filters, true)

	searchSource := elastic.NewSearchSource().
		Query(filtQuery).
		FetchSource(false).
		Size(1)

	controlTermsAgg := elastic.NewTermsAggregation().Field("profiles.controls.id").
		Size(int(size)).
		Order("_key", true)

	controlTermsAgg.SubAggregation("impact",
		elastic.NewTermsAggregation().Field("profiles.controls.impact").Size(1))

	controlTermsAgg.SubAggregation("title",
		elastic.NewTermsAggregation().Field("profiles.controls.title").Size(1))

	controlTermsAgg.SubAggregation("skipped",
		elastic.NewFilterAggregation().Filter(elastic.NewTermQuery("profiles.controls.status", "skipped")))

	controlTermsAgg.SubAggregation("failed",
		elastic.NewFilterAggregation().Filter(elastic.NewTermQuery("profiles.controls.status", "failed")))

	controlTermsAgg.SubAggregation("passed",
		elastic.NewFilterAggregation().Filter(elastic.NewTermQuery("profiles.controls.status", "passed")))

	controlTermsAgg.SubAggregation("end_time", elastic.NewReverseNestedAggregation().SubAggregation("most_recent_report",
		elastic.NewTermsAggregation().Field("end_time").Size(1)))

	profileInfoAgg := elastic.NewReverseNestedAggregation().Path("profiles").SubAggregation("version",
		elastic.NewTermsAggregation().Field("profiles.name"))

	profileInfoAgg.SubAggregation("sha",
		elastic.NewTermsAggregation().Field("profiles.sha256"))

	profileInfoAgg.SubAggregation("title",
		elastic.NewTermsAggregation().Field("profiles.title"))

	profileInfoAgg.SubAggregation("version",
		elastic.NewTermsAggregation().Field("profiles.version"))

	controlTermsAgg.SubAggregation("profile", profileInfoAgg)

	controlsQuery := &elastic.BoolQuery{}
	// Going through all filters to find the ones prefixed with 'control_tag', e.g. 'control_tag:nist'
	for filterType := range filters {
		if strings.HasPrefix(filterType, "control_tag:") {
			_, tagKey := leftSplit(filterType, ":")
			termQuery := backend.newNestedTermQueryFromControlTagsFilter(tagKey, filters[filterType])
			controlsQuery = controlsQuery.Must(termQuery)
		}
	}
	if len(filters["control_status"]) > 0 {
		controlStatusQuery := backend.newTermQueryFromFilter("profiles.controls.status", filters["control_status"])
		controlsQuery = controlsQuery.Must(controlStatusQuery)
		logrus.Infof("controls status query %v", controlStatusQuery)
	}
	if len(filters["control_name"]) > 0 {
		controlTitlesQuery := backend.newTermQueryFromFilter("profiles.controls.title.lower", filters["control_name"])
		controlsQuery = controlsQuery.Should(controlTitlesQuery)
		logrus.Infof("controls name query %v", controlTitlesQuery)
	}
	if len(filters["control"]) > 0 {
		controlIdsQuery := backend.newTermQueryFromFilter("profiles.controls.id", filters["control"])
		controlsQuery = controlsQuery.Should(controlIdsQuery)
		logrus.Infof("controls ids query %v", controlIdsQuery)
	}

	filteredControls := elastic.NewFilterAggregation().Filter(controlsQuery)
	filteredControls.SubAggregation("control", controlTermsAgg)
	controlsAgg := elastic.NewNestedAggregation().Path("profiles.controls")
	controlsAgg.SubAggregation("filtered_controls", filteredControls)

	profilesQuery := &elastic.BoolQuery{}
	if len(filters["profile_name"]) > 0 {
		profileTitlesQuery := backend.newTermQueryFromFilter("profiles.title.lower", filters["profile_name"])
		profilesQuery = profilesQuery.Should(profileTitlesQuery)
		logrus.Infof("profiles name query %v", profileTitlesQuery)
	}

	if len(filters["profile_id"]) > 0 {
		profilesShaQuery := backend.newTermQueryFromFilter("profiles.sha256", filters["profile_id"])
		profilesQuery = profilesQuery.Should(profilesShaQuery)
	}

	filteredProfiles := elastic.NewFilterAggregation().Filter(profilesQuery)
	filteredProfiles.SubAggregation("controls", controlsAgg)
	outterMostProfilesAgg := elastic.NewNestedAggregation().Path("profiles")
	outterMostProfilesAgg.SubAggregation("filtered_profiles", filteredProfiles)

	searchSource.Aggregation("profiles", outterMostProfilesAgg)

	source, err := searchSource.Source()
	if err != nil {
		return nil, errors.Wrapf(err, "%s unable to get Source", myName)
	}
	LogQueryPartMin(esIndex, source, fmt.Sprintf("%s query", myName))

	searchResult, err := client.Search().
		SearchSource(searchSource).
		Index(esIndex).
		Do(ctx)

	if err != nil {
		logrus.Errorf("%s search failed", myName)
		return nil, err
	}

	logrus.Debugf("Search query took %d milliseconds\n", searchResult.TookInMillis)

	if outerProfilesAggResult, found := searchResult.Aggregations.Nested("profiles"); found {
		if outerFilteredProfiles, found := outerProfilesAggResult.Aggregations.Filter("filtered_profiles"); found {
			if outerControlsAggResult, found := outerFilteredProfiles.Aggregations.Nested("controls"); found {
				if filteredControls, found := outerControlsAggResult.Aggregations.Filter("filtered_controls"); found {
					if controlBuckets, found := filteredControls.Aggregations.Terms("control"); found && len(controlBuckets.Buckets) > 0 {
						for _, controlBucket := range controlBuckets.Buckets {
							contListItem, err := backend.getControlItem(controlBucket)
							if err != nil {
								return nil, err
							}
							contListItems = append(contListItems, &contListItem)
						}
					}
				}
			}
		}
	}

	contListItemList := &reportingapi.ControlItems{ControlItems: contListItems}
	return contListItemList, nil
}

func (backend *ES2Backend) getControlItem(controlBucket *elastic.AggregationBucketKeyItem) (reportingapi.ControlItem, error) {
	contListItem := reportingapi.ControlItem{}
	id, ok := controlBucket.Key.(string)
	if !ok {
		logrus.Errorf("could not convert the value of controlBucket: %v, to a string!", controlBucket)
	}
	contListItem.Id = id
	controlSummary := &reportingapi.ControlSummary{
		Passed:  &reportingapi.Total{},
		Skipped: &reportingapi.Total{},
		Failed:  &reportingapi.Failed{},
	}
	if aggResult, found := controlBucket.Aggregations.Terms("title"); found {
		//there can only be one
		bucket := aggResult.Buckets[0]
		title, ok := bucket.Key.(string)
		if !ok {
			logrus.Errorf("could not convert the value of title: %v, to a string!", bucket)
		}
		contListItem.Title = title
	}
	if impactAggResult, found := controlBucket.Aggregations.Terms("impact"); found {
		//there can only be one
		impactBucket := impactAggResult.Buckets[0]
		impactAsNumber, ok := impactBucket.Key.(float64)
		if !ok {
			logrus.Errorf("could not convert the value of impact: %v, to a float!", impactBucket)
		}
		contListItem.Impact = float32(impactAsNumber)
		controlSummary.Total = int32(impactBucket.DocCount)
	}
	profileMin := &reportingapi.ProfileMin{}
	if profileResult, found := controlBucket.Aggregations.ReverseNested("profile"); found {
		if result, found := profileResult.Terms("sha"); found && len(result.Buckets) > 0 {
			sha := result.Buckets[0]
			name := sha.KeyNumber
			profileMin.Id = string(name)
		}
		if result, found := profileResult.Terms("title"); found && len(result.Buckets) > 0 {
			title := result.Buckets[0]
			name := title.KeyNumber
			profileMin.Title = string(name)
		}
		if result, found := profileResult.Terms("version"); found && len(result.Buckets) > 0 {
			version := result.Buckets[0]
			name := version.KeyNumber
			profileMin.Version = string(name)
		}
	}
	contListItem.Profile = profileMin

	if endTimeResult, found := controlBucket.Aggregations.ReverseNested("end_time"); found {
		if result, found := endTimeResult.Terms("most_recent_report"); found && len(result.Buckets) > 0 {
			endTime := result.Buckets[0]
			name := endTime.KeyAsString

			endTimeAsTime, err := time.Parse(time.RFC3339, *name)
			if err != nil {
				return reportingapi.ControlItem{}, errors.Wrapf(err, "%s time error: ", *name)
			}

			timestamp, err := ptypes.TimestampProto(endTimeAsTime)
			if err != nil {
				return reportingapi.ControlItem{}, errors.Wrapf(err, "%s time error: ", *name)
			} else {
				contListItem.EndTime = timestamp
			}
		}
	}
	if passed, found := controlBucket.Aggregations.Filter("passed"); found {
		controlSummary.Passed.Total = int32(passed.DocCount)
	}
	if skipped, found := controlBucket.Aggregations.Filter("skipped"); found {
		controlSummary.Skipped.Total = int32(skipped.DocCount)
	}
	if failed, found := controlBucket.Aggregations.Filter("failed"); found {
		total := int32(failed.DocCount)
		controlSummary.Failed.Total = total
		if contListItem.Impact < 0.4 {
			controlSummary.Failed.Minor = total
		} else if contListItem.Impact < 0.7 {
			controlSummary.Failed.Major = total
		} else {
			controlSummary.Failed.Critical = total
		}
	}
	contListItem.ControlSummary = controlSummary
	return contListItem, nil
}

//getFiltersQuery - builds up an elasticsearch query filter based on the filters map that is passed in
//  arguments: filters - is a map of filters that serve as the source for generated es query filters
//             latestOnly - specifies whether or not we are only interested in retrieving only the latest report
//  return *elastic.BoolQuery
func (backend ES2Backend) getFiltersQuery(filters map[string][]string, latestOnly bool) *elastic.BoolQuery {
	utils.DeDupFilters(filters)

	typeQuery := elastic.NewTypeQuery(mappings.DocType)

	boolQuery := elastic.NewBoolQuery()
	boolQuery = boolQuery.Must(typeQuery)

	// These are filter types where we use ElasticSearch Term Queries
	filterTypes := []string{"environment", "organization", "chef_server", "chef_tags",
		"policy_group", "policy_name", "status", "node_name", "platform", "platform_with_version",
		"role", "recipe", "inspec_version", "ipaddress"}

	for _, filterType := range filterTypes {
		if len(filters[filterType]) > 0 {
			ESFieldName := backend.getESFieldName(filterType)
			termQuery := backend.newTermQueryFromFilter(ESFieldName, filters[filterType])
			boolQuery = boolQuery.Must(termQuery)
		}
	}

	if len(filters["control_name"]) > 0 {
		termQuery := backend.newNestedTermQueryFromFilter("profiles.controls.title.lower", "profiles.controls",
			filters["control_name"])
		boolQuery = boolQuery.Must(termQuery)
	}

	if len(filters["profile_name"]) > 0 {
		termQuery := backend.newNestedTermQueryFromFilter("profiles.title.lower", "profiles", filters["profile_name"])
		boolQuery = boolQuery.Must(termQuery)
	}

	if len(filters["profile_with_version"]) > 0 {
		termQuery := backend.newNestedTermQueryFromFilter("profiles.full.lower", "profiles", filters["profile_with_version"])
		boolQuery = boolQuery.Must(termQuery)
	}

	if len(filters["node_id"]) > 0 {
		termQuery := elastic.NewTermsQuery("node_uuid", stringArrayToInterfaceArray(filters["node_id"])...)
		boolQuery = boolQuery.Must(termQuery)
	}

	numberOfProfiles := len(filters["profile_id"])
	numberOfControls := len(filters["control"])
	if numberOfProfiles > 0 || numberOfControls > 0 {
		profileBaseFscIncludes := []string{"profiles.name", "profiles.sha256", "profiles.version"}
		profileLevelFscIncludes := []string{"profiles.controls_sums", "profiles.status"}
		controlLevelFscIncludes := []string{"profiles.controls.id", "profiles.controls.status",
			"profiles.controls.impact"}

		profileAndControlQuery := getProfileAndControlQuery(filters, profileBaseFscIncludes,
			profileLevelFscIncludes, controlLevelFscIncludes)
		boolQuery = boolQuery.Must(profileAndControlQuery)
	}

	if len(filters["start_time"]) > 0 || len(filters["end_time"]) > 0 {
		endTime := firstOrEmpty(filters["end_time"])
		startTime := firstOrEmpty(filters["start_time"])

		timeRangeQuery := elastic.NewRangeQuery("end_time")
		if len(startTime) > 0 {
			timeRangeQuery.Gte(startTime)
		}
		if len(endTime) > 0 {
			timeRangeQuery.Lte(endTime)
		}

		boolQuery = boolQuery.Must(timeRangeQuery)
	}

	if len(filters["projects"]) > 0 {
		projectsQuery := elastic.NewBoolQuery()

		if stringutils.SliceContains(filters["projects"], authzConstants.UnassignedProjectID) {
			emptyProjectQuery := elastic.NewBoolQuery()
			emptyProjectQuery.MustNot(elastic.NewExistsQuery("projects"))
			projectsQuery.Should(emptyProjectQuery)
		}

		assignedProjectIds := stringutils.SliceFilter(filters["projects"], func(projectId string) bool {
			return projectId != authzConstants.UnassignedProjectID
		})

		if len(assignedProjectIds) > 0 {
			projectMatchQuery := elastic.NewTermsQuery("projects", stringArrayToInterfaceArray(assignedProjectIds)...)
			projectsQuery.Should(projectMatchQuery)
		}

		boolQuery = boolQuery.Filter(projectsQuery)
	}

	if len(filters["job_id"]) > 0 {
		termQuery := elastic.NewTermsQuery("job_uuid", stringArrayToInterfaceArray(filters["job_id"])...)
		boolQuery = boolQuery.Must(termQuery)
	} else if latestOnly {
		//only if there is no job_id filter set, do we want the daily latest
		termQuery := elastic.NewTermsQuery("daily_latest", true)
		boolQuery = boolQuery.Must(termQuery)
	}

	// Going through all filters to find the ones prefixed with 'control_tag', e.g. 'control_tag:nist'
	for filterType := range filters {
		if strings.HasPrefix(filterType, "control_tag:") {
			_, tagKey := leftSplit(filterType, ":")
			termQuery := backend.newNestedTermQueryFromControlTagsFilter(tagKey, filters[filterType])
			boolQuery = boolQuery.Must(termQuery)
		}
	}

	return boolQuery
}

func (backend ES2Backend) getESFieldName(filterType string) string {
	ESFieldName := filterType
	switch filterType {
	case "chef_server":
		ESFieldName = "source_fqdn.lower"
	case "inspec_version":
		ESFieldName = "version.lower"
	case "organization":
		ESFieldName = "organization_name.lower"
	case "platform":
		ESFieldName = "platform.name.lower"
	case "platform_with_version":
		ESFieldName = "platform.full.lower"
	case "recipe":
		ESFieldName = "recipes.lower"
	case "role":
		ESFieldName = "roles.lower"
	case "environment":
		ESFieldName = "environment.lower"
	case "chef_tags":
		ESFieldName = "chef_tags.lower"
	case "policy_group":
		ESFieldName = "policy_group.lower"
	case "policy_name":
		ESFieldName = "policy_name.lower"
	case "node_name":
		ESFieldName = "node_name.lower"
	}

	return ESFieldName
}

func (backend ES2Backend) newTermQueryFromFilter(ESField string,
	filters []string) *elastic.BoolQuery {
	refinedValues := make([]string, 0, 0)
	filterQuery := elastic.NewBoolQuery()

	for _, value := range filters {
		if containsWildcardChar(value) {
			wildQuery := elastic.NewWildcardQuery(ESField, value)
			filterQuery = filterQuery.Should(wildQuery)
		} else {
			refinedValues = append(refinedValues, value)
		}
	}
	if len(refinedValues) > 0 {
		termQuery := elastic.NewTermsQuery(ESField, stringArrayToInterfaceArray(refinedValues)...)
		filterQuery = filterQuery.Should(termQuery)
	}
	return filterQuery
}

func (backend ES2Backend) newNestedTermQueryFromFilter(ESField string, ESFieldPath string,
	filterValues []string) *elastic.BoolQuery {
	refinedValues := make([]string, 0, 0)
	filterQuery := elastic.NewBoolQuery()

	for _, filterValue := range filterValues {
		if containsWildcardChar(filterValue) {
			wildQuery := elastic.NewWildcardQuery(ESField, filterValue)
			nestedQuery := elastic.NewNestedQuery(ESFieldPath, wildQuery)
			filterQuery = filterQuery.Should(nestedQuery)
		} else {
			refinedValues = append(refinedValues, filterValue)
		}
	}
	if len(refinedValues) > 0 {
		termQuery := elastic.NewTermsQuery(ESField, stringArrayToInterfaceArray(refinedValues)...)
		nestedQuery := elastic.NewNestedQuery(ESFieldPath, termQuery)
		filterQuery = filterQuery.Should(nestedQuery)
	}

	return filterQuery
}

// Returns an ElasticSearch nested query to filter reports by control tags
func (backend ES2Backend) newNestedTermQueryFromControlTagsFilter(tagKey string, tagValues []string) *elastic.NestedQuery {
	refinedValues := make([]string, 0, 0)
	ESFieldPath := "profiles.controls.string_tags"
	ESFieldTagKey := "profiles.controls.string_tags.key.lower"
	ESFieldTagValues := "profiles.controls.string_tags.values.lower"

	// Add ElasticSearch query filter for the control tag key
	tagKey = strings.ToLower(tagKey)
	boolQuery := elastic.NewBoolQuery()
	if containsWildcardChar(tagKey) {
		boolQuery.Must(elastic.NewWildcardQuery(ESFieldTagKey, tagKey))
	} else {
		boolQuery.Must(elastic.NewTermsQuery(ESFieldTagKey, tagKey))
	}

	// Add ElasticSearch query filters for the control tag value(s)
	insideBoolQuery := elastic.NewBoolQuery()
	insideBool := false
	for _, tagValue := range tagValues {
		tagValue = strings.ToLower(tagValue)
		if containsWildcardChar(tagValue) {
			insideBoolQuery.Should(elastic.NewWildcardQuery(ESFieldTagValues, tagValue))
			insideBool = true
		} else {
			refinedValues = append(refinedValues, tagValue)
		}
	}
	// Here we handle control tag value(s) without wildcard characters
	if len(refinedValues) > 0 {
		termQuery := elastic.NewTermsQuery(ESFieldTagValues, stringArrayToInterfaceArray(refinedValues)...)
		insideBoolQuery.Should(termQuery)
		insideBool = true
	}
	if insideBool {
		boolQuery.Must(insideBoolQuery)
	}

	nestedQuery := elastic.NewNestedQuery(ESFieldPath, boolQuery)
	return nestedQuery
}

func containsWildcardChar(value string) bool {
	wildcardValues := []string{"*", "?"}

	for _, wildcard := range wildcardValues {
		if strings.Contains(value, wildcard) {
			return true
		}
	}

	return false
}

//  getFiltersQueryForDeepReport - builds up an elasticsearch query filter based on the reportId filters map passed in.
//   This func ignores all but the profile_id and control filter.  If a single profile_id or a single profile_id and a
//   child control from the single profile_id is passed in, then we are in deep filtering mode and the
//   fetchSourceContexts will be adjusted accordingly
//    arguments:
//			reportId - the id of the report we are querying for.
//			filters - is a map of filters that serve as the source for generated es query filters
//    return *elastic.BoolQuery
func (backend ES2Backend) getFiltersQueryForDeepReport(reportId string,
	filters map[string][]string) *elastic.BoolQuery {

	utils.DeDupFilters(filters)
	typeQuery := elastic.NewTypeQuery(mappings.DocType)

	boolQuery := elastic.NewBoolQuery()
	boolQuery = boolQuery.Must(typeQuery)

	idsQuery := elastic.NewIdsQuery(mappings.DocType)
	idsQuery.Ids(reportId)
	boolQuery = boolQuery.Must(idsQuery)

	if len(filters["projects"]) > 0 {
		projectsQuery := elastic.NewBoolQuery()

		if stringutils.SliceContains(filters["projects"], authzConstants.UnassignedProjectID) {
			emptyProjectQuery := elastic.NewBoolQuery()
			emptyProjectQuery.MustNot(elastic.NewExistsQuery("projects"))
			projectsQuery.Should(emptyProjectQuery)
		}

		assignedProjectIds := stringutils.SliceFilter(filters["projects"], func(projectId string) bool {
			return projectId != authzConstants.UnassignedProjectID
		})

		if len(assignedProjectIds) > 0 {
			projectMatchQuery := elastic.NewTermsQuery("projects", stringArrayToInterfaceArray(assignedProjectIds)...)
			projectsQuery.Should(projectMatchQuery)
		}

		boolQuery = boolQuery.Filter(projectsQuery)
	}

	numberOfProfiles := len(filters["profile_id"])
	numberOfControls := len(filters["control"])
	if numberOfProfiles > 0 || numberOfControls > 0 {
		profileBaseFscIncludes := []string{
			"profiles.depends",
			"profiles.name",
			"profiles.sha256",
			"profiles.status",
			"profiles.skip_message",
			"profiles.version"}

		profileLevelFscIncludes := []string{"profiles"}
		controlLevelFscIncludes := []string{"profiles.controls"}

		profileAndControlQuery := getProfileAndControlQuery(
			filters,
			profileBaseFscIncludes,
			profileLevelFscIncludes,
			controlLevelFscIncludes,
		)
		boolQuery = boolQuery.Must(profileAndControlQuery)
	}

	return boolQuery
}

func getProfileAndControlQuery(filters map[string][]string, profileBaseFscIncludes, profileLevelFscIncludes,
	controlLevelFscIncludes []string) *elastic.NestedQuery {
	var profileIds string
	var profileLevelFsc *elastic.FetchSourceContext
	numberOfProfiles := len(filters["profile_id"])
	numberOfControls := len(filters["control"])
	if numberOfProfiles == 0 {
		profileIds = ".*"
	} else {
		profileIds = strings.Join(filters["profile_id"], "|")
		if numberOfProfiles == 1 && numberOfControls <= 1 {
			//we are deep and when that's the case, we know we want at least this stuff.
			// if we don't have controls, we'll add even more to the profile level context (below)
			profileLevelFsc = elastic.NewFetchSourceContext(true).
				Include(profileBaseFscIncludes...)
		}
	}
	profileBoolQuery := elastic.NewBoolQuery()
	stringQuery := elastic.NewQueryStringQuery(fmt.Sprintf("profiles.sha256:/(%s)/", profileIds))
	profileBoolQuery.Must(stringQuery)

	var nestedQuery *elastic.NestedQuery
	nestedQuery = elastic.NewNestedQuery("profiles", profileBoolQuery)

	//add in the control query if it exists
	if numberOfControls > 0 {
		controlIds := strings.Join(filters["control"], "|")
		var nestedControlQuery *elastic.NestedQuery

		controlQuery := elastic.NewQueryStringQuery(fmt.Sprintf("profiles.controls.id:/(%s)/", controlIds))
		nestedControlQuery = elastic.NewNestedQuery("profiles.controls", controlQuery)

		if numberOfProfiles == 1 && numberOfControls == 1 {
			//we are control deep and when that's the case, we know we want this stuff in fsc.
			//we also know that we want to nest it so that we may use it for reports or node info
			fsc := elastic.NewFetchSourceContext(true).
				Include(controlLevelFscIncludes...)
			nestedControlQuery.InnerHit(elastic.NewInnerHit().FetchSourceContext(fsc))
		}

		profileBoolQuery = profileBoolQuery.Must(nestedControlQuery)
	} else if numberOfProfiles == 1 && profileLevelFsc != nil {
		//we are profile deep and when that's the case, we know we want this stuff in fsc.
		//for deep filtering, we only need to include profile level controls_sums and status at profile level,
		// if we are filtering at profile level.. to include it when we are at control level would be inefficient
		profileLevelFsc.Include(profileLevelFscIncludes...)
	}

	//we are deep so lets add the fetch source context
	if numberOfProfiles == 1 && numberOfControls <= 1 {
		nestedQuery.InnerHit(elastic.NewInnerHit().FetchSourceContext(profileLevelFsc))
	}
	return nestedQuery
}
