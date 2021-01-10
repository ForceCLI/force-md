package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/profile"
)

func init() {
	profileCmd.AddCommand(profile.FieldPermissionsCmd)
	profileCmd.AddCommand(profile.TidyCmd)
	profileCmd.AddCommand(profile.ObjectPermissionsCmd)
	rootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage Profiles",
}
