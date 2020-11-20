package profile

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/profile"
)

var (
	objectName string
)

func init() {
	ObjectPermissionsCmd.Flags().StringVarP(&objectName, "object", "o", "", "object name")
	ObjectPermissionsCmd.Flags().BoolP("create", "c", false, "allow create")
	ObjectPermissionsCmd.Flags().BoolP("delete", "d", false, "allow delete")
	ObjectPermissionsCmd.Flags().BoolP("edit", "e", false, "allow edit")
	ObjectPermissionsCmd.Flags().BoolP("read", "r", false, "allow read")
	ObjectPermissionsCmd.Flags().BoolP("modify-all", "m", false, "allow modify all")
	ObjectPermissionsCmd.Flags().BoolP("view-all", "v", false, "allow view all")
	ObjectPermissionsCmd.Flags().BoolP("no-create", "C", false, "disallow create")
	ObjectPermissionsCmd.Flags().BoolP("no-delete", "D", false, "disallow delete")
	ObjectPermissionsCmd.Flags().BoolP("no-edit", "E", false, "disallow edit")
	ObjectPermissionsCmd.Flags().BoolP("no-read", "R", false, "disallow read")
	ObjectPermissionsCmd.Flags().BoolP("no-modify-all", "M", false, "disallow modify all")
	ObjectPermissionsCmd.Flags().BoolP("no-view-all", "V", false, "disallow view all")
	ObjectPermissionsCmd.Flags().SortFlags = false
	ObjectPermissionsCmd.MarkFlagRequired("object")
}

var ObjectPermissionsCmd = &cobra.Command{
	Use:   "object-permissions",
	Short: "Update object permissions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := permissionsToUpdate(cmd)
		for _, file := range args {
			updateObjectPermissions(file, perms)
		}
	},
}

func textValue(cmd *cobra.Command, flag string) (t profile.BooleanText) {
	if cmd.Flags().Changed(flag) {
		val, _ := cmd.Flags().GetBool(flag)
		t = profile.BooleanText{
			Text: strconv.FormatBool(val),
		}
	}
	antiFlag := "no-" + flag
	if cmd.Flags().Changed(antiFlag) {
		val, _ := cmd.Flags().GetBool(antiFlag)
		t = profile.BooleanText{
			Text: strconv.FormatBool(!val),
		}
	}
	return t
}

func permissionsToUpdate(cmd *cobra.Command) profile.ObjectPermissions {
	perms := profile.ObjectPermissions{}
	perms.AllowCreate = textValue(cmd, "create")
	perms.AllowDelete = textValue(cmd, "delete")
	perms.AllowEdit = textValue(cmd, "edit")
	perms.AllowRead = textValue(cmd, "read")
	perms.ModifyAllRecords = textValue(cmd, "modify-all")
	perms.ViewAllRecords = textValue(cmd, "view-all")
	return perms
}

func updateObjectPermissions(file string, perms profile.ObjectPermissions) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
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
