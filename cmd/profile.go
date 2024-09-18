package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/profile"
)

func init() {
	profileCmd.AddCommand(profile.FieldPermissionsCmd)
	profileCmd.AddCommand(profile.TidyCmd)
	profileCmd.AddCommand(profile.ObjectPermissionsCmd)
	profileCmd.AddCommand(profile.TabCmd)
	profileCmd.AddCommand(profile.UserPermissionsCmd)
	profileCmd.AddCommand(profile.CustomPermissionsCmd)
	profileCmd.AddCommand(profile.ApplicationCmd)
	profileCmd.AddCommand(profile.FlowCmd)
	profileCmd.AddCommand(profile.ApexCmd)
	profileCmd.AddCommand(profile.VisualforceCmd)
	profileCmd.AddCommand(profile.LayoutCmd)
	profileCmd.AddCommand(profile.RecordTypeCmd)
	profileCmd.AddCommand(profile.NewCmd)
	profileCmd.AddCommand(profile.MergeCmd)
	profileCmd.AddCommand(profile.LoginFlowCmd)
	RootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage Profiles",
}
