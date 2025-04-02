package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/platformEventChannelMembers"
)

func init() {
	platformEventChannelMemberCmd.AddCommand(platformEventChannelMembers.NewCmd)
	RootCmd.AddCommand(platformEventChannelMemberCmd)
}

var platformEventChannelMemberCmd = &cobra.Command{
	Use:   "platformEventChannelMember [command] [flags] [filename]...",
	Short: "Manage Platform Event Channel Members",
}
