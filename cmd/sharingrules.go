package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/sharingrules"
)

func init() {
	sharingRulesCmd.AddCommand(sharingrules.CriteriaCmd)
	sharingRulesCmd.AddCommand(sharingrules.OwnerCmd)
	RootCmd.AddCommand(sharingRulesCmd)
}

var sharingRulesCmd = &cobra.Command{
	Use:   "sharingrules",
	Short: "Manage Sharing Rules",
}
