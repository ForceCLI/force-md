package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
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

func (p *PermissionSet) DeleteCustomPermission(permissionName string) error {
	found := false
	newPerms := p.CustomPermissions[:0]
	for _, f := range p.CustomPermissions {
		if f.Name == permissionName {
			found = true
		} else {
			newPerms = append(newPerms, f)
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	p.CustomPermissions = newPerms
	return nil
}

func (p *PermissionSet) GetCustomPermissions() CustomPermissionList {
	return p.CustomPermissions
}

func (p *PermissionSet) GetEnabledCustomPermissions() []string {
	var permissions []string
	for _, v := range p.CustomPermissions {
		if v.Enabled.ToBool() {
			permissions = append(permissions, v.Name)
		}
	}
	return permissions
}
