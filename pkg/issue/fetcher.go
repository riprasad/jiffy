package issue

import (
	"log"

	"github.com/andygrunwald/go-jira"
)

func GetInfo(token string, jql string) []jira.Issue {
	tp := jira.BearerAuthTransport{
		Token: token,
	}

	jiraClient, _ := jira.NewClient(tp.Client(), "https://issues.redhat.com/")

	issues, _, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		log.Fatal(err)
	}

	return issues
}
