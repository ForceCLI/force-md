package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var VisualforcePageExistsError = errors.New("visualforce page already exists")

func (p *PermissionSet) AddPageAccess(pageName string) error {
	for _, v := range p.PageAccesses {
		if v.ApexPage == pageName {
			return VisualforcePageExistsError
		}
	}
	p.PageAccesses = append(p.PageAccesses, PageAccess{
		ApexPage: pageName,
		Enabled:  TrueText,
	})
	p.CustomPermissions.Tidy()
	return nil
}
