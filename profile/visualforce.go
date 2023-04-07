package profile

import (
	"fmt"

	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/permissionset"
)

var VisualforcePageExistsError = errors.New("visualforce page already exists")

func (p *Profile) AddVisualforcePageAccess(pageName string) error {
	for _, f := range p.PageAccesses {
		if f.ApexPage == pageName {
			return VisualforcePageExistsError
		}
	}

	p.PageAccesses = append(p.PageAccesses, permissionset.PageAccess{ApexPage: pageName, Enabled: TrueText})
	p.PageAccesses.Tidy()
	return nil
}

func (p *Profile) DeleteVisualforcePageAccess(pageName string) error {
	found := false
	newPages := p.PageAccesses[:0]
	for _, f := range p.PageAccesses {
		if f.ApexPage == pageName {
			found = true
		} else {
			newPages = append(newPages, f)
		}
	}
	if !found {
		return errors.New("page not found")
	}
	p.PageAccesses = newPages
	return nil
}

func (p *Profile) GetVisualforcePageVisibility() permissionset.PageAccessList {
	return p.PageAccesses
}

func (p *Profile) GetEnabledPageAccesses() []string {
	var pages []string
	for _, v := range p.PageAccesses {
		if v.Enabled.ToBool() {
			pages = append(pages, v.ApexPage)
		}
	}
	return pages
}

func (p *Profile) CloneVisualforcePageAccess(src, dest string) error {
	for _, f := range p.PageAccesses {
		if f.ApexPage == dest {
			return fmt.Errorf("%s page already exists", dest)
		}
	}
	found := false
	for _, f := range p.PageAccesses {
		if f.ApexPage == src {
			found = true
			clone := permissionset.PageAccess{}
			clone.Enabled.Text = f.Enabled.Text
			clone.ApexPage = dest
			p.PageAccesses = append(p.PageAccesses, clone)
		}
	}
	if !found {
		return fmt.Errorf("source page %s not found", src)
	}
	p.PageAccesses.Tidy()
	return nil
}
