package cve

import (
	"errors"
	"sort"
)

func SortCVEs(cve []CVE) {
	// Sort the CVE title list based on Threat Severity and Bug ID
	sort.Slice(cve, func(i, j int) bool {
		priorityI, errI := getSeverityPriority(cve[i].ThreatSeverity)
		priorityJ, errJ := getSeverityPriority(cve[j].ThreatSeverity)

		// Handle the error
		if errI != nil || errJ != nil {
			panic("Invalid threat severity")
		}

		if priorityI != priorityJ {
			// Sort by Threat Severity (Critical > Important > Moderate > Low)
			return priorityI > priorityJ
		}
		// If Threat Severity is the same, sort by Bug ID
		return cve[i].BugzillaID < cve[j].BugzillaID
	})
}

// getSeverityPriority is a helper function to get the priority of threat severity
func getSeverityPriority(severity string) (int, error) {
	switch severity {
	case "Critical":
		return 4, nil
	case "Important":
		return 3, nil
	case "Moderate":
		return 2, nil
	case "Low":
		return 1, nil
	default:
		return 0, errors.New("invalid threat severity")
	}
}
