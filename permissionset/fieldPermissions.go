package permissionset

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

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
			return errors.New("field already exists")
		}
	}

	p.FieldPermissions = append(p.FieldPermissions, defaultFieldPermissions(fieldName))
	p.FieldPermissions.Tidy()
	return nil
}

func (p *PermissionSet) GetFieldPermissions() FieldPermissionsList {
	return p.FieldPermissions
}

func defaultFieldPermissions(fieldName string) FieldPermissions {
	var falseBooleanText = BooleanText{
		Text: "false",
	}

	fp := FieldPermissions{
		Field:    FieldName{fieldName},
		Editable: falseBooleanText,
		Readable: falseBooleanText,
	}
	return fp
}
