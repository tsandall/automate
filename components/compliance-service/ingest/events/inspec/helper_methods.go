package inspec

const (
	ResultStatusPassed         string = "passed"
	ResultStatusSkipped        string = "skipped"
	ResultStatusFailed         string = "failed"
	ResultStatusInformational  string = "informational"
	ResultStatusLoaded         string = "loaded"
	ControlImpactInformational string = "informational"
	ControlImpactMinor         string = "minor"
	ControlImpactMajor         string = "major"
	ControlImpactHigh          string = "high"
	ControlImpactCritical      string = "critical"
)

// Status calculates the overall status of all controls based on all results
func (control *Control) Status() (status string) {
	status = ResultStatusPassed
	for _, result := range control.Results {
		if result.Status == ResultStatusFailed {
			status = ResultStatusFailed
			break
		} else if result.Status == ResultStatusSkipped {
			status = ResultStatusSkipped
		}
	}
	return status
}

// ImpactName returns a human readable name for the impact
func (control *Control) ImpactName() (impact string) {
	if control.Impact == 0.0 {
		return ControlImpactInformational
	} else if control.Impact < 0.4 {
		return ControlImpactMinor
	} else if control.Impact < 0.7 {
		return ControlImpactMajor
	} else if control.Impact < 0.9 {
		return ControlImpactHigh
	}
	return ControlImpactCritical
}

// CVSS 3.0 (https://nvd.nist.gov/vuln-metrics/cvss)
// None | 0.0
// Low | 0.1-3.9
// Medium | 4.0-6.9
// High | 7.0-8.9
// Critical | 9.0-10.0
