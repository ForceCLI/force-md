package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/application"
)

func init() {
	applicationCmd.AddCommand(application.ActionCmd)
	applicationCmd.AddCommand(application.TabCmd)
	applicationCmd.AddCommand(application.TidyCmd)
	RootCmd.AddCommand(applicationCmd)
}

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manage Applications",
}
