package cmd

import (
	"github.com/ForceCLI/force-md/cmd/reportFolder"
	"github.com/spf13/cobra"
)

func init() {
	reportFolderCmd.AddCommand(reportFolder.FolderSharesCmd)
	RootCmd.AddCommand(reportFolderCmd)
}

var reportFolderCmd = &cobra.Command{
	Use:   "reportfolder",
	Short: "Manage Report Folders",
}
