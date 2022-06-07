package permissionset

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

var (
	sourceFileName string
)

func init() {
	MergeCmd.Flags().StringVarP(&sourceFileName, "source", "s", "", "source permission set")
	MergeCmd.MarkFlagRequired("source")
}

var MergeCmd = &cobra.Command{
	Use:                   "merge -s path/to/Source.permissionset [filename]...",
	Short:                 "Merge permissions",
	Long:                  "Apply permissions granted in source permission set",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		grant, err := grantedPermissions(sourceFileName)
		if err != nil {
			log.Fatal("loading source permissions failed: " + err.Error())
		}
		for _, file := range args {
			mergePermissions(file, grant)
		}
	},
}

func mergePermissions(file string, grant permissionset.PermissionSet) {
	p, err := permissionset.Open(file)
	if err != nil {
		log.Warn("parsing permissionset failed: " + err.Error())
		return
	}
	for _, a := range grant.ApplicationVisibilities {
		err = p.AddApplicationVisibility(a.Application)
		if err != nil && err != permissionset.ApplicationExistsError {
			log.Warn(fmt.Sprintf("adding application %s permissions failed for %s: %s", a.Application, file, err.Error()))
			return
		}
	}
	for _, c := range grant.ClassAccesses {
		err = p.AddClass(c.ApexClass)
		if err != nil && err != permissionset.ClassExistsError {
			log.Warn(fmt.Sprintf("adding apex class %s permissions failed for %s: %s", c.ApexClass, file, err.Error()))
			return
		}
	}
	for _, c := range grant.CustomPermissions {
		err = p.AddCustomPermission(c.Name)
		if err != nil && err != permissionset.CustomPermissionExistsError {
			log.Warn(fmt.Sprintf("adding custom permission %s failed for %s: %s", c.Name, file, err.Error()))
			return
		}
	}
	for _, o := range grant.ObjectPermissions {
		objectName := o.Object.Text
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
	for _, f := range grant.FieldPermissions {
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
	for _, v := range grant.PageAccesses {
		err = p.AddVisualforcePageAccess(v.ApexPage)
		if err != nil && err != permissionset.VisualforcePageExistsError {
			log.Warn(fmt.Sprintf("adding visualforce page %s failed for %s: %s", v.ApexPage, file, err.Error()))
			return
		}
	}
	for _, r := range grant.RecordTypeVisibilities {
		err = p.AddRecordType(r.RecordType)
		if err != nil && err != permissionset.RecordTypeExistsError {
			log.Warn(fmt.Sprintf("adding record type %s failed for %s: %s", r.RecordType, file, err.Error()))
			return
		}
	}
	for _, t := range grant.TabSettings {
		err = p.AddTab(t.Tab)
		if err != nil && err != permissionset.TabExistsError {
			log.Warn(fmt.Sprintf("adding tab %s failed for %s: %s", t.Tab, file, err.Error()))
			return
		}
	}
	for _, u := range grant.UserPermissions {
		err = p.AddUserPermission(u.Name)
		if err != nil && err != permissionset.UserPermissionExistsError {
			log.Warn(fmt.Sprintf("adding user permission %s failed for %s: %s", u.Name, file, err.Error()))
			return
		}
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func grantedPermissions(file string) (permissionset.PermissionSet, error) {
	granted := permissionset.PermissionSet{
		ApplicationVisibilities: make(permissionset.ApplicationVisibilityList, 0),
		ClassAccesses:           make(permissionset.ApexClassList, 0),
		CustomPermissions:       make(permissionset.CustomPermissionList, 0),
		FieldPermissions:        make(permissionset.FieldPermissionsList, 0),
		ObjectPermissions:       make(permissionset.ObjectPermissionsList, 0),
		PageAccesses:            make(permissionset.PageAccessList, 0),
		RecordTypeVisibilities:  make(permissionset.RecordTypeList, 0),
		TabSettings:             make(permissionset.TabSettingsList, 0),
		UserPermissions:         make(permissionset.UserPermissionList, 0),
	}
	p, err := permissionset.Open(sourceFileName)
	if err != nil {
		log.Fatal("parsing source permissionset failed: " + err.Error())
		return *p, err
	}
	for _, a := range p.ApplicationVisibilities {
		if a.Visible.ToBool() {
			granted.ApplicationVisibilities = append(granted.ApplicationVisibilities, permissionset.ApplicationVisibility{
				Application: a.Application,
				Visible:     TrueText,
			})
		}
	}
	for _, c := range p.ClassAccesses {
		if c.Enabled.ToBool() {
			granted.ClassAccesses = append(granted.ClassAccesses, permissionset.ApexClass{
				ApexClass: c.ApexClass,
				Enabled:   TrueText,
			})
		}
	}
	for _, c := range p.CustomPermissions {
		if c.Enabled.ToBool() {
			granted.CustomPermissions = append(granted.CustomPermissions, permissionset.CustomPermission{
				Name:    c.Name,
				Enabled: TrueText,
			})
		}
	}
	for _, f := range p.FieldPermissions {
		permissionsGranted := false
		fieldPermsGranted := permissionset.FieldPermissions{
			Field: f.Field,
		}
		if f.Readable.ToBool() {
			fieldPermsGranted.Readable = TrueText
			permissionsGranted = true
		}
		if f.Editable.ToBool() {
			fieldPermsGranted.Editable = TrueText
			permissionsGranted = true
		}
		if permissionsGranted {
			granted.FieldPermissions = append(granted.FieldPermissions, fieldPermsGranted)
		}
	}
	for _, o := range p.ObjectPermissions {
		permissionsGranted := false
		objectPermsGranted := permissionset.ObjectPermissions{
			Object: o.Object,
		}
		if o.AllowCreate.ToBool() {
			objectPermsGranted.AllowCreate = TrueText
			permissionsGranted = true
		}
		if o.AllowRead.ToBool() {
			objectPermsGranted.AllowRead = TrueText
			permissionsGranted = true
		}
		if o.AllowEdit.ToBool() {
			objectPermsGranted.AllowEdit = TrueText
			permissionsGranted = true
		}
		if o.AllowDelete.ToBool() {
			objectPermsGranted.AllowDelete = TrueText
			permissionsGranted = true
		}
		if o.ViewAllRecords.ToBool() {
			objectPermsGranted.ViewAllRecords = TrueText
			permissionsGranted = true
		}
		if o.ModifyAllRecords.ToBool() {
			objectPermsGranted.ModifyAllRecords = TrueText
			permissionsGranted = true
		}
		if permissionsGranted {
			granted.ObjectPermissions = append(granted.ObjectPermissions, objectPermsGranted)
		}
	}
	for _, p := range p.PageAccesses {
		if p.Enabled.ToBool() {
			granted.PageAccesses = append(granted.PageAccesses, permissionset.PageAccess{
				ApexPage: p.ApexPage,
				Enabled:  TrueText,
			})
		}
	}
	for _, r := range p.RecordTypeVisibilities {
		if r.Visible.ToBool() {
			granted.RecordTypeVisibilities = append(granted.RecordTypeVisibilities, permissionset.RecordType{
				RecordType: r.RecordType,
				Visible:    TrueText,
			})
		}
	}
	for _, t := range p.TabSettings {
		if t.IsVisible() {
			granted.TabSettings = append(granted.TabSettings, permissionset.TabSettings{
				Tab:        t.Tab,
				Visibility: t.Visibility,
			})
		}
	}
	for _, u := range p.UserPermissions {
		if u.Enabled.ToBool() {
			granted.UserPermissions = append(granted.UserPermissions, permissionset.UserPermission{
				Name:    u.Name,
				Enabled: TrueText,
			})
		}
	}
	return granted, nil
}
