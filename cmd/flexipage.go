package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/cmd/flexipage"
)

func init() {
	flexiPageCmd.AddCommand(flexipage.FieldCmd)
	RootCmd.AddCommand(flexiPageCmd)
}

var flexiPageCmd = &cobra.Command{
	Use:   "flexipage",
	Short: "Manage Lightning Pages (FlexiPages)",
}
