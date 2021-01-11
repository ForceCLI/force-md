package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var applicationName string

func init() {
	deleteApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	deleteApplicationCmd.MarkFlagRequired("application")

	ApplicationCmd.AddCommand(deleteApplicationCmd)
}

var ApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manage application visibility",
}

var deleteApplicationCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete application visibility",
	Long:  "Delete application visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteApplicationVisibility(file, applicationName)
		}
	},
}

func deleteApplicationVisibility(file string, applicationName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteApplicationVisibility(applicationName)
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
