package dashboard

import (
	"fmt"

	"github.com/ForceCLI/force-md/dashboard"
	"github.com/ForceCLI/force-md/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	reportsDeleteCmd.Flags().StringP("report", "r", "", "report name")
	reportsDeleteCmd.MarkFlagRequired("report")

	ReportsCmd.AddCommand(reportsListCmd)
	ReportsCmd.AddCommand(reportsDeleteCmd)
}

var ReportsCmd = &cobra.Command{
	Use:   "reports",
	Short: "Manage reports used by dashboard",
}

var reportsListCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List reports used by dashboard",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listReports(file)
		}
	},
}

var reportsDeleteCmd = &cobra.Command{
	Use:                   "delete -r <report> [filename]...",
	Short:                 "Delete report from dashboard",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			report, _ := cmd.Flags().GetString("report")
			deleteReport(file, report)
		}
	},
}

func listReports(file string) {
	a, err := dashboard.Open(file)
	if err != nil {
		log.Warn("parsing dashboard failed: " + err.Error())
		return
	}
	reports := a.GetReports()
	for _, r := range reports {
		fmt.Println(r)
	}
}

func deleteReport(file string, report string) {
	a, err := dashboard.Open(file)
	if err != nil {
		log.Warn("parsing dashboard failed: " + err.Error())
		return
	}
	err = a.DeleteReport(report)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(a, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
