package profile

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/permissionset"
	"github.com/ForceCLI/force-md/profile"
)

var sourceField string
var fieldName string

func init() {
	editFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	editFieldCmd.Flags().BoolP("edit", "e", false, "allow edit")
	editFieldCmd.Flags().BoolP("read", "r", false, "allow read")
	editFieldCmd.Flags().BoolP("no-edit", "E", false, "disallow edit")
	editFieldCmd.Flags().BoolP("no-read", "R", false, "disallow read")
	editFieldCmd.Flags().SortFlags = false
	editFieldCmd.MarkFlagRequired("field")

	addFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	addFieldCmd.Flags().BoolP("edit", "e", false, "allow edit")
	addFieldCmd.Flags().BoolP("read", "r", false, "allow read")
	addFieldCmd.MarkFlagRequired("field")

	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	cloneCmd.Flags().StringVarP(&sourceField, "source", "s", "", "source field name")
	cloneCmd.Flags().StringVarP(&fieldName, "field", "f", "", "new field name")
	cloneCmd.MarkFlagRequired("source")
	cloneCmd.MarkFlagRequired("field")

	tableFieldsCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	tableFieldsCmd.Flags().StringVarP(&objectName, "object", "o", "", "object")
	tableFieldsCmd.Flags().BoolP("edit", "e", false, "allow edit")
	tableFieldsCmd.Flags().BoolP("read", "r", false, "allow read")
	tableFieldsCmd.Flags().BoolP("no-edit", "E", false, "disallow edit")
	tableFieldsCmd.Flags().BoolP("no-read", "R", false, "disallow read")

	listFieldsCmd.Flags().StringVarP(&objectName, "object", "o", "", "object")
	listFieldsCmd.Flags().BoolP("edit", "e", false, "allow edit")
	listFieldsCmd.Flags().BoolP("read", "r", false, "allow read")
	listFieldsCmd.Flags().BoolP("no-edit", "E", false, "disallow edit")
	listFieldsCmd.Flags().BoolP("no-read", "R", false, "disallow read")

	FieldPermissionsCmd.AddCommand(addFieldCmd)
	FieldPermissionsCmd.AddCommand(editFieldCmd)
	FieldPermissionsCmd.AddCommand(deleteFieldCmd)
	FieldPermissionsCmd.AddCommand(listFieldsCmd)
	FieldPermissionsCmd.AddCommand(cloneCmd)
	FieldPermissionsCmd.AddCommand(tableFieldsCmd)
}

var FieldPermissionsCmd = &cobra.Command{
	Use:   "field-permissions",
	Short: "Manage field permissions",
}

var cloneCmd = &cobra.Command{
	Use:   "clone -s SObject.Field -f SObject.Field [flags] [filename]...",
	Short: "Clone field permissions",
	Long:  "Clone field permissions in profiles for a new field",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			cloneField(file)
		}
	},
}

var editFieldCmd = &cobra.Command{
	Use:   "edit -f SObject.Field [flags] [filename]...",
	Short: "Update field permissions",
	Long:  "Update field permissions in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := fieldPermissionsFromFlags(cmd)
		for _, file := range args {
			updateFieldPermissions(file, perms)
		}
	},
}

var addFieldCmd = &cobra.Command{
	Use:   "add -f SObject.Field [flags] [filename]...",
	Short: "Add field permissions",
	Long:  "Add field permissions in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := fieldPermissionsFromFlags(cmd)
		for _, file := range args {
			addFieldPermissions(file)
			updateFieldPermissions(file, perms)
		}
	},
}

var deleteFieldCmd = &cobra.Command{
	Use:   "delete -f SObject.Field [flags] [filename]...",
	Short: "Delete field permissions",
	Long:  "Delete field permissions in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteFieldPermissions(file, fieldName)
		}
	},
}

var listFieldsCmd = &cobra.Command{
	Use:                   "list [filename]...",
	Short:                 "List fields",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		perms := fieldPermissionsFromFlags(cmd)
		for _, file := range args {
			listFields(file, perms)
		}
	},
}

var tableFieldsCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Field Permissions in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := fieldPermissionsFromFlags(cmd)
		tableFieldPermissions(args, perms)
	},
}

func fieldPermissionsFromFlags(cmd *cobra.Command) permissionset.FieldPermissions {
	perms := permissionset.FieldPermissions{}
	perms.Editable = textValue(cmd, "edit")
	perms.Readable = textValue(cmd, "read")
	return perms
}

func cloneField(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.CloneFieldPermissions(sourceField, fieldName)
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

func addFieldPermissions(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.AddFieldPermissions(fieldName)
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

func deleteFieldPermissions(file string, fieldName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteFieldPermissions(fieldName)
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

func updateFieldPermissions(file string, perms permissionset.FieldPermissions) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.SetFieldPermissions(fieldName, perms)
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

func listFields(file string, toApply permissionset.FieldPermissions) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	var filters []profile.FieldFilter
	if fieldName != "" && !strings.ContainsRune(fieldName, '.') && objectName != "" {
		fieldName = objectName + "." + fieldName
	}
	if fieldName != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return strings.ToLower(f.Field) == strings.ToLower(fieldName)
		})
	}
	if objectName != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return strings.HasPrefix(strings.ToLower(f.Field), strings.ToLower(objectName+"."))
		})
	}
	if toApply.Readable.Text != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return toApply.Readable.ToBool() == f.Readable.ToBool()
		})
	}
	if toApply.Editable.Text != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return toApply.Editable.ToBool() == f.Editable.ToBool()
		})
	}
	fields := p.GetFieldPermissions(filters...)
	for _, a := range fields {
		var perms []string
		if a.Readable.Text == "true" {
			perms = append(perms, "read")
		}
		if a.Editable.Text == "true" {
			perms = append(perms, "write")
		}
		permsString := "no access"
		if len(perms) > 0 {
			permsString = strings.Join(perms, "-")
		}
		fmt.Printf("%s: %s\n", a.Field, permsString)
	}
}

func tableFieldPermissions(files []string, toApply permissionset.FieldPermissions) {
	var filters []profile.FieldFilter
	if fieldName != "" && !strings.ContainsRune(fieldName, '.') && objectName != "" {
		fieldName = objectName + "." + fieldName
	}
	if fieldName != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return strings.ToLower(f.Field) == strings.ToLower(fieldName)
		})
	}
	if objectName != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return strings.HasPrefix(strings.ToLower(f.Field), strings.ToLower(objectName+"."))
		})
	}
	if toApply.Readable.Text != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return toApply.Readable.ToBool() == f.Readable.ToBool()
		})
	}
	if toApply.Editable.Text != "" {
		filters = append(filters, func(f permissionset.FieldPermissions) bool {
			return toApply.Editable.ToBool() == f.Editable.ToBool()
		})
	}
	type perm struct {
		fields  permissionset.FieldPermissionsList
		profile string
	}
	var perms []perm
	for _, file := range files {
		p, err := profile.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		profileName := internal.TrimSuffixToEnd(path.Base(file), ".profile")
		perms = append(perms, perm{fields: p.GetFieldPermissions(filters...), profile: profileName})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "Field", "Readable", "Editable"})
	table.SetRowLine(true)
	for _, perm := range perms {
		for _, f := range perm.fields {
			table.Append([]string{perm.profile, f.Field, f.Readable.Text, f.Editable.Text})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}
