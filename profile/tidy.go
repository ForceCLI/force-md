package profile

import (
	"sort"
)

func (p *Profile) Tidy() {
	p.ApplicationVisibilities.Tidy()
	p.ClassAccesses.Tidy()
	p.FieldPermissions.Tidy()
	p.ObjectPermissions.Tidy()
	p.LayoutAssignments.Tidy()
	p.PageAccesses.Tidy()
	p.RecordTypeVisibilities.Tidy()
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab < p.TabVisibilities[j].Tab
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

func (ca ApexClassList) Tidy() {
	sort.Slice(ca, func(i, j int) bool {
		return ca[i].ApexClass < ca[j].ApexClass
	})
}

func (fp FieldPermissionsList) Tidy() {
	sort.Slice(fp, func(i, j int) bool {
		return fp[i].Field.Text < fp[j].Field.Text
	})
}

func (av ApplicationVisibilityList) Tidy() {
	sort.Slice(av, func(i, j int) bool {
		return av[i].Application < av[j].Application
	})
}

func (la LayoutAssignmentList) Tidy() {
	sort.Slice(la, func(i, j int) bool {
		left := la[i].Layout
		right := la[j].Layout
		if la[i].RecordType != nil {
			left += la[i].RecordType.Text
		}
		if la[j].RecordType != nil {
			right += la[j].RecordType.Text
		}

		return left < right
	})
}

func (up UserPermissionList) Tidy() {
	sort.Slice(up, func(i, j int) bool {
		return up[i].Name < up[j].Name
	})
}

func (pa PageAccessList) Tidy() {
	sort.Slice(pa, func(i, j int) bool {
		return pa[i].ApexPage < pa[j].ApexPage
	})
}

func (rt RecordTypeVisibilityList) Tidy() {
	sort.Slice(rt, func(i, j int) bool {
		return rt[i].RecordType < rt[j].RecordType
	})
}
