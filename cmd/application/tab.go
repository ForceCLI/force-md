package application

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/application"
	"github.com/octoberswimmer/force-md/internal"
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
	Short: "Manage tabs",
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
	Use:                   "delete -c TabName [filename]...",
	Short:                 "Delete tab",
	Long:                  "Delete tab from application",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteTab(file, tabName)
		}
	},
}

var listTabsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List tabs",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listTabs(file)
		}
	},
}

func addTab(file string) {
	p, err := application.Open(file)
	if err != nil {
		log.Warn("parsing application failed: " + err.Error())
		return
	}
	p.AddTab(tabName)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func deleteTab(file string, tabName string) {
	p, err := application.Open(file)
	if err != nil {
		log.Warn("parsing application failed: " + err.Error())
		return
	}
	err = p.DeleteTab(tabName)
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

func listTabs(file string) {
	p, err := application.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	tabs := p.GetTabs()
	for _, a := range tabs {
		fmt.Println(a.Text)
	}
}
