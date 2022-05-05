package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/matchingrules"
)

func init() {
	matchingRulesCmd.AddCommand(matchingrules.ListCmd)
	matchingRulesCmd.AddCommand(matchingrules.DeleteCmd)
	RootCmd.AddCommand(matchingRulesCmd)
}

var matchingRulesCmd = &cobra.Command{
	Use:   "matchingrules",
	Short: "Manage Matching Rules",
}
