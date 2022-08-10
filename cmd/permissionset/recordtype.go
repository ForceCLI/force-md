package permissionset

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var sourceRecordType string
var recordTypeName string

func init() {
	addRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type name")
	addRecordTypeCmd.MarkFlagRequired("recordtype")

	deleteRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type name")
	deleteRecordTypeCmd.MarkFlagRequired("recordtype")

	RecordTypeCmd.AddCommand(addRecordTypeCmd)
	RecordTypeCmd.AddCommand(deleteRecordTypeCmd)
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
}

var RecordTypeCmd = &cobra.Command{
	Use:   "recordtype",
	Short: "Manage record type visibility",
}

var addRecordTypeCmd = &cobra.Command{
	Use:                   "add -r SObject.RecordType [filename]...",
	Short:                 "Add record type",
	Long:                  "Add record type to permission sets",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addRecordTypeVisibility(file)
		}
	},
}

var deleteRecordTypeCmd = &cobra.Command{
	Use:                   "delete -r SObject.RecordType [filename]...",
	Short:                 "Delete record type",
	Long:                  "Delete record type visisiblity from permission sets",
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

func addRecordTypeVisibility(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
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
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission sets failed: " + err.Error())
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

func listRecordTypes(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission sets failed: " + err.Error())
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
		permsString := "no access"
		if len(perms) > 0 {
			permsString = strings.Join(perms, " ")
		}
		fmt.Printf("%s: %s\n", a.RecordType, permsString)
	}
}
