package jiffy

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/riprasad/jiffy/pkg/cve"
	"github.com/riprasad/jiffy/pkg/issue"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jiffy",
	Short: "Intro command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jiffy is a small helpful robot that automates laborious and time-consuming manual tasks of collecting critical CVE information required for filing Red Hat Security Advisories (RHSA). Named after its ability to perform tasks swiftly, Jiffy streamlines the process saving you valuable time and effort.")
	},
}

var Info = &cobra.Command{
	Use:   "Info",
	Short: "Print CVEs info of passed token and jql",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("in info section start")
		viper.SetConfigFile(".env")
		viper.ReadInConfig()

		token := viper.Get("token").(string)
		jql := viper.Get("jql").(string)
		issues := issue.GetInfo(token, jql)
		cves := cve.CurateCveDetails(issues)

		// print in proper table form on stdout
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "%s\t%s\t\n", "BugzillaID", "JiraKey")
		for _, cveInfo := range cves {
			fmt.Printf("%s %s ", cveInfo.BugzillaID, cveInfo.JiraKey)
		}
		w.Flush()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jiffy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(Info)
}
