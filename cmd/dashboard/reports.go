package dashboard

import (
	"fmt"

	"github.com/ForceCLI/force-md/dashboard"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ReportsCmd.AddCommand(reportsListCmd)
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
