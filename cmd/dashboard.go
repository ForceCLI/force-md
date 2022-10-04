package cmd

import (
	"github.com/octoberswimmer/force-md/cmd/dashboard"
	"github.com/spf13/cobra"
)

func init() {
	dashboardCmd.AddCommand(dashboard.EditCmd)
	dashboardCmd.AddCommand(dashboard.ReportsCmd)
	RootCmd.AddCommand(dashboardCmd)
}

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Manage Dashboards",
}
