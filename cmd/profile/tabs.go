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
	addTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	addTabCmd.MarkFlagRequired("tab")

	deleteTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	deleteTabCmd.MarkFlagRequired("tab")

	TabCmd.AddCommand(addTabCmd)
	TabCmd.AddCommand(deleteTabCmd)
	TabCmd.AddCommand(listTabsCmd)
}

var TabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Manage tab visibility",
}

var addTabCmd = &cobra.Command{
	Use:   "add -t TabName [flags] [filename]...",
	Short: "Add tab visibility",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addTab(file)
		}
	},
}

var deleteTabCmd = &cobra.Command{
	Use:   "delete -t TabName [flags] [filename]...",
	Short: "Delete tab visibility",
	Long:  "Delete tab visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteTabVisibility(file, tabName)
		}
	},
}

var listTabsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List tab visibility",
	Long:                  "List tab visibility in profiles",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listTabVisibility(file)
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

func addTab(file string) {
	p, err := profile.Open(file)
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

func listTabVisibility(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	tabs := p.GetTabs()
	for _, t := range tabs {
		fmt.Printf("%s: %s\n", t.Tab, t.Visibility)
	}
}
