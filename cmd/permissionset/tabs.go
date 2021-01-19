package permissionset

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var tabName string

func init() {
	addTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	addTabCmd.MarkFlagRequired("tab")

	TabCmd.AddCommand(addTabCmd)
}

var TabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Manage tab visibility",
}

var addTabCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tab visibility",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addTab(file)
		}
	},
}

func addTab(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	p.AddTab(tabName)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
