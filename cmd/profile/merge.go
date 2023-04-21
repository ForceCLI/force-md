package profile

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/profile"
)

var (
	sourceFileName string
)

func init() {
	MergeCmd.Flags().StringVarP(&sourceFileName, "source", "s", "", "source profile")
	MergeCmd.MarkFlagRequired("source")
}

var MergeCmd = &cobra.Command{
	Use:                   "merge -s path/to/Source.profile [filename]...",
	Short:                 "Merge profiles",
	Long:                  "Apply permissions granted in source profile",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		apply, err := profile.Open(sourceFileName)
		if err != nil {
			log.Fatal("loading source profile failed: " + err.Error())
		}
		for _, file := range args {
			mergePermissions(file, *apply)
		}
	},
}

func mergePermissions(file string, apply profile.Profile) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	for _, a := range apply.ApplicationVisibilities {
		err = p.AddApplicationVisibility(a.Application, a.Default.ToBool())
		if err != nil && err != profile.ApplicationExistsError {
			log.Warn(fmt.Sprintf("adding application %s permissions failed for %s: %s", a.Application, file, err.Error()))
			return
		}
	}
	for _, c := range apply.ClassAccesses {
		err = p.AddClass(c.ApexClass)
		if err != nil && err != profile.ClassExistsError {
			log.Warn(fmt.Sprintf("adding apex class %s permissions failed for %s: %s", c.ApexClass, file, err.Error()))
			return
		}
	}
	for _, o := range apply.ObjectPermissions {
		objectName := o.Object
		err = p.AddObjectPermissions(objectName)
		if err != nil && err != profile.ObjectExistsError {
			log.Warn(fmt.Sprintf("adding object %s permissions failed for %s: %s", objectName, file, err.Error()))
			return
		}
		err = p.SetObjectPermissions(objectName, o)
		if err != nil {
			log.Warn(fmt.Sprintf("updating object %s permissions failed for %s: %s", objectName, file, err.Error()))
			return
		}
	}
	for _, f := range apply.FieldPermissions {
		fieldName := f.Field
		err = p.AddFieldPermissions(fieldName)
		if err != nil && err != profile.FieldExistsError {
			log.Warn(fmt.Sprintf("adding field %s permissions failed for %s: %s", fieldName, file, err.Error()))
			return
		}
		err = p.SetFieldPermissions(fieldName, f)
		if err != nil {
			log.Warn(fmt.Sprintf("updating field %s permissions failed for %s: %s", fieldName, file, err.Error()))
			return
		}
	}
	for _, v := range apply.PageAccesses {
		err = p.AddVisualforcePageAccess(v.ApexPage)
		if err != nil && err != profile.VisualforcePageExistsError {
			log.Warn(fmt.Sprintf("adding visualforce page %s failed for %s: %s", v.ApexPage, file, err.Error()))
			return
		}
	}
	for _, r := range apply.RecordTypeVisibilities {
		err = p.AddRecordType(r.RecordType)
		if err != nil && err != profile.RecordTypeExistsError {
			log.Warn(fmt.Sprintf("adding record type %s failed for %s: %s", r.RecordType, file, err.Error()))
			return
		}
	}
	for _, t := range apply.TabVisibilities {
		err = p.AddTab(t.Tab)
		if err != nil && err != profile.TabExistsError {
			log.Warn(fmt.Sprintf("adding tab %s failed for %s: %s", t.Tab, file, err.Error()))
			return
		}
	}
	for _, l := range apply.LayoutAssignments {
		bits := strings.SplitN(l.Layout, "-", 2)
		objectName := bits[0]
		layoutName := bits[1]
		if l.RecordType != nil {
			p.SetObjectLayoutForRecordType(objectName, layoutName, l.RecordType.Text)
		} else {
			p.SetObjectLayout(objectName, layoutName)
		}
	}
	for _, u := range apply.UserPermissions {
		err = p.AddUserPermission(u.Name)
		if err != nil && err != profile.UserPermissionExistsError {
			log.Warn(fmt.Sprintf("adding user permission %s failed for %s: %s", u.Name, file, err.Error()))
			return
		}
	}
	for _, u := range apply.LoginIPRanges {
		err = p.AddLoginIPRange(u.StartAddress, u.EndAddress, u.Description)
		if err != nil && err != profile.DuplicateIPRangeError {
			log.Warn(fmt.Sprintf("adding ip range %s failed for %s -> %s: %s", u.StartAddress, u.EndAddress, file, err.Error()))
			return
		}
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
