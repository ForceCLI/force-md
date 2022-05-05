package profile

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var sourceRecordType string
var recordTypeName string

func init() {
	editRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type name")
	editRecordTypeCmd.Flags().BoolP("visible", "v", false, "visible")
	editRecordTypeCmd.Flags().BoolP("default", "d", false, "default")
	editRecordTypeCmd.Flags().BoolP("no-visible", "V", false, "not visible")
	editRecordTypeCmd.Flags().BoolP("no-default", "D", false, "not default")
	editRecordTypeCmd.Flags().SortFlags = false
	editRecordTypeCmd.MarkFlagRequired("recordtype")

	addRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type name")
	addRecordTypeCmd.MarkFlagRequired("recordtype")

	deleteRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type name")
	deleteRecordTypeCmd.MarkFlagRequired("recordtype")

	cloneRecordTypeCmd.Flags().StringVarP(&sourceRecordType, "source", "s", "", "source record type name")
	cloneRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "new record type name")
	cloneRecordTypeCmd.MarkFlagRequired("source")
	cloneRecordTypeCmd.MarkFlagRequired("recordtype")

	RecordTypeCmd.AddCommand(addRecordTypeCmd)
	RecordTypeCmd.AddCommand(editRecordTypeCmd)
	RecordTypeCmd.AddCommand(deleteRecordTypeCmd)
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
	RecordTypeCmd.AddCommand(cloneRecordTypeCmd)
}

var RecordTypeCmd = &cobra.Command{
	Use:   "recordtype",
	Short: "Manage record type visibility",
}

var cloneRecordTypeCmd = &cobra.Command{
	Use:                   "clone -s SObject.RecordType -f SObject.RecordType [filename]...",
	Short:                 "Clone record type visibility",
	Long:                  "Clone record type visibility in profiles for a new record type",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			cloneRecordType(file)
		}
	},
}

var editRecordTypeCmd = &cobra.Command{
	Use:   "edit -f SObject.RecordType [flags] [filename]...",
	Short: "Update record types",
	Long:  "Update record types in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := recordTypeVisibilityToUpdate(cmd)
		for _, file := range args {
			updateRecordTypeVisibility(file, perms)
		}
	},
}

var addRecordTypeCmd = &cobra.Command{
	Use:                   "add -f SObject.RecordType [filename]...",
	Short:                 "Add record type",
	Long:                  "Add record type to profiles",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addRecordTypeVisibility(file)
		}
	},
}

var deleteRecordTypeCmd = &cobra.Command{
	Use:                   "delete -f SObject.RecordType [filename]...",
	Short:                 "Delete record type",
	Long:                  "Delete record type visisiblity from profiles",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteRecordTypeVisibility(file, recordTypeName)
		}
	},
}

var listRecordTypesCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List record types",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRecordTypes(file)
		}
	},
}

func recordTypeVisibilityToUpdate(cmd *cobra.Command) profile.RecordTypeVisibility {
	perms := profile.RecordTypeVisibility{}
	perms.Visible = textValue(cmd, "visible")
	perms.Default = textValue(cmd, "default")
	return perms
}

func cloneRecordType(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.CloneRecordType(sourceRecordType, recordTypeName)
	if err != nil {
		log.Warn(fmt.Sprintf("clone failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func addRecordTypeVisibility(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.AddRecordType(recordTypeName)
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

func deleteRecordTypeVisibility(file string, recordTypeName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteRecordType(recordTypeName)
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

func updateRecordTypeVisibility(file string, perms profile.RecordTypeVisibility) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.SetRecordTypeVisibility(recordTypeName, perms)
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

func listRecordTypes(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	recordTypes := p.GetRecordTypeVisibility()
	for _, a := range recordTypes {
		var perms []string
		if a.Visible.Text == "true" {
			perms = append(perms, "visible")
		} else {
			perms = append(perms, "not visible")
		}
		if a.Default.Text == "true" {
			perms = append(perms, "(default)")
		}
		permsString := "no access"
		if len(perms) > 0 {
			permsString = strings.Join(perms, " ")
		}
		fmt.Printf("%s: %s\n", a.RecordType, permsString)
	}
}