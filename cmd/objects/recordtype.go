package objects

import (
	"fmt"
	"html"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects"
	rt "github.com/ForceCLI/force-md/objects/recordtype"
)

var (
	recordTypeName string
	picklistValue  string
)

func init() {
	deleteRecordTypeCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")

	writeRecordTypesCmd.Flags().StringP("directory", "d", "", "directory where record types should be output")
	writeRecordTypesCmd.MarkFlagRequired("directory")

	tableRecordTypesCmd.Flags().BoolP("active", "a", false, "is active")
	tableRecordTypesCmd.Flags().BoolP("no-active", "A", false, "is not active")

	listRecordTypesCmd.Flags().BoolP("active", "a", false, "is active")
	listRecordTypesCmd.Flags().BoolP("no-active", "A", false, "is not active")

	RecordTypeCmd.AddCommand(deleteRecordTypeCmd)
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
	RecordTypeCmd.AddCommand(tableRecordTypesCmd)
	RecordTypeCmd.AddCommand(recordtypePicklistCmd)
	RecordTypeCmd.AddCommand(writeRecordTypesCmd)

	recordtypePicklistTableCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	recordtypePicklistTableCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")

	recordtypePicklistListCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	recordtypePicklistListCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")

	recordtypePicklistListCmd.MarkFlagRequired("field")
	recordtypePicklistListCmd.MarkFlagRequired("recordtype")

	recordtypePicklistAddCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	recordtypePicklistAddCmd.Flags().StringVarP(&recordTypeName, "recordtype", "r", "", "record type")
	recordtypePicklistAddCmd.Flags().StringVarP(&picklistValue, "value", "v", "", "picklist value")

	recordtypePicklistAddCmd.MarkFlagRequired("field")
	recordtypePicklistAddCmd.MarkFlagRequired("recordtype")
	recordtypePicklistAddCmd.MarkFlagRequired("value")

	recordtypePicklistCmd.AddCommand(recordtypePicklistTableCmd)
	recordtypePicklistCmd.AddCommand(recordtypePicklistListCmd)
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

var recordtypePicklistListCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List record type picklist options",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listRecordTypePicklistOptions(file, recordTypeName)
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
		perms := recordTypeVisibilityFromFlags(cmd)
		for _, file := range args {
			listRecordType(file, perms)
		}
	},
}

var tableRecordTypesCmd = &cobra.Command{
	Use:                   "table [filename]...",
	Short:                 "List object record types in a table",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := recordTypeVisibilityFromFlags(cmd)
		tableRecordType(args, perms)
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

var writeRecordTypesCmd = &cobra.Command{
	Use:                   "write -d directory [filename]...",
	Short:                 "Split object record types into separate files",
	Long:                  "Split object record types into separate metadata files to match sfdx's source format",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("directory")
		for _, file := range args {
			writeRecordTypes(file, dir)
		}
	},
}

func listRecordType(file string, filter rt.RecordType) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	var filters []rt.RecordTypeFilter
	filters = append(filters, func(r rt.RecordType) bool {
		if filter.Active.Text != "" && filter.Active.ToBool() != r.Active.ToBool() {
			return false
		}
		return true
	})
	recordTypes := o.GetRecordTypes(filters...)
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
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
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
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
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

func tableRecordType(files []string, filter rt.RecordType) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Object", "Record Type", "Active"})
	table.SetAutoMergeCells(true)
	table.SetAutoMergeCellsByColumnIndex([]int{0, 1})
	table.SetRowLine(true)
	var filters []rt.RecordTypeFilter
	filters = append(filters, func(r rt.RecordType) bool {
		if filter.Active.Text != "" && filter.Active.ToBool() != r.Active.ToBool() {
			return false
		}
		return true
	})
	for _, file := range files {
		o, err := objects.Open(file)
		if err != nil {
			log.Warn("parsing object failed: " + err.Error())
			return
		}
		objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
		recordTypes := o.GetRecordTypes(filters...)

		for _, r := range recordTypes {
			table.Append([]string{objectName, r.FullName, r.Active.String()})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}

func tableRecordTypePicklistOptions(file string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
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

func listRecordTypePicklistOptions(file string, recordTypeName string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	objectName := internal.TrimSuffixToEnd(path.Base(file), ".object")
	recordTypes := o.GetRecordTypes()

	fieldName = strings.TrimPrefix(fieldName, objectName+".")
	recordTypeName = strings.TrimPrefix(recordTypeName, objectName+".")

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
					fmt.Println(html.UnescapeString(s))
				} else {
					panic(err.Error())
				}
			}
		}
	}
}

func writeRecordTypes(file string, recordTypesDir string) {
	o, err := objects.Open(file)
	if err != nil {
		log.Warn("parsing object failed: " + err.Error())
		return
	}
	recordTypes := o.GetRecordTypes()
	for _, f := range recordTypes {
		recordType := rt.RecordTypeMetadata{
			RecordType: f,
			Xmlns:      o.Xmlns,
		}
		err = internal.WriteToFile(recordType, recordTypesDir+"/"+f.FullName+".recordType-meta.xml")
		if err != nil {
			log.Warn("write failed: " + err.Error())
			return
		}
	}
}

func recordTypeVisibilityFromFlags(cmd *cobra.Command) rt.RecordType {
	perms := rt.RecordType{}
	perms.Active = textValue(cmd, "active")
	return perms
}

func textValue(cmd *cobra.Command, flag string) (t BooleanText) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetBool(flag)
		t = BooleanText{
			Text: strconv.FormatBool(val),
		}
	}
	antiFlag := "no-" + flag
	if cmd.Flags().Changed(antiFlag) {
		val, _ := cmd.Flags().GetBool(antiFlag)
		t = BooleanText{
			Text: strconv.FormatBool(!val),
		}
	}
	return t
}
