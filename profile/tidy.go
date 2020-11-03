package profile

import (
	"sort"
)

func (p *Profile) Tidy() {
	sort.Slice(p.ApplicationVisibilities, func(i, j int) bool {
		return p.ApplicationVisibilities[i].Application.Text < p.ApplicationVisibilities[j].Application.Text
	})
	sort.Slice(p.ClassAccesses, func(i, j int) bool {
		return p.ClassAccesses[i].ApexClass.Text < p.ClassAccesses[j].ApexClass.Text
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
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab.Text < p.TabVisibilities[j].Tab.Text
	})
	sort.Slice(p.UserPermissions, func(i, j int) bool {
		return p.UserPermissions[i].Name.Text < p.UserPermissions[j].Name.Text
	})
}
