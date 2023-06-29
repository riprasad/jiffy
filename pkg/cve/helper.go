package cve

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/andygrunwald/go-jira"
)

type CVE struct {
	ID              string
	ThreatSeverity  string
	BugzillaID      string
	AffectedPackage string
	Summary         string
	JiraKey         string
}

func extractCveDetails(jiraTitle string) (string, string, string, string) {
	// Extracting the CVE number
	startIndex := 0
	endIndex := strings.Index(jiraTitle, " ")
	/* if startIndex == -1 || endIndex == -1 {
		return "", "", errors.New("cannot extract CVE information")
	} */
	cveID := jiraTitle[startIndex:endIndex]

	// Extracting the affected package name
	startIndex = strings.Index(jiraTitle, " ")
	endIndex = strings.Index(jiraTitle, ":")
	affectedPackage := jiraTitle[startIndex+1 : endIndex]

	// Extracting the CVE summary
	startIndex = strings.Index(jiraTitle, ":")
	endIndex = strings.Index(jiraTitle, "[")
	cveSummary := jiraTitle[startIndex+2 : endIndex-1]

	// Extracting the affected product name
	startIndex = strings.Index(jiraTitle, "[")
	endIndex = strings.Index(jiraTitle, "]")
	affectedProduct := jiraTitle[startIndex+1 : endIndex]

	return cveID, affectedPackage, cveSummary, affectedProduct
}

func checkRegexMatch(input string) (bool, error) {
	regexPattern := "^CVE-(.*?) (.*?): (.*?) \\[(.*?)\\]"
	matched, err := regexp.MatchString(regexPattern, input)
	if err != nil {
		return false, errors.New("error occurred while matching regex")
	}
	return matched, nil
}

func CurateCveDetails(issues []jira.Issue) []CVE {

	fmt.Println("Number of issue: ", len(issues))
	// Create a slice of CveTitle struct to store the CVE titles with their details
	cveDetails := make([]CVE, len(issues))

	for i, issue := range issues {
		jiraKey := issue.Key
		jiraTitle := issue.Fields.Summary

		matched, err := checkRegexMatch(jiraTitle)
		if err != nil {
			fmt.Println("Error:", err)
		} else if matched {
			cveID, affectedPackage, cveSummary, _ := extractCveDetails(jiraTitle)
			cveStatus, err := getInfo(cveID)
			if err != nil {
				log.Fatal(err)
			}

			threatSeverity := cveStatus.ThreatSeverity
			bugId := cveStatus.Bugzilla.ID

			cveDetails[i] = CVE{
				ID:              cveID,
				ThreatSeverity:  threatSeverity,
				BugzillaID:      bugId,
				AffectedPackage: affectedPackage,
				Summary:         cveSummary,
				JiraKey:         jiraKey,
			}
		} else {
			fmt.Println("Error: ", errors.New("jira issue title does not match the expected pattern"))
		}

	}
	SortCVEs(cveDetails)

	return cveDetails
}
