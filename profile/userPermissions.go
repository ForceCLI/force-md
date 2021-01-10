package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) DeleteUserPermission(permissionName string) error {
	found := false
	for i, f := range p.UserPermissions {
		if f.Name.Text == permissionName {
			p.UserPermissions = append(p.UserPermissions[:i], p.UserPermissions[i+1:]...)
			found = true
		}
	}
	if !found {
		return errors.New("permission not found")
	}
	return nil
}
