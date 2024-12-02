package profile

import (
	"strings"

	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/metadata/permissionset"
)

var UserPermissionExistsError = errors.New("user permission already exists")

type UserPermissionFilter func(permissionset.UserPermission) bool

func (p *Profile) AddUserPermission(permissionName string) error {
	for _, f := range p.UserPermissions {
		if f.Name == permissionName {
			return UserPermissionExistsError
		}
	}
	p.UserPermissions = append(p.UserPermissions, permissionset.UserPermission{
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

func (p *Profile) GetUserPermissions(filters ...UserPermissionFilter) permissionset.UserPermissionList {
	var permissions []permissionset.UserPermission

PERMS:
	for _, v := range p.UserPermissions {
		for _, filter := range filters {
			if !filter(v) {
				continue PERMS
			}
		}
		permissions = append(permissions, v)
	}
	return permissions
}

func (p *Profile) GetEnabledUserPermissions() []string {
	var permissions []string
	for _, u := range p.UserPermissions {
		if u.Enabled.ToBool() {
			permissions = append(permissions, u.Name)
		}
	}
	return permissions
}
