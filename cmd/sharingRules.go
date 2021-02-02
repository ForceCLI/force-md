package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/sharingRules"
)

func init() {
	sharingRulesCmd.AddCommand(sharingRules.CriteriaRulesCmd)
	sharingRulesCmd.AddCommand(sharingRules.OwnerRulesCmd)
	RootCmd.AddCommand(sharingRulesCmd)
}

var sharingRulesCmd = &cobra.Command{
	Use:   "sharingRules",
	Short: "Manage Sharing Rules",
}
