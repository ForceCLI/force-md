package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/reportType"
)

func init() {
	reportTypeCmd.AddCommand(reportType.FieldCmd)
	reportTypeCmd.AddCommand(reportType.SectionCmd)
	RootCmd.AddCommand(reportTypeCmd)
}

var reportTypeCmd = &cobra.Command{
	Use:   "reporttype [command] [flags] [filename]...",
	Short: "Manage Report Types",
}
