package rhsa

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/riprasad/jiffy/pkg/config"
	"github.com/riprasad/jiffy/pkg/cve"
	"github.com/riprasad/jiffy/pkg/issue"
	"github.com/spf13/cobra"
)

// will show CVE infos like BugzillaID, JiraKey, AffectedPackage, Summary and ID
func Curate() *cobra.Command {
	var outputFormat string

	cfg := config.GetConfiguration()
	if cfg.JiraToken == "" || cfg.Jql == "" {
		fmt.Println(`Either Jira Token or Jql is empty. Please pass both values in "config.yaml" to continue :)`)
		os.Exit(1)
	}

	token := cfg.JiraToken
	jql := cfg.Jql
	issues := issue.GetInfo(token, jql)
	cves := cve.CurateCveDetails(issues)

	cmd := &cobra.Command{
		Use:     "curate",
		Short:   "Get cve issues from jira",
		Long:    "",
		Example: "",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			if outputFormat == "" {
				for _, cveInfo := range cves {
					fmt.Printf("%s %s ", cveInfo.BugzillaID, cveInfo.JiraKey)
				}

				fmt.Println()
				fmt.Println()

				for _, cveInfo := range cves {
					fmt.Printf("%s: %s (%s)\n", cveInfo.AffectedPackage, cveInfo.Summary, cveInfo.ID)
				}
			}

			if outputFormat == "table" {
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
			}
		},
	}
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "", `Format of your default Output is normal form and "= table" for table form`)
	return cmd
}
