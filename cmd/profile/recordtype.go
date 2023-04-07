package profile

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/profile"
)

var sourceRecordType string

func init() {
	editRecordTypeCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type name")
	editRecordTypeCmd.Flags().BoolP("visible", "v", false, "visible")
	editRecordTypeCmd.Flags().BoolP("default", "d", false, "default")
	editRecordTypeCmd.Flags().BoolP("no-visible", "V", false, "not visible")
	editRecordTypeCmd.Flags().BoolP("no-default", "D", false, "not default")
	editRecordTypeCmd.Flags().SortFlags = false
	editRecordTypeCmd.MarkFlagRequired("recordtype")

	addRecordTypeCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type name")
	addRecordTypeCmd.MarkFlagRequired("recordtype")

	deleteRecordTypeCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type name")
	deleteRecordTypeCmd.MarkFlagRequired("recordtype")

	cloneRecordTypeCmd.Flags().StringVarP(&sourceRecordType, "source", "s", "", "source record type name")
	cloneRecordTypeCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "new record type name")
	cloneRecordTypeCmd.MarkFlagRequired("source")
	cloneRecordTypeCmd.MarkFlagRequired("recordtype")

	tableRecordTypeCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	tableRecordTypeCmd.Flags().StringVarP(&recordType, "recordtype", "r", "", "record type name")
	tableRecordTypeCmd.Flags().BoolP("visible", "v", false, "visible")
	tableRecordTypeCmd.Flags().BoolP("default", "d", false, "default")
	tableRecordTypeCmd.Flags().BoolP("no-visible", "V", false, "not visible")
	tableRecordTypeCmd.Flags().BoolP("no-default", "D", false, "not default")
	tableRecordTypeCmd.Flags().SortFlags = false

	RecordTypeCmd.AddCommand(addRecordTypeCmd)
	RecordTypeCmd.AddCommand(editRecordTypeCmd)
	RecordTypeCmd.AddCommand(deleteRecordTypeCmd)
	RecordTypeCmd.AddCommand(listRecordTypesCmd)
	RecordTypeCmd.AddCommand(cloneRecordTypeCmd)
	RecordTypeCmd.AddCommand(tableRecordTypeCmd)
}

var RecordTypeCmd = &cobra.Command{
	Use:   "recordtype",
	Short: "Manage record type visibility",
}

var cloneRecordTypeCmd = &cobra.Command{
	Use:                   "clone -s SObject.RecordType -r SObject.RecordType [filename]...",
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
	Use:   "edit -r SObject.RecordType [flags] [filename]...",
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
	Use:                   "add -r SObject.RecordType [filename]...",
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
	Use:                   "delete -r SObject.RecordType [filename]...",
	Short:                 "Delete record type",
	Long:                  "Delete record type visisiblity from profiles",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteRecordTypeVisibility(file, recordType)
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

var tableRecordTypeCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Record Types in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := recordTypeVisibilityToUpdate(cmd)
		tableRecordTypes(args, perms)
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
	err = p.CloneRecordType(sourceRecordType, recordType)
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
	err = p.AddRecordType(recordType)
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

func deleteRecordTypeVisibility(file string, recordType string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteRecordType(recordType)
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
	err = p.SetRecordTypeVisibility(recordType, perms)
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

func tableRecordTypes(files []string, filter profile.RecordTypeVisibility) {
	var filters []profile.RecordTypeFilter
	if objectName != "" {
		filters = append(filters, func(f profile.RecordTypeVisibility) bool {
			pieces := strings.Split(f.RecordType, ".")
			if len(pieces) != 2 {
				return false
			}
			return strings.ToLower(pieces[0]) == strings.ToLower(objectName)
		})
	}
	if recordType != "" {
		fullRecordTypeName := strings.ToLower(recordType)
		if objectName != "" {
			recordType = strings.TrimPrefix(strings.ToLower(recordType), strings.ToLower(objectName)+".")
			fullRecordTypeName = objectName + "." + recordType
		}
		filters = append(filters, func(f profile.RecordTypeVisibility) bool {
			return strings.ToLower(f.RecordType) == fullRecordTypeName
		})
	}
	flagFilter := func(o profile.RecordTypeVisibility) bool {
		if filter.Visible.Text != "" && filter.Visible.ToBool() != o.Visible.ToBool() {
			return false
		}
		if filter.Default.Text != "" && filter.Default.ToBool() != o.Default.ToBool() {
			return false
		}
		return true
	}
	filters = append(filters, flagFilter)
	type visibility struct {
		recordTypes profile.RecordTypeVisibilityList
		profile     string
	}
	var visibilities []visibility
	for _, file := range files {
		p, err := profile.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		profileName := strings.TrimSuffix(path.Base(file), ".profile")
		visibilities = append(visibilities, visibility{recordTypes: p.GetRecordTypeVisibility(filters...), profile: profileName})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "RecordType", "Default", "Visible", "PersonAccountDefault"})
	table.SetRowLine(true)
	for _, vis := range visibilities {
		for _, o := range vis.recordTypes {
			table.Append([]string{vis.profile, o.RecordType,
				o.Default.Text,
				o.Visible.Text,
				strconv.FormatBool(o.PersonAccountDefault.ToBool()),
			})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
