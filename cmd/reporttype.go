package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/reporttype"
)

func init() {
	reportTypeCmd.AddCommand(reporttype.RemoveCmd)
	rootCmd.AddCommand(reportTypeCmd)
}

var reportTypeCmd = &cobra.Command{
	Use:   "reporttype",
	Short: "Manage Report Types",
}
