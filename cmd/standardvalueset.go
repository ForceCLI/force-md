package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/standardvalueset"
)

func init() {
	standardValueSetCmd.AddCommand(standardvalueset.ListCmd)
	RootCmd.AddCommand(standardValueSetCmd)
}

var standardValueSetCmd = &cobra.Command{
	Use:   "standardvalueset",
	Short: "Manage Standard Value Sets",
}
