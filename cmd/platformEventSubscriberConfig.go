package cmd

import (
	"github.com/ForceCLI/force-md/cmd/platformEventSubscriberConfig"
	"github.com/spf13/cobra"
)

func init() {
	platformEventSubscriberConfigCmd.AddCommand(platformEventSubscriberConfig.EditCmd)
	RootCmd.AddCommand(platformEventSubscriberConfigCmd)
}

var platformEventSubscriberConfigCmd = &cobra.Command{
	Use:   "platformEventSubscriberConfig",
	Short: "Manage Platform Event Subscriber Config Metadata",
}
