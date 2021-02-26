package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var sourceField string
var fieldName string

func init() {
	cloneCmd.Flags().StringVarP(&sourceField, "source", "s", "", "source field name")
	cloneCmd.Flags().StringVarP(&fieldName, "field", "f", "", "new field name")
	cloneCmd.MarkFlagRequired("source")
	cloneCmd.MarkFlagRequired("field")

	addFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	addFieldCmd.MarkFlagRequired("field")

	deleteFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	deleteFieldCmd.MarkFlagRequired("field")

	editFieldCmd.Flags().StringVarP(&fieldName, "field", "f", "", "field name")
	editFieldCmd.Flags().BoolP("edit", "e", false, "allow edit")
	editFieldCmd.Flags().BoolP("read", "r", false, "allow read")
	editFieldCmd.Flags().BoolP("no-edit", "E", false, "disallow edit")
	editFieldCmd.Flags().BoolP("no-read", "R", false, "disallow read")
	editFieldCmd.Flags().SortFlags = false
	editFieldCmd.MarkFlagRequired("field")

	FieldPermissionsCmd.AddCommand(cloneCmd)
	FieldPermissionsCmd.AddCommand(addFieldCmd)
	FieldPermissionsCmd.AddCommand(editFieldCmd)
	FieldPermissionsCmd.AddCommand(deleteFieldCmd)
}

var FieldPermissionsCmd = &cobra.Command{
	Use:   "field-permissions",
	Short: "Manage field permissions",
}

var cloneCmd = &cobra.Command{
	Use:   "clone -s SObject.Field -f SObject.Field [flags] [filename]...",
	Short: "Clone field permissions",
	Long:  "Clone field permissions in permission sets for a new field",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addNewField(file)
		}
	},
}

func addNewField(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
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

var editFieldCmd = &cobra.Command{
	Use:   "edit -f SObject.Field [flags] [filename]...",
	Short: "Update field permissions",
	Long:  "Update field permissions in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := fieldPermissionsToUpdate(cmd)
		for _, file := range args {
			updateFieldPermissions(file, perms)
		}
	},
}

var addFieldCmd = &cobra.Command{
	Use:   "add -f SObject.Field [flags] [filename]...",
	Short: "Add field permissions",
	Long:  "Add field permissions in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addFieldPermissions(file)
		}
	},
}

var deleteFieldCmd = &cobra.Command{
	Use:   "delete -f SObject.Field [flags] [filename]...",
	Short: "Delete field permissions",
	Long:  "Delete field permissions in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteFieldPermissions(file, fieldName)
		}
	},
}

func addFieldPermissions(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
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
	p, err := permissionset.Open(file)
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
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
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

func fieldPermissionsToUpdate(cmd *cobra.Command) permissionset.FieldPermissions {
	perms := permissionset.FieldPermissions{}
	perms.Editable = textValue(cmd, "edit")
	perms.Readable = textValue(cmd, "read")
	return perms
}
