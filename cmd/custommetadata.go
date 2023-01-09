package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/custommetadata"
)

func init() {
	customMetadataCmd.AddCommand(custommetadata.TableCmd)
	RootCmd.AddCommand(customMetadataCmd)
}

var customMetadataCmd = &cobra.Command{
	Use:   "custommetadata [command] [flags] [filename]...",
	Short: "Manage Custom Metadata",
}
