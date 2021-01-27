package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/objects"
)

func init() {
	objectsCmd.AddCommand(objects.FieldCmd)
	objectsCmd.AddCommand(objects.TidyCmd)
	RootCmd.AddCommand(objectsCmd)
}

var objectsCmd = &cobra.Command{
	Use:   "objects",
	Short: "Manage Custom and Standard Objects",
}
