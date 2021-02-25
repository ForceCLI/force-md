package permissionset

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var (
	objectName string
)

func init() {
	editObjectCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	editObjectCmd.Flags().BoolP("create", "c", false, "allow create")
	editObjectCmd.Flags().BoolP("delete", "d", false, "allow delete")
	editObjectCmd.Flags().BoolP("edit", "e", false, "allow edit")
	editObjectCmd.Flags().BoolP("read", "r", false, "allow read")
	editObjectCmd.Flags().BoolP("modify-all", "m", false, "allow modify all")
	editObjectCmd.Flags().BoolP("view-all", "v", false, "allow view all")
	editObjectCmd.Flags().BoolP("no-create", "C", false, "disallow create")
	editObjectCmd.Flags().BoolP("no-delete", "D", false, "disallow delete")
	editObjectCmd.Flags().BoolP("no-edit", "E", false, "disallow edit")
	editObjectCmd.Flags().BoolP("no-read", "R", false, "disallow read")
	editObjectCmd.Flags().BoolP("no-modify-all", "M", false, "disallow modify all")
	editObjectCmd.Flags().BoolP("no-view-all", "V", false, "disallow view all")
	editObjectCmd.Flags().SortFlags = false
	editObjectCmd.MarkFlagRequired("object")

	addObjectCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	addObjectCmd.MarkFlagRequired("object")

	deleteObjectCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	deleteObjectCmd.MarkFlagRequired("object")

	ObjectPermissionsCmd.AddCommand(editObjectCmd)
	ObjectPermissionsCmd.AddCommand(addObjectCmd)
	ObjectPermissionsCmd.AddCommand(deleteObjectCmd)
}

var ObjectPermissionsCmd = &cobra.Command{
	Use:   "object-permissions",
	Short: "Update object permissions",
}

var editObjectCmd = &cobra.Command{
	Use:   "edit -o SObject [flags] [filename]...",
	Short: "Update object permissions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := objectPermissionsToUpdate(cmd)
		for _, file := range args {
			updateObjectPermissions(file, perms)
		}
	},
}

var addObjectCmd = &cobra.Command{
	Use:   "add -o SObject [flags] [filename]...",
	Short: "Add object permissions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addObjectPermissions(file)
		}
	},
}

var deleteObjectCmd = &cobra.Command{
	Use:   "delete -o SObject [flags] [filename]...",
	Short: "Delete object permissions",
	Long:  "Delete object permissions and related field permissions in permission sets",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteObjectPermissions(file, objectName)
		}
	},
}

func textValue(cmd *cobra.Command, flag string) (t permissionset.BooleanText) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetBool(flag)
		t = permissionset.BooleanText{
			Text: strconv.FormatBool(val),
		}
	}
	antiFlag := "no-" + flag
	if cmd.Flags().Changed(antiFlag) {
		val, _ := cmd.Flags().GetBool(antiFlag)
		t = permissionset.BooleanText{
			Text: strconv.FormatBool(!val),
		}
	}
	return t
}

func objectPermissionsToUpdate(cmd *cobra.Command) permissionset.ObjectPermissions {
	perms := permissionset.ObjectPermissions{}
	perms.AllowCreate = textValue(cmd, "create")
	perms.AllowDelete = textValue(cmd, "delete")
	perms.AllowEdit = textValue(cmd, "edit")
	perms.AllowRead = textValue(cmd, "read")
	perms.ModifyAllRecords = textValue(cmd, "modify-all")
	perms.ViewAllRecords = textValue(cmd, "view-all")
	return perms
}

func updateObjectPermissions(file string, perms permissionset.ObjectPermissions) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.SetObjectPermissions(objectName, perms)
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

func addObjectPermissions(file string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	err = p.AddObjectPermissions(objectName)
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

func deleteObjectPermissions(file string, objectName string) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permission set failed: " + err.Error())
		return
	}
	p.DeleteObjectPermissions(objectName)
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
