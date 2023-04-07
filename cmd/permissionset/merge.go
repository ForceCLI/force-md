package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/permissionGranter"
	"github.com/ForceCLI/force-md/permissionset"
)

var (
	sourceFileName string
)

func init() {
	MergeCmd.Flags().StringVarP(&sourceFileName, "source", "s", "", "source permission set or profile")
	MergeCmd.MarkFlagRequired("source")
}

var MergeCmd = &cobra.Command{
	Use:                   "merge -s path/to/Source.permissionset [filename]...",
	Short:                 "Merge permissions",
	Long:                  "Apply permissions granted in source permission set or profile",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		grant, err := permissionGranter.Open(sourceFileName)
		if err != nil {
			log.Fatal("loading source permissions failed: " + err.Error())
		}
		for _, file := range args {
			mergePermissions(file, grant)
		}
	},
}

func mergePermissions(file string, granter permissionGranter.PermissionGranter) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	for _, a := range granter.GetVisibleApplications() {
		err = p.AddApplicationVisibility(a)
		if err != nil && err != permissionset.ApplicationExistsError {
			log.Warn(fmt.Sprintf("adding application %s permissions failed for %s: %s", a, file, err.Error()))
			return
		}
	}
	for _, c := range granter.GetEnabledClasses() {
		err = p.AddClass(c)
		if err != nil && err != permissionset.ClassExistsError {
			log.Warn(fmt.Sprintf("adding apex class %s permissions failed for %s: %s", c, file, err.Error()))
			return
		}
	}
	for _, c := range granter.GetEnabledCustomPermissions() {
		err = p.AddCustomPermission(c)
		if err != nil && err != permissionset.CustomPermissionExistsError {
			log.Warn(fmt.Sprintf("adding custom permission %s failed for %s: %s", c, file, err.Error()))
			return
		}
	}
	for _, o := range granter.GetGrantedObjectPermissions() {
		objectName := o.Object
		err = p.AddObjectPermissions(objectName)
		if err != nil && err != permissionset.ObjectExistsError {
			log.Warn(fmt.Sprintf("adding object %s permissions failed for %s: %s", objectName, file, err.Error()))
			return
		}
		err = p.SetObjectPermissions(objectName, o)
		if err != nil {
			log.Warn(fmt.Sprintf("updating object %s permissions failed for %s: %s", objectName, file, err.Error()))
			return
		}
	}
	for _, f := range granter.GetGrantedFieldPermissions() {
		fieldName := f.Field.Text
		err = p.AddFieldPermissions(fieldName)
		if err != nil && err != permissionset.FieldExistsError {
			log.Warn(fmt.Sprintf("adding field %s permissions failed for %s: %s", fieldName, file, err.Error()))
			return
		}
		err = p.SetFieldPermissions(fieldName, f)
		if err != nil {
			log.Warn(fmt.Sprintf("updating field %s permissions failed for %s: %s", fieldName, file, err.Error()))
			return
		}
	}
	for _, v := range granter.GetEnabledPageAccesses() {
		err = p.AddVisualforcePageAccess(v)
		if err != nil && err != permissionset.VisualforcePageExistsError {
			log.Warn(fmt.Sprintf("adding visualforce page %s failed for %s: %s", v, file, err.Error()))
			return
		}
	}
	for _, r := range granter.GetVisibleRecordTypes() {
		err = p.AddRecordType(r)
		if err != nil && err != permissionset.RecordTypeExistsError {
			log.Warn(fmt.Sprintf("adding record type %s failed for %s: %s", r, file, err.Error()))
			return
		}
	}
	for _, u := range granter.GetEnabledUserPermissions() {
		err = p.AddUserPermission(u)
		if err != nil && err != permissionset.UserPermissionExistsError {
			log.Warn(fmt.Sprintf("adding user permission %s failed for %s: %s", u, file, err.Error()))
			return
		}
	}
	p.Tidy()
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
