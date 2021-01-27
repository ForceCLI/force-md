package profile

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var applicationName string
var isDefault bool

func init() {
	deleteApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	deleteApplicationCmd.MarkFlagRequired("application")

	addApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	addApplicationCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "make default application")
	addApplicationCmd.MarkFlagRequired("application")

	ApplicationCmd.AddCommand(addApplicationCmd)
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

var addApplicationCmd = &cobra.Command{
	Use:   "add",
	Short: "Add application visibility",
	Long:  "Add application visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addApplicationVisibility(file, applicationName, isDefault)
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

func addApplicationVisibility(file string, applicationName string, isDefault bool) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.AddApplicationVisibility(applicationName, isDefault)
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
