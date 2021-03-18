package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var CustomPermissionExistsError = errors.New("custom permission already exists")

func (p *PermissionSet) AddCustomPermission(permission string) error {
	for _, c := range p.CustomPermissions {
		if c.Name == permission {
			return CustomPermissionExistsError
		}
	}
	p.CustomPermissions = append(p.CustomPermissions, CustomPermission{
		Name:    permission,
		Enabled: TrueText,
	})
	p.CustomPermissions.Tidy()
	return nil
}
