package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var UserPermissionExistsError = errors.New("user permissions already exists")

func (p *PermissionSet) AddUserPermission(permissionName string) error {
	for _, f := range p.UserPermissions {
		if f.Name == permissionName {
			return UserPermissionExistsError
		}
	}
	p.UserPermissions = append(p.UserPermissions, UserPermission{
		Name:    permissionName,
		Enabled: TrueText,
	})
	p.UserPermissions.Tidy()
	return nil
}

func (p *PermissionSet) DeleteUserPermission(permissionName string) error {
	found := false
	newPerms := p.UserPermissions[:0]
	for _, f := range p.UserPermissions {
		if f.Name == permissionName {
			found = true
		} else {
			newPerms = append(newPerms, f)
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	p.UserPermissions = newPerms
	return nil
}

func (p *PermissionSet) GetUserPermissions() UserPermissionList {
	return p.UserPermissions
}

func (p *PermissionSet) GetEnabledUserPermissions() []string {
	var permissions []string
	for _, u := range p.UserPermissions {
		if u.Enabled.ToBool() {
			permissions = append(permissions, u.Name)
		}
	}
	return permissions
}
