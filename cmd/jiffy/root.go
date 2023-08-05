package jiffy

import (
	"fmt"
	"os"

	"github.com/riprasad/jiffy/cmd/jiffy/rhsa"
	"github.com/riprasad/jiffy/pkg/config"
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
func RhsaCurate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rhsa",
		Short: "Print CVEs info of passed token and jql",
	}
	cmd.AddCommand(rhsa.Curate())
	return cmd
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
	rootCmd.AddCommand(RhsaCurate())
}
