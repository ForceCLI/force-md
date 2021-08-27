package profile

import (
	"strings"

	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var UserPermissionExistsError = errors.New("user permissions already exists")

func (p *Profile) AddUserPermission(permissionName string) error {
	for _, f := range p.UserPermissions {
		if f.Name == permissionName {
			return UserPermissionExistsError
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

func (p *Profile) EnableUserPermission(permissionName string) error {
	found := false
	for i, u := range p.UserPermissions {
		if strings.ToLower(u.Name) == strings.ToLower(permissionName) {
			found = true
			p.UserPermissions[i].Enabled.Text = "true"
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	return nil
}

func (p *Profile) DisableUserPermission(permissionName string) error {
	found := false
	for i, u := range p.UserPermissions {
		if strings.ToLower(u.Name) == strings.ToLower(permissionName) {
			found = true
			p.UserPermissions[i].Enabled.Text = "false"
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	return nil
}

func (p *Profile) GetUserPermissions() UserPermissionList {
	return p.UserPermissions
}
