package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/profile"
)

func init() {
	profileCmd.AddCommand(profile.FieldPermissionsCmd)
	profileCmd.AddCommand(profile.TidyCmd)
	profileCmd.AddCommand(profile.ObjectPermissionsCmd)
	profileCmd.AddCommand(profile.TabCmd)
	profileCmd.AddCommand(profile.UserPermissionsCmd)
	profileCmd.AddCommand(profile.ApplicationCmd)
	profileCmd.AddCommand(profile.FlowCmd)
	profileCmd.AddCommand(profile.ApexCmd)
	profileCmd.AddCommand(profile.VisualforceCmd)
	profileCmd.AddCommand(profile.LayoutCmd)
	profileCmd.AddCommand(profile.NewCmd)
	rootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage Profiles",
}
