package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/permissionsetgroup"
)

func init() {
	permissionSetGroupCmd.AddCommand(permissionsetgroup.PermissionSetCmd)
	RootCmd.AddCommand(permissionSetGroupCmd)
}

var permissionSetGroupCmd = &cobra.Command{
	Use:   "permissionsetgroup",
	Short: "Manage Permission Set Groups",
}
