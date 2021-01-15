package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/permissionset"
)

func init() {
	permissionSetCmd.AddCommand(permissionset.FieldPermissionsCmd)
	permissionSetCmd.AddCommand(permissionset.TidyCmd)
	permissionSetCmd.AddCommand(permissionset.ApexClassCmd)
	rootCmd.AddCommand(permissionSetCmd)
}

var permissionSetCmd = &cobra.Command{
	Use:   "permissionset",
	Short: "Manage Permission Sets",
}
