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
	p.FieldPermissions.Tidy()
	p.ObjectPermissions.Tidy()
	sort.Slice(p.PageAccesses, func(i, j int) bool {
		return p.PageAccesses[i].ApexPage.Text < p.PageAccesses[j].ApexPage.Text
	})
	sort.Slice(p.TabSettings, func(i, j int) bool {
		return p.TabSettings[i].Tab < p.TabSettings[j].Tab
	})
	p.UserPermissions.Tidy()
	sort.Slice(p.RecordTypeVisibilities, func(i, j int) bool {
		return p.RecordTypeVisibilities[i].RecordType.Text < p.RecordTypeVisibilities[j].RecordType.Text
	})
	sort.Slice(p.CustomPermissions, func(i, j int) bool {
		return p.CustomPermissions[i].Name.Text < p.CustomPermissions[j].Name.Text
	})
}

func (op ObjectPermissionsList) Tidy() {
	sort.Slice(op, func(i, j int) bool {
		return op[i].Object.Text < op[j].Object.Text
	})
}

func (fp FieldPermissionsList) Tidy() {
	sort.Slice(fp, func(i, j int) bool {
		return fp[i].Field.Text < fp[j].Field.Text
	})
}

func (up UserPermissionList) Tidy() {
	sort.Slice(up, func(i, j int) bool {
		return up[i].Name < up[j].Name
	})
}
