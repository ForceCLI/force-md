package profile

import (
	"strings"

	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/permissionset"
)

var CustomPermissionExistsError = errors.New("custom permission already exists")

type CustomPermissionFilter func(permissionset.CustomPermission) bool

func (p *Profile) GetEnabledCustomPermissions() []string {
	var permissions []string
	for _, v := range p.CustomPermissions {
		if v.Enabled.ToBool() {
			permissions = append(permissions, v.Name)
		}
	}
	return permissions
}

func (p *Profile) AddCustomPermission(permissionName string) error {
	for _, f := range p.CustomPermissions {
		if f.Name == permissionName {
			return CustomPermissionExistsError
		}
	}
	p.CustomPermissions = append(p.CustomPermissions, permissionset.CustomPermission{
		Name:    permissionName,
		Enabled: BooleanText{"true"},
	})
	p.CustomPermissions.Tidy()
	return nil
}

func (p *Profile) DeleteCustomPermission(permissionName string) error {
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

func (p *Profile) EnableCustomPermission(permissionName string) error {
	found := false
	for i, u := range p.CustomPermissions {
		if strings.ToLower(u.Name) == strings.ToLower(permissionName) {
			found = true
			p.CustomPermissions[i].Enabled.Text = "true"
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	return nil
}

func (p *Profile) DisableCustomPermission(permissionName string) error {
	found := false
	for i, u := range p.CustomPermissions {
		if strings.ToLower(u.Name) == strings.ToLower(permissionName) {
			found = true
			p.CustomPermissions[i].Enabled.Text = "false"
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	return nil
}

func (p *Profile) GetCustomPermissions(filters ...CustomPermissionFilter) permissionset.CustomPermissionList {
	var permissions []permissionset.CustomPermission

PERMS:
	for _, v := range p.CustomPermissions {
		for _, filter := range filters {
			if !filter(v) {
				continue PERMS
			}
		}
		permissions = append(permissions, v)
	}
	return permissions
}
