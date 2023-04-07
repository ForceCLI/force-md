package objects

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects"
)

var (
	recordTypeName string
	picklistValue  string
)

func init() {
	deleteRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")

	RecordTypeCmd.AddCommand(deleteRecordTypeCmd)
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
	RecordTypeCmd.AddCommand(recordtypePicklistCmd)

	recordtypePicklistTableCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	recordtypePicklistTableCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")

	recordtypePicklistAddCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	recordtypePicklistAddCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")
	recordtypePicklistAddCmd.Flags().StringVarP(&picklistValue, "value", "v", "", "picklist value")

	recordtypePicklistAddCmd.MarkFlagRequired("field")
	recordtypePicklistAddCmd.MarkFlagRequired("recordtype")
	recordtypePicklistAddCmd.MarkFlagRequired("value")

	recordtypePicklistCmd.AddCommand(recordtypePicklistTableCmd)
	recordtypePicklistCmd.AddCommand(recordtypePicklistAddCmd)
}

var RecordTypeCmd = &cobra.Command{
	Use:                   "recordtype",
	Short:                 "Manage object record type metadata",
	DisableFlagsInUseLine: true,
}

var recordtypePicklistCmd = &cobra.Command{
	Use:                   "picklist",
	Short:                 "Manage record type picklist options",
	DisableFlagsInUseLine: true,
}

var recordtypePicklistTableCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "Display record type picklist options",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			tableRecordTypePicklistOptions(file)
		}
	},
}

var recordtypePicklistAddCmd = &cobra.Command{
	Use:   "add [flags] [filename]...",
	Short: "Assign picklist value to record type",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			assignPicklistValueToRecordType(file)
		}
	},
}

var listRecordTypesCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List object record types",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRecordType(file)
		}
	},
}

var deleteRecordTypeCmd = &cobra.Command{
	Use:   "delete [flags] [filename]...",
	Short: "Delete object record type",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteRecordType(file)
		}
	},
}

func listRecordType(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	recordTypes := o.GetRecordTypes()
	for _, f := range recordTypes {
		fmt.Printf("%s.%s\n", objectName, f.FullName)
	}
}

func deleteRecordType(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	err = o.DeleteRecordType(strings.TrimPrefix(recordTypeName, objectName+"."))
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

func assignPicklistValueToRecordType(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	field := strings.TrimPrefix(fieldName, objectName+".")
	recordType := strings.TrimPrefix(recordTypeName, objectName+".")

	err = o.AddFieldPicklistValue(field, recordType, picklistValue)
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

func tableRecordTypePicklistOptions(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := strings.TrimSuffix(path.Base(file), ".object")
	recordTypes := o.GetRecordTypes()

	fieldName = strings.TrimPrefix(fieldName, objectName+".")
	recordTypeName = strings.TrimPrefix(recordTypeName, objectName+".")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Record Type", "Field", "Value", "Default"})
	table.SetAutoMergeCells(true)
	table.SetAutoMergeCellsByColumnIndex([]int{0, 1})
	table.SetRowLine(true)
	for _, r := range recordTypes {
		if recordTypeName != "" && strings.ToLower(r.FullName) != strings.ToLower(recordTypeName) {
			continue
		}
		for _, p := range r.PicklistValues {
			if fieldName != "" && strings.ToLower(p.Picklist) != strings.ToLower(fieldName) {
				continue
			}
			for _, v := range p.Values {
				if s, err := url.QueryUnescape(v.FullName); err == nil {
					table.Append([]string{r.FullName, objectName + "." + p.Picklist, s, v.Default.Text})
				} else {
					panic(err.Error())
				}
			}
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
