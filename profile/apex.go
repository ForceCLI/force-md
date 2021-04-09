package profile

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var ClassExistsError = errors.New("apex class already exists")

func (p *Profile) AddClass(className string) error {
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

func (p *Profile) DeleteApexClassAccess(apexClassName string) error {
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

func (p *Profile) GetApexClasses() ApexClassList {
	return p.ClassAccesses
}
