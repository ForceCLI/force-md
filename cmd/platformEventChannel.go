package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/platformEventChannels"
)

func init() {
	platformEventChannelCmd.AddCommand(platformEventChannels.NewCmd)
	RootCmd.AddCommand(platformEventChannelCmd)
}

var platformEventChannelCmd = &cobra.Command{
	Use:   "platformEventChannel [command] [flags] [filename]...",
	Short: "Manage Platform Event Channels",
}
