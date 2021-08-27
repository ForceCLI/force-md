package profile

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var FieldExistsError = errors.New("field already exists")

func (p *Profile) SetFieldPermissions(fieldName string, updates FieldPermissions) error {
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

func (p *Profile) DeleteFieldPermissions(fieldName string) error {
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

func (p *Profile) AddFieldPermissions(fieldName string) error {
	for _, f := range p.FieldPermissions {
		if f.Field.Text == fieldName {
			return FieldExistsError
		}
	}

	p.FieldPermissions = append(p.FieldPermissions, defaultFieldPermissions(fieldName))
	p.FieldPermissions.Tidy()
	return nil
}

func (p *Profile) GetFieldPermissions() FieldPermissionsList {
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

func (p *Profile) CloneFieldPermissions(src, dest string) error {
	for _, f := range p.FieldPermissions {
		if f.Field.Text == dest {
			return fmt.Errorf("%s field already exists", dest)
		}
	}
	found := false
	for _, f := range p.FieldPermissions {
		if f.Field.Text == src {
			found = true
			clone := FieldPermissions{}
			clone.Editable.Text = f.Editable.Text
			clone.Readable.Text = f.Readable.Text
			clone.Field.Text = dest
			p.FieldPermissions = append(p.FieldPermissions, clone)
		}
	}
	if !found {
		return fmt.Errorf("source field %s not found", src)
	}
	p.FieldPermissions.Tidy()
	return nil
}
