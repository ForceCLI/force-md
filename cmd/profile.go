package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/profile"
)

var sourceField string
var newField string

func init() {
	cloneCmd.Flags().StringVarP(&sourceField, "source", "s", "", "source field name")
	cloneCmd.Flags().StringVarP(&newField, "field", "f", "", "new field name")
	cloneCmd.MarkFlagRequired("source")
	cloneCmd.MarkFlagRequired("field")
	profileCmd.AddCommand(fieldPermissionsCmd)
	fieldPermissionsCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(profileCmd)

	profileCmd.AddCommand(tidyCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage Profiles",
}

var fieldPermissionsCmd = &cobra.Command{
	Use:   "field-permissions",
	Short: "Manage field permissions",
}

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Tidy profile metadata",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tidy(file)
		}
	},
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone field permissions",
	Long:  "Clone field permissions in profiles for a new field",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addNewField(file)
		}
	},
}

func tidy(file string) {
	p, err := profile.ParseProfile(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	p.Tidy()
	err = p.Write(file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func addNewField(file string) {
	p, err := profile.ParseProfile(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.CloneFieldPermissions(sourceField, newField)
	if err != nil {
		log.Warn(fmt.Sprintf("clone failed for %s: %s", file, err.Error()))
		return
	}
	err = p.Write(file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
