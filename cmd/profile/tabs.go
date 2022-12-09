package profile

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var tabName string
var tabVisibility TabVisibility

type TabVisibility enumflag.Flag

const (
	DefaultOn TabVisibility = iota
	DefaultOff
	Hidden
)

var TabVisibilityIds = map[TabVisibility][]string{
	DefaultOn:  {"DefaultOn"},
	DefaultOff: {"DefaultOff"},
	Hidden:     {"Hidden"},
}

func init() {
	addTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	addTabCmd.MarkFlagRequired("tab")

	deleteTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	deleteTabCmd.MarkFlagRequired("tab")

	editTabCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")
	editTabCmd.Flags().VarP(enumflag.New(&tabVisibility, "visibility", TabVisibilityIds, enumflag.EnumCaseInsensitive),
		"visibility", "v", "tab visibility; can be 'DefaultOn', 'DefaultOff', or 'Hidden'")
	editTabCmd.MarkFlagRequired("tab")
	editTabCmd.MarkFlagRequired("visibility")

	tableTabsCmd.Flags().StringVarP(&tabName, "tab", "t", "", "tab name")

	TabCmd.AddCommand(addTabCmd)
	TabCmd.AddCommand(deleteTabCmd)
	TabCmd.AddCommand(editTabCmd)
	TabCmd.AddCommand(listTabsCmd)
	TabCmd.AddCommand(tableTabsCmd)
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

var tableTabsCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Tabs in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableTabPermissions(args)
	},
}

var editTabCmd = &cobra.Command{
	Use:   "edit -t TabName [flags] [filename]...",
	Short: "Edit tab visibility",
	Long:  "Edit tab visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			updateTabVisibility(file, tabVisibility)
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

func updateTabVisibility(file string, v TabVisibility) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	var visibility string
	switch v {
	case DefaultOn:
		visibility = "DefaultOn"
	case DefaultOff:
		visibility = "DefaultOff"
	case Hidden:
		visibility = "Hidden"
	}
	err = p.SetTabVisibility(tabName, visibility)
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

func tableTabPermissions(files []string) {
	var filters []profile.TabFilter
	if tabName != "" {
		filters = append(filters, func(f profile.TabVisibility) bool {
			return strings.ToLower(f.Tab) == strings.ToLower(tabName)
		})
	}
	type perm struct {
		tabs    profile.TabVisibilityList
		profile string
	}
	var perms []perm
	for _, file := range files {
		p, err := profile.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		profileName := strings.TrimSuffix(path.Base(file), ".profile")
		perms = append(perms, perm{tabs: p.GetTabs(filters...), profile: profileName})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "Tab", "Visibility"})
	table.SetRowLine(true)
	for _, perm := range perms {
		for _, t := range perm.tabs {
			table.Append([]string{perm.profile, t.Tab,
				t.Visibility,
			})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
