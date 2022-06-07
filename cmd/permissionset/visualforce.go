package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var pageName string
var sourcePageName string

func init() {
	addVisualforcePageCmd.Flags().StringVarP(&pageName, "page", "p", "", "visualforce page name")
	addVisualforcePageCmd.MarkFlagRequired("page")

	deleteVisualforcePageCmd.Flags().StringVarP(&pageName, "page", "p", "", "visualforce page name")
	deleteVisualforcePageCmd.MarkFlagRequired("page")

	cloneVisualforcePageCmd.Flags().StringVarP(&pageName, "page", "p", "", "visualforce page name")
	cloneVisualforcePageCmd.Flags().StringVarP(&sourcePageName, "source", "s", "", "source page name")
	cloneVisualforcePageCmd.MarkFlagRequired("page")
	cloneVisualforcePageCmd.MarkFlagRequired("source")

	VisualforceCmd.AddCommand(addVisualforcePageCmd)
	VisualforceCmd.AddCommand(cloneVisualforcePageCmd)
	VisualforceCmd.AddCommand(deleteVisualforcePageCmd)
	VisualforceCmd.AddCommand(listVisualforcePageCmd)
}

var VisualforceCmd = &cobra.Command{
	Use:   "visualforce",
	Short: "Manage visualforce page visibility",
}

var addVisualforcePageCmd = &cobra.Command{
	Use:                   "add -p PageName [filename]...",
	Short:                 "Add visualforce page access",
	Long:                  "Add visualforce page access to permission sets",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addVisualforceAccess(file, pageName)
		}
	},
}

var cloneVisualforcePageCmd = &cobra.Command{
	Use:                   "clone -s Page -f Page [filename]...",
	Short:                 "Clone visualforce permissions",
	Long:                  "Clone visualforce page permissions in permission sets for a new page",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			cloneVisualforcePageAccess(file)
		}
	},
}

var deleteVisualforcePageCmd = &cobra.Command{
	Use:                   "delete -p PageName [flags] [filename]...",
	Short:                 "Delete visualforce page access",
	Long:                  "Delete visualforce page access from permission sets",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteVisualforceAccess(file, pageName)
		}
	},
}

var listVisualforcePageCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List VisualForce page visibility",
	Long:                  "List VisualForce page visibility in permission sets",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listVisualforcePages(file)
		}
	},
}

func deleteVisualforceAccess(file string, pageName string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	err = p.DeleteVisualforcePageAccess(pageName)
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

func addVisualforceAccess(file string, pageName string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	err = p.AddVisualforcePageAccess(pageName)
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

func cloneVisualforcePageAccess(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	err = p.CloneVisualforcePageAccess(sourcePageName, pageName)
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

func listVisualforcePages(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	pages := p.GetVisualforcePageVisibility()
	for _, p := range pages {
		access := "disabled"
		if p.Enabled.ToBool() {
			access = "enabled"
		}
		fmt.Printf("%s: %s\n", p.ApexPage, access)
	}
}
