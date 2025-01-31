package server

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	rrule "github.com/teambition/rrule-go"

	es "github.com/chef/automate/api/interservice/es_sidecar"
	"github.com/chef/automate/api/interservice/ingest"
	"github.com/chef/automate/components/ingest-service/backend"
	"github.com/chef/automate/components/ingest-service/config"
	"github.com/chef/automate/lib/cereal"
	"github.com/chef/automate/lib/cereal/patterns"
	"github.com/chef/automate/lib/datalifecycle/purge"
)

var (
	MissingNodesWorkflowName = cereal.NewWorkflowName("missing_nodes")
	MissingNodesScheduleName = "periodic_missing_nodes"

	DeleteNodesWorkflowName = cereal.NewWorkflowName("delete_nodes")
	DeleteNodesScheduleName = "periodic_delete_nodes"

	MissingNodesForDeletionWorkflowName = cereal.NewWorkflowName("missing_nodes_for_deletion")
	MissingNodesForDeletionScheduleName = "periodic_missing_nodes_for_deletion"
)

func InitializeJobManager(c *cereal.Manager, client backend.Client, esSidecarClient es.EsSidecarClient) error {
	err := patterns.RegisterSingleTaskWorkflowExecutor(c, DeleteNodesWorkflowName, false, &DeleteExpiredMarkedNodesTask{client}, cereal.TaskExecutorOpts{})
	if err != nil {
		return err
	}
	err = patterns.RegisterSingleTaskWorkflowExecutor(c, MissingNodesWorkflowName, false, &MarkNodesMissingTask{client}, cereal.TaskExecutorOpts{})
	if err != nil {
		return err
	}
	err = patterns.RegisterSingleTaskWorkflowExecutor(c, MissingNodesForDeletionWorkflowName, false, &MarkMissingNodesForDeletionTask{client}, cereal.TaskExecutorOpts{})
	if err != nil {
		return err
	}
	err = purge.ConfigureManager(
		c, PurgeWorkflowName, purge.WithTaskEsSidecarClient(esSidecarClient),
	)
	if err != nil {
		return errors.Wrap(err, "failed to configure purge workflow")
	}

	return nil
}

func MigrateJobsSchedule(ctx context.Context, c *cereal.Manager, oldConfigFile string) error {
	jc, err := config.OldJobConfigFromFile(oldConfigFile)
	if err != nil {
		log.WithError(err).Warn("failed to read old job config from disk, defaults will be used")
	}

	for idx, name := range config.JobList {
		scheduleName := fmt.Sprintf("periodic_%s", name)
		config, err := jc.ConfigForJob(idx)
		if err != nil {
			return err
		}

		d, err := time.ParseDuration(config.Every)
		if err != nil {
			return errors.Wrap(err, "could not parse job configuration")
		}
		r, err := rrule.NewRRule(rrule.ROption{
			Freq:     rrule.SECONDLY,
			Interval: int(d.Seconds()),
			Dtstart:  time.Now(),
		})
		if err != nil {
			return errors.Wrap(err, "could not create recurrence rule for job configuration")
		}

		err = c.CreateWorkflowSchedule(ctx, scheduleName, cereal.NewWorkflowName(name), config.Threshold, config.Running, r)
		if err == cereal.ErrWorkflowScheduleExists {
			log.Infof("Schedule for %s already exists, not migrating", scheduleName)
		} else if err != nil {
			return errors.Wrap(err, "could not create scheduled workflow")
		}
	}
	return nil
}

func workflowNameToInstanceName(workflowName cereal.WorkflowName) string {
	return fmt.Sprintf("periodic_%s", workflowName)
}

type DeleteExpiredMarkedNodesTask struct {
	Client backend.Client
}

func (t *DeleteExpiredMarkedNodesTask) Run(ctx context.Context, task cereal.Task) (interface{}, error) {
	var threshold string
	err := task.GetParameters(&threshold)
	if err != nil {
		return nil, errors.Wrap(err, "could not get threshold parameter")
	}

	return nil, deleteExpiredMarkedNodes(ctx, threshold, t.Client)
}

type MarkNodesMissingTask struct {
	Client backend.Client
}

func (t *MarkNodesMissingTask) Run(ctx context.Context, task cereal.Task) (interface{}, error) {
	var threshold string
	err := task.GetParameters(&threshold)
	if err != nil {
		return nil, errors.Wrap(err, "could not get threshold parameter")
	}

	return nil, markNodesMissing(ctx, threshold, t.Client)
}

type MarkMissingNodesForDeletionTask struct {
	Client backend.Client
}

func (t *MarkMissingNodesForDeletionTask) Run(ctx context.Context, task cereal.Task) (interface{}, error) {
	var threshold string
	err := task.GetParameters(&threshold)
	if err != nil {
		return nil, errors.Wrap(err, "could not get threshold parameter")
	}

	return nil, markMissingNodesForDeletion(ctx, threshold, t.Client)
}

// markNodesMissing is a job that will call the backend to mark
// all nodes that haven't checked in passed the threshold
func markNodesMissing(ctx context.Context, threshold string, client backend.Client) error {
	logctx := log.WithFields(log.Fields{
		"threshold": threshold,
		"job":       "MarkMissingNodesForDeletion",
	})

	logctx.Debug("Starting job")
	updateCount, err := client.MarkNodesMissing(ctx, threshold)
	if err != nil {
		logctx.WithError(err).Error("Job failed")
		return err
	}

	f := log.Fields{"nodes_updated": updateCount}
	if updateCount > 0 {
		logctx.WithFields(f).Info("Job updated nodes")
	} else {
		logctx.WithFields(f).Debug("Job ran without updates")
	}
	return nil
}

// markMissingNodesForDeletion is a job that will call the backend to mark all missing nodes
// that haven't checked in passed the threshold ready for deletion
func markMissingNodesForDeletion(ctx context.Context, threshold string, client backend.Client) error {
	logctx := log.WithFields(log.Fields{
		"threshold": threshold,
		"job":       "MarkMissingNodesForDeletion",
	})

	logctx.Debug("Starting job")
	updateCount, err := client.MarkMissingNodesForDeletion(ctx, threshold)
	if err != nil {
		log.WithError(err).Error("Job failed")
		return err
	}

	f := log.Fields{"nodes_updated": updateCount}
	if updateCount > 0 {
		logctx.WithFields(f).Info("Job updated nodes")
	} else {
		logctx.WithFields(f).Debug("Job ran without updates")
	}
	return nil
}

// deleteExpiredMarkedNodes is a job that will call the backend to delete all expired
// nodes marked for deletion
func deleteExpiredMarkedNodes(ctx context.Context, threshold string, client backend.Client) error {
	logctx := log.WithFields(log.Fields{
		"job":       "DeleteExpiredMarkedNodes",
		"threshold": threshold,
	})

	logctx.Debug("Starting Job")
	updateCount, err := client.DeleteMarkedNodes(ctx, threshold)
	if err != nil {
		logctx.WithError(err).Error("Job Failed")
		return err
	}

	f := log.Fields{"nodes_deleted": updateCount}
	if updateCount > 0 {
		logctx.WithFields(f).Info("Job deleted nodes")
	} else {
		logctx.WithFields(f).Debug("Job ran without updates")
	}
	return nil
}

func JobSettingsToUpdateOpts(settings *ingest.JobSettings, oldSchedule *cereal.Schedule) ([]cereal.WorkflowScheduleUpdateOpt, bool, error) {
	shouldRunNow := false
	err := settings.Validate()
	if err != nil {
		return nil, shouldRunNow, err
	}

	ret := make([]cereal.WorkflowScheduleUpdateOpt, 0)
	if e := settings.GetEvery(); len(e) > 0 {
		// Convert duration to an rrule
		d, err := time.ParseDuration(e)
		if err != nil {
			// Unlikely as validate already checked this
			return nil, shouldRunNow, err
		}
		oldRecurrence, err := oldSchedule.GetRRule()
		if err != nil {
			return nil, shouldRunNow, errors.Wrap(err, "could not parse existing recurrence rule")
		}

		r, err := rrule.NewRRule(rrule.ROption{
			Freq:     rrule.SECONDLY,
			Interval: int(d.Seconds()),
			Dtstart:  oldRecurrence.OrigOptions.Dtstart,
		})
		if err != nil {
			return nil, shouldRunNow, errors.Wrap(err, "could not construct new recurrence rule")
		}
		ret = append(ret, cereal.UpdateRecurrence(r))
	}

	if t := settings.GetThreshold(); len(t) > 0 {
		var oldThreshold string
		err := oldSchedule.GetParameters(&oldThreshold)
		if err != nil {
			oldThreshold = ""
		}

		if settings.GetRunning() && oldThreshold != t {
			shouldRunNow = true
		}
		ret = append(ret, cereal.UpdateParameters(t))
	}

	if !oldSchedule.Enabled && settings.GetRunning() {
		shouldRunNow = true
	}

	ret = append(ret, cereal.UpdateEnabled(settings.GetRunning()))

	return ret, shouldRunNow, nil
}
