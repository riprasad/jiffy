package main

import (
	"fmt"

	jiffy "github.com/riprasad/jiffy/cmd"
)

func main() {
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()

	// token := viper.Get("token").(string)
	// jql := viper.Get("jql").(string)

	// for COBRA CLI
	jiffy.Execute()

	// issues := issue.GetInfo(token, jql)
	// cves := cve.CurateCveDetails(issues)

	// for _, cveInfo := range cves {
	// 	fmt.Printf("%s %s ", cveInfo.BugzillaID, cveInfo.JiraKey)
	// }

	fmt.Println()
	fmt.Println()

	// for _, cveInfo := range cves {
	// 	fmt.Printf("%s: %s (%s)\n", cveInfo.AffectedPackage, cveInfo.Summary, cveInfo.ID)
	// }
	
}