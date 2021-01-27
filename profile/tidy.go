package profile

import (
	"sort"
)

func (p *Profile) Tidy() {
	p.ApplicationVisibilities.Tidy()
	sort.Slice(p.ClassAccesses, func(i, j int) bool {
		return p.ClassAccesses[i].ApexClass.Text < p.ClassAccesses[j].ApexClass.Text
	})
	sort.Slice(p.FieldPermissions, func(i, j int) bool {
		return p.FieldPermissions[i].Field.Text < p.FieldPermissions[j].Field.Text
	})
	p.ObjectPermissions.Tidy()
	sort.Slice(p.PageAccesses, func(i, j int) bool {
		return p.PageAccesses[i].ApexPage.Text < p.PageAccesses[j].ApexPage.Text
	})
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab < p.TabVisibilities[j].Tab
	})
	sort.Slice(p.RecordTypeVisibilities, func(i, j int) bool {
		return p.RecordTypeVisibilities[i].RecordType.Text < p.RecordTypeVisibilities[j].RecordType.Text
	})
	sort.Slice(p.FlowAccesses, func(i, j int) bool {
		return p.FlowAccesses[i].Flow.Text < p.FlowAccesses[j].Flow.Text
	})
	p.UserPermissions.Tidy()
}

func (op ObjectPermissionsList) Tidy() {
	sort.Slice(op, func(i, j int) bool {
		return op[i].Object.Text < op[j].Object.Text
	})
}

func (av ApplicationVisibilityList) Tidy() {
	sort.Slice(av, func(i, j int) bool {
		return av[i].Application < av[j].Application
	})
}

func (up UserPermissionList) Tidy() {
	sort.Slice(up, func(i, j int) bool {
		return up[i].Name < up[j].Name
	})
}
