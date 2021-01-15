package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var pageName string

func init() {
	deleteVisualforcePageCmd.Flags().StringVarP(&pageName, "page", "p", "", "visualforce page name")
	deleteVisualforcePageCmd.MarkFlagRequired("page")

	VisualforceCmd.AddCommand(deleteVisualforcePageCmd)
}

var VisualforceCmd = &cobra.Command{
	Use:   "visualforce",
	Short: "Manage visualforce page visibility",
}

var deleteVisualforcePageCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete visualforce page access",
	Long:  "Delete visualforce page access in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteVisualforceAccess(file, pageName)
		}
	},
}

func deleteVisualforceAccess(file string, pageName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
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
