package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/queue"
)

func init() {
	queueCmd.AddCommand(queue.MemberCmd)
	RootCmd.AddCommand(queueCmd)
}

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Manage Queues",
}
