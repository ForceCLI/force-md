package cmd

import (
	"github.com/octoberswimmer/force-md/cmd/globalvalueset"
	"github.com/spf13/cobra"
)

func init() {
	globalValueSetCmd.AddCommand(globalvalueset.EditCmd)
	globalValueSetCmd.AddCommand(globalvalueset.TidyCmd)
	RootCmd.AddCommand(globalValueSetCmd)
}

var globalValueSetCmd = &cobra.Command{
	Use:   "globalvalueset",
	Short: "Manage Global Value Sets",
}
