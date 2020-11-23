package permissionset

import (
	"sort"
)

func (p *PermissionSet) AddClass(className string) {
	p.ClassAccesses = append(p.ClassAccesses, ApexClass{
		ApexClass: className,
		Enabled:   "true",
	})
	sort.Slice(p.ClassAccesses, func(i, j int) bool {
		return p.ClassAccesses[i].ApexClass < p.ClassAccesses[j].ApexClass
	})
}
