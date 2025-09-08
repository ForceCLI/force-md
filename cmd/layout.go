package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/layout"
)

func init() {
	layoutCmd.AddCommand(layout.FieldCmd)
	RootCmd.AddCommand(layoutCmd)
}

var layoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "Manage Page Layouts",
}
