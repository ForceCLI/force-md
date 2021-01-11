package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) DeleteVisualforcePageAccess(pageName string) error {
	found := false
	newPages := p.PageAccesses[:0]
	for _, f := range p.PageAccesses {
		if f.ApexPage.Text == pageName {
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
