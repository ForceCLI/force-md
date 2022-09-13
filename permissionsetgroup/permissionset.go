package permissionsetgroup

import (
	"strings"

	"github.com/pkg/errors"
)

var PermissionSetExistsError = errors.New("permission set already exists")

func (p *PermissionSetGroup) AddPermissionSet(permissionSetName string) error {
	for _, c := range p.PermissionSets {
		if strings.ToLower(c.Text) == strings.ToLower(permissionSetName) {
			return PermissionSetExistsError
		}
	}
	p.PermissionSets = append(p.PermissionSets, PermissionSet{
		Text: permissionSetName,
	})
	p.PermissionSets.Tidy()
	return nil
}

func (p *PermissionSetGroup) DeletePermissionSet(permissionSetName string) error {
	found := false
	newPermissionSets := p.PermissionSets[:0]
	for _, f := range p.PermissionSets {
		if strings.ToLower(f.Text) == strings.ToLower(permissionSetName) {
			found = true
		} else {
			newPermissionSets = append(newPermissionSets, f)
		}
	}
	if !found {
		return errors.New("permission set not found")
	}
	p.PermissionSets = newPermissionSets
	return nil
}

func (p *PermissionSetGroup) GetPermissionSets() PermissionSetList {
	return p.PermissionSets
}
