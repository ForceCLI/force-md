package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
)

var ClassExistsError = errors.New("apex class already exists")

func (p *PermissionSet) AddClass(className string) error {
	for _, c := range p.ClassAccesses {
		if c.ApexClass == className {
			return ClassExistsError
		}
	}
	p.ClassAccesses = append(p.ClassAccesses, ApexClass{
		ApexClass: className,
		Enabled:   TrueText,
	})
	p.ClassAccesses.Tidy()
	return nil
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

func (p *PermissionSet) GetApexClasses() ApexClassList {
	return p.ClassAccesses
}

func (p *PermissionSet) GetEnabledClasses() []string {
	var classes []string
	for _, v := range p.ClassAccesses {
		if v.Enabled.ToBool() {
			classes = append(classes, v.ApexClass)
		}
	}
	return classes
}
