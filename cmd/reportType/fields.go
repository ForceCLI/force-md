package reportType

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/reportType"
)

var (
	fieldName string
	tableName string
	section   string
)

func init() {
	addFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	addFieldCmd.Flags().StringVarP(&tableName, "table", "t", "", "table name")
	addFieldCmd.Flags().StringVarP(&section, "section", "s", "", "section")
	addFieldCmd.MarkFlagRequired("field")
	addFieldCmd.MarkFlagRequired("table")
	addFieldCmd.MarkFlagRequired("section")

	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	FieldCmd.AddCommand(listFieldsCmd)
	FieldCmd.AddCommand(addFieldCmd)
	FieldCmd.AddCommand(deleteFieldCmd)
}

var FieldCmd = &cobra.Command{
	Use:                   "fields",
	Short:                 "Manage report type fields",
	DisableFlagsInUseLine: true,
}

var listFieldsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List report type fields",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listFields(file)
		}
	},
}

var addFieldCmd = &cobra.Command{
	Use:                   "add -s Section -t Table -f Field [filename]...",
	Short:                 "Add field to report type",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addField(file, section, tableName, fieldName)
		}
	},
}

var deleteFieldCmd = &cobra.Command{
	Use:                   "delete -f Field [filename]...",
	Short:                 "Delete field from report type",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteField(file, fieldName)
		}
	},
}

func listFields(file string) {
	o, err := reportType.Open(file)
	if err != nil {
		log.Warn("parsing report type failed: " + err.Error())
		return
	}
	fields := o.GetFields()
	for _, f := range fields {
		fmt.Printf("%s.%s\n", f.Table, f.Field)
	}
}

func addField(file string, section, tableName, fieldName string) {
	o, err := reportType.Open(file)
	if err != nil {
		log.Warn("parsing report type failed: " + err.Error())
		return
	}
	err = o.AddField(section, tableName, fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func deleteField(file string, fieldName string) {
	o, err := reportType.Open(file)
	if err != nil {
		log.Warn("parsing report type failed: " + err.Error())
		return
	}
	err = o.DeleteField(fieldName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(o, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
