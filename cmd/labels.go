package cmd

import (
	"github.com/ForceCLI/force-md/cmd/labels"
	"github.com/spf13/cobra"
)

func init() {
	labelsCmd.AddCommand(labels.TableCmd)
	RootCmd.AddCommand(labelsCmd)
}

var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Manage Custom Labels",
}
