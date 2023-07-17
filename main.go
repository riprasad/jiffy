package main

import (
	"fmt"

	"github.com/riprasad/jiffy/pkg/cve"
	"github.com/riprasad/jiffy/pkg/issue"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println("port is: ", viper.Get("token"))
	fmt.Println("username is: ", viper.Get("jql"))

	jql := viper.Get("jql").(string)

	//jql := `project = IPT AND status in ("Selected for Development") AND labels = security ORDER BY status DESC, created DESC, duedate`
	issues := issue.GetInfo(jql)

	cves := cve.CurateCveDetails(issues)

	for _, cveInfo := range cves {
		fmt.Printf("%s %s ", cveInfo.BugzillaID, cveInfo.JiraKey)
	}

	fmt.Println()
	fmt.Println()

	for _, cveInfo := range cves {
		fmt.Printf("%s: %s (%s)\n", cveInfo.AffectedPackage, cveInfo.Summary, cveInfo.ID)
	}

}
