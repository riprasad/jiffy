package issue

import (
	"log"

	"github.com/andygrunwald/go-jira"
)

func GetInfo(jql string) []jira.Issue {
	tp := jira.BearerAuthTransport{
		Token: "ODg1ODg0NTM1NDA4OuWMXiwdwXnSTQ4bbjowWQHP51ec",
	}

	jiraClient, _ := jira.NewClient(tp.Client(), "https://issues.redhat.com/")

	issues, _, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		log.Fatal(err)
	}

	return issues
}
