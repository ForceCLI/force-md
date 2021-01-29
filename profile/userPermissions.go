package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) AddUserPermission(permissionName string) error {
	for _, f := range p.UserPermissions {
		if f.Name == permissionName {
			return errors.New("permission already exists")
		}
	}
	p.UserPermissions = append(p.UserPermissions, UserPermission{
		Name:    permissionName,
		Enabled: BooleanText{"true"},
	})
	p.UserPermissions.Tidy()
	return nil
}

func (p *Profile) DeleteUserPermission(permissionName string) error {
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

func (p *Profile) GetUserPermissions() UserPermissionList {
	return p.UserPermissions
}
