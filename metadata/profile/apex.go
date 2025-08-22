package profile

import (
	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/metadata/permissionset"
)

var ClassExistsError = errors.New("apex class already exists")

func (p *Profile) AddClass(className string) error {
	for _, c := range p.ClassAccesses {
		if c.ApexClass == className {
			return ClassExistsError
		}
	}
	p.ClassAccesses = append(p.ClassAccesses, permissionset.ApexClass{
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

func (p *Profile) EnableApexClassAccess(apexClassName string) error {
	found := false
	for i, f := range p.ClassAccesses {
		if f.ApexClass == apexClassName {
			found = true
			p.ClassAccesses[i].Enabled = TrueText
		}
	}
	if !found {
		return errors.New("class not found")
	}
	return nil
}

func (p *Profile) DisableApexClassAccess(apexClassName string) error {
	found := false
	for i, f := range p.ClassAccesses {
		if f.ApexClass == apexClassName {
			found = true
			p.ClassAccesses[i].Enabled = FalseText
		}
	}
	if !found {
		return errors.New("class not found")
	}
	return nil
}

func (p *Profile) GetApexClasses() permissionset.ApexClassList {
	return p.ClassAccesses
}

func (p *Profile) GetEnabledClasses() []string {
	var classes []string
	for _, v := range p.ClassAccesses {
		if v.Enabled.ToBool() {
			classes = append(classes, v.ApexClass)
		}
	}
	return classes
}

func (p *Profile) CloneApexClassAccess(src, dest string) error {
	for _, c := range p.ClassAccesses {
		if c.ApexClass == dest {
			return errors.New("apex class already exists")
		}
	}
	found := false
	for _, c := range p.ClassAccesses {
		if c.ApexClass == src {
			found = true
			clone := permissionset.ApexClass{}
			clone.Enabled.Text = c.Enabled.Text
			clone.ApexClass = dest
			p.ClassAccesses = append(p.ClassAccesses, clone)
			p.ClassAccesses.Tidy()
			break
		}
	}
	if !found {
		return errors.New("source apex class not found")
	}
	return nil
}
