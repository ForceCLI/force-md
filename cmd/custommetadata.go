package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/custommetadata"
)

func init() {
	customMetadataCmd.AddCommand(custommetadata.TableCmd)
	customMetadataCmd.AddCommand(custommetadata.ListCmd)
	customMetadataCmd.AddCommand(custommetadata.NewCmd)
	customMetadataCmd.AddCommand(custommetadata.EditCmd)
	RootCmd.AddCommand(customMetadataCmd)
}

var customMetadataCmd = &cobra.Command{
	Use:   "custommetadata [command] [flags] [filename]...",
	Short: "Manage Custom Metadata",
}
