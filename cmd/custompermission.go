package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/custompermission"
)

func init() {
	customPermissionCmd.AddCommand(custompermission.NewCmd)
	RootCmd.AddCommand(customPermissionCmd)
}

var customPermissionCmd = &cobra.Command{
	Use:   "custompermission [command] [flags] [filename]...",
	Short: "Manage Custom Permissions",
}
