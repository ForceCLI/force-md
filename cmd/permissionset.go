package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/permissionset"
)

func init() {
	permissionSetCmd.AddCommand(permissionset.FieldPermissionsCmd)
	permissionSetCmd.AddCommand(permissionset.TidyCmd)
	permissionSetCmd.AddCommand(permissionset.TabCmd)
	permissionSetCmd.AddCommand(permissionset.ApexCmd)
	permissionSetCmd.AddCommand(permissionset.VisualforceCmd)
	permissionSetCmd.AddCommand(permissionset.NewCmd)
	permissionSetCmd.AddCommand(permissionset.ObjectPermissionsCmd)
	permissionSetCmd.AddCommand(permissionset.UserPermissionsCmd)
	permissionSetCmd.AddCommand(permissionset.CustomPermissionsCmd)
	permissionSetCmd.AddCommand(permissionset.CustomMetadataTypesCmd)
	permissionSetCmd.AddCommand(permissionset.ApplicationCmd)
	permissionSetCmd.AddCommand(permissionset.RecordTypeCmd)
	permissionSetCmd.AddCommand(permissionset.MergeCmd)
	permissionSetCmd.AddCommand(permissionset.EditCmd)
	RootCmd.AddCommand(permissionSetCmd)
}

var permissionSetCmd = &cobra.Command{
	Use:   "permissionset",
	Short: "Manage Permission Sets",
}
