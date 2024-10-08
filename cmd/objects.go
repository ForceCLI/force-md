package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/objects"
)

func init() {
	objectsCmd.AddCommand(objects.FieldCmd)
	objectsCmd.AddCommand(objects.IndexCmd)
	objectsCmd.AddCommand(objects.FieldSetCmd)
	objectsCmd.AddCommand(objects.RecordTypeCmd)
	objectsCmd.AddCommand(objects.TidyCmd)
	objectsCmd.AddCommand(objects.ValidationRuleCmd)
	objectsCmd.AddCommand(objects.ActionCmd)
	objectsCmd.AddCommand(objects.WebLinkCmd)
	RootCmd.AddCommand(objectsCmd)
}

var objectsCmd = &cobra.Command{
	Use:   "objects [command] [flags] [filename]...",
	Short: "Manage Custom and Standard Objects",
}
