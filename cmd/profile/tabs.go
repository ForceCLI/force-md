package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var tabName string

func init() {
	deleteTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	deleteTabCmd.MarkFlagRequired("tab")

	TabCmd.AddCommand(deleteTabCmd)
}

var TabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Manage tab visibility",
}

var deleteTabCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tab visibility",
	Long:  "Delete tab visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteTabVisibility(file, tabName)
		}
	},
}

func deleteTabVisibility(file string, tabName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteTabVisibility(tabName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
