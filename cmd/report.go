package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/report"
)

func init() {
	reportCmd.AddCommand(report.FieldCmd)
	RootCmd.AddCommand(reportCmd)
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Manage Reports",
}
