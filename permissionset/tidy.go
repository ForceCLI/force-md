package permissionset

import (
	"sort"
)

func (p *PermissionSet) Tidy() {
	sort.Slice(p.ApplicationVisibilities, func(i, j int) bool {
		return p.ApplicationVisibilities[i].Application.Text < p.ApplicationVisibilities[j].Application.Text
	})
	sort.Slice(p.ClassAccesses, func(i, j int) bool {
		return p.ClassAccesses[i].ApexClass < p.ClassAccesses[j].ApexClass
	})
	sort.Slice(p.FieldPermissions, func(i, j int) bool {
		return p.FieldPermissions[i].Field.Text < p.FieldPermissions[j].Field.Text
	})
	sort.Slice(p.ObjectPermissions, func(i, j int) bool {
		return p.ObjectPermissions[i].Object.Text < p.ObjectPermissions[j].Object.Text
	})
	sort.Slice(p.PageAccesses, func(i, j int) bool {
		return p.PageAccesses[i].ApexPage.Text < p.PageAccesses[j].ApexPage.Text
	})
	sort.Slice(p.TabSettings, func(i, j int) bool {
		return p.TabSettings[i].Tab < p.TabSettings[j].Tab
	})
	sort.Slice(p.UserPermissions, func(i, j int) bool {
		return p.UserPermissions[i].Name.Text < p.UserPermissions[j].Name.Text
	})
	sort.Slice(p.RecordTypeVisibilities, func(i, j int) bool {
		return p.RecordTypeVisibilities[i].RecordType.Text < p.RecordTypeVisibilities[j].RecordType.Text
	})
	sort.Slice(p.CustomPermissions, func(i, j int) bool {
		return p.CustomPermissions[i].Name.Text < p.CustomPermissions[j].Name.Text
	})
}
