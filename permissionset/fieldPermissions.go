package permissionset

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var FieldExistsError = errors.New("field already exists")

type FieldFilter func(FieldPermissions) bool

func (p *PermissionSet) SetFieldPermissions(fieldName string, updates FieldPermissions) error {
	found := false
	for i, f := range p.FieldPermissions {
		if f.Field.Text == fieldName {
			found = true
			if err := mergo.Merge(&updates, f); err != nil {
				return errors.Wrap(err, "merging permissions")
			}
			p.FieldPermissions[i] = updates
		}
	}
	if !found {
		return fmt.Errorf("field not found: %s", fieldName)
	}
	return nil
}

func (p *PermissionSet) DeleteFieldPermissions(fieldName string) error {
	found := false
	newPerms := p.FieldPermissions[:0]
	for _, f := range p.FieldPermissions {
		if f.Field.Text == fieldName {
			found = true
		} else {
			newPerms = append(newPerms, f)
		}
	}
	if !found {
		return errors.New("field not found")
	}
	p.FieldPermissions = newPerms
	return nil
}

func (p *PermissionSet) AddFieldPermissions(fieldName string) error {
	for _, f := range p.FieldPermissions {
		if f.Field.Text == fieldName {
			return FieldExistsError
		}
	}

	p.FieldPermissions = append(p.FieldPermissions, defaultFieldPermissions(fieldName))
	p.FieldPermissions.Tidy()
	return nil
}

func (p *PermissionSet) GetFieldPermissions(filters ...FieldFilter) FieldPermissionsList {
	var fieldPermissions FieldPermissionsList
FIELDS:
	for _, f := range p.FieldPermissions {
		for _, filter := range filters {
			if !filter(f) {
				continue FIELDS
			}
		}
		fieldPermissions = append(fieldPermissions, f)
	}
	return fieldPermissions
}

func defaultFieldPermissions(fieldName string) FieldPermissions {
	fp := FieldPermissions{
		Field:    FieldName{fieldName},
		Editable: FalseText,
		Readable: FalseText,
	}
	return fp
}

func (p *PermissionSet) GetGrantedFieldPermissions() []FieldPermissions {
	var fieldPermissions FieldPermissionsList
	for _, f := range p.FieldPermissions {
		permissionsGranted := false
		fieldPermsGranted := FieldPermissions{
			Field: f.Field,
		}
		if f.Readable.ToBool() {
			fieldPermsGranted.Readable = TrueText
			permissionsGranted = true
		}
		if f.Editable.ToBool() {
			fieldPermsGranted.Editable = TrueText
			permissionsGranted = true
		}
		if permissionsGranted {
			fieldPermissions = append(fieldPermissions, fieldPermsGranted)
		}
	}
	return fieldPermissions
}
