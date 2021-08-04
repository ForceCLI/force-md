package cmd

import (
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/cmd/pkg"
)

func init() {
	packageCmd.AddCommand(pkg.AddCmd)
	packageCmd.AddCommand(pkg.DeleteCmd)
	packageCmd.AddCommand(pkg.TidyCmd)
	packageCmd.AddCommand(pkg.ListCmd)
	RootCmd.AddCommand(packageCmd)
}

var packageCmd = &cobra.Command{
	Use:   "package [command] [flags] [filename]...",
	Short: "Manage package.xml or destructiveChanges[Pre|Post].xml",
}
