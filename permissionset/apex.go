package permissionset

import (
	"sort"

	"github.com/pkg/errors"
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

func (p *PermissionSet) DeleteApexClassAccess(apexClassName string) error {
	found := false
	newClasses := p.ClassAccesses[:0]
	for _, f := range p.ClassAccesses {
		if f.ApexClass == apexClassName {
			found = true
		} else {
			newClasses = append(newClasses, f)
		}
	}
	if !found {
		return errors.New("class not found")
	}
	p.ClassAccesses = newClasses
	return nil
}
