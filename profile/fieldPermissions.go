package profile

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/permissionset"
)

var FieldExistsError = errors.New("field already exists")

type FieldFilter func(permissionset.FieldPermissions) bool

func (p *Profile) SetFieldPermissions(fieldName string, updates permissionset.FieldPermissions) error {
	found := false
	for i, f := range p.FieldPermissions {
		if strings.ToLower(f.Field) == strings.ToLower(fieldName) {
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

func (p *Profile) DeleteFieldPermissions(fieldName string) error {
	found := false
	newPerms := p.FieldPermissions[:0]
	for _, f := range p.FieldPermissions {
		if strings.ToLower(f.Field) == strings.ToLower(fieldName) {
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

func (p *Profile) AddFieldPermissions(fieldName string) error {
	for _, f := range p.FieldPermissions {
		if strings.ToLower(f.Field) == strings.ToLower(fieldName) {
			return FieldExistsError
		}
	}

	p.FieldPermissions = append(p.FieldPermissions, defaultFieldPermissions(fieldName))
	p.FieldPermissions.Tidy()
	return nil
}

func (p *Profile) GetFieldPermissions(filters ...FieldFilter) permissionset.FieldPermissionsList {
	var fieldPermissions permissionset.FieldPermissionsList
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

func defaultFieldPermissions(fieldName string) permissionset.FieldPermissions {
	fp := permissionset.FieldPermissions{
		Field:    fieldName,
		Editable: FalseText,
		Readable: FalseText,
	}
	return fp
}

func (p *Profile) CloneFieldPermissions(src, dest string) error {
	for _, f := range p.FieldPermissions {
		if strings.ToLower(f.Field) == strings.ToLower(dest) {
			return fmt.Errorf("%s field already exists", dest)
		}
	}
	found := false
	for _, f := range p.FieldPermissions {
		if strings.ToLower(f.Field) == strings.ToLower(src) {
			found = true
			clone := permissionset.FieldPermissions{}
			clone.Editable.Text = f.Editable.Text
			clone.Readable.Text = f.Readable.Text
			clone.Field = dest
			p.FieldPermissions = append(p.FieldPermissions, clone)
		}
	}
	if !found {
		return fmt.Errorf("source field %s not found", src)
	}
	p.FieldPermissions.Tidy()
	return nil
}

func (p *Profile) GetGrantedFieldPermissions() []permissionset.FieldPermissions {
	var fieldPermissions permissionset.FieldPermissionsList
	for _, f := range p.FieldPermissions {
		permissionsGranted := false
		fieldPermsGranted := permissionset.FieldPermissions{
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
