package profile

import (
	"sort"
)

func (p *Profile) Tidy() {
	p.ApplicationVisibilities.Tidy()
	p.ClassAccesses.Tidy()
	p.CustomPermissions.Tidy()
	p.FieldPermissions.Tidy()
	p.ObjectPermissions.Tidy()
	p.LayoutAssignments.Tidy()
	p.PageAccesses.Tidy()
	p.RecordTypeVisibilities.Tidy()
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab < p.TabVisibilities[j].Tab
	})
	sort.Slice(p.FlowAccesses, func(i, j int) bool {
		return p.FlowAccesses[i].Flow < p.FlowAccesses[j].Flow
	})
	p.UserPermissions.Tidy()
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

func (rt RecordTypeVisibilityList) Tidy() {
	sort.Slice(rt, func(i, j int) bool {
		return rt[i].RecordType < rt[j].RecordType
	})
}
