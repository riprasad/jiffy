package jiffy

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/riprasad/jiffy/pkg/config"
	"github.com/riprasad/jiffy/pkg/cve"
	"github.com/riprasad/jiffy/pkg/issue"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jiffy",
	Short: "Intro command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jiffy is a small helpful robot that automates laborious and time-consuming manual tasks of collecting critical CVE information required for filing Red Hat Security Advisories (RHSA). Named after its ability to perform tasks swiftly, Jiffy streamlines the process saving you valuable time and effort.")
	},
}

// will show CVE infos like BugzillaID, JiraKey, AffectedPackage, Summary and ID
var Info = &cobra.Command{
	Use:   "info",
	Short: "Print CVEs info of passed token and jql",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfiguration()
		if cfg.JiraToken == "" || cfg.Jql == "" {
			fmt.Println(`Either Jira Token or Jql is empty. Please pass both values in "config.yaml" to continue :)`)
			os.Exit(1)
		}

		token := cfg.JiraToken
		jql := cfg.Jql
		issues := issue.GetInfo(token, jql)
		cves := cve.CurateCveDetails(issues)

		// print in proper table form on stdout
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "%s\t%s\t\n", "Bugzilla_ID", "Jira_Key")
		for _, cveInfo := range cves {
			fmt.Printf("%s %s ", cveInfo.BugzillaID, cveInfo.JiraKey)
		}
		w.Flush()

		fmt.Println()
		fmt.Println()

		// print in proper table form on stdout
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", "ID", "Affected_Package", "Summary")
		for _, cveInfo := range cves {
			fmt.Printf("%s %s %s", cveInfo.ID, cveInfo.AffectedPackage, cveInfo.Summary)
		}
		w.Flush()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.InitConfiguration()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(Info)
}
