package profile

import (
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

func (p *Profile) SetObjectPermissions(objectName string, updates ObjectPermissions) error {
	found := false
	for i, f := range p.ObjectPermissions {
		if f.Object.Text == objectName {
			found = true
			if err := mergo.Merge(&updates, f); err != nil {
				return errors.Wrap(err, "merging permissions")
			}
			p.ObjectPermissions[i] = updates
		}
	}
	if !found {
		return errors.New("object not found")
	}
	return nil
}

func defaultObjectPermissions(objectName string) ObjectPermissions {
	var falseBooleanText = BooleanText{
		Text: "false",
	}

	op := ObjectPermissions{
		Object:           ObjectName{objectName},
		AllowCreate:      falseBooleanText,
		AllowDelete:      falseBooleanText,
		AllowEdit:        falseBooleanText,
		AllowRead:        falseBooleanText,
		ModifyAllRecords: falseBooleanText,
		ViewAllRecords:   falseBooleanText,
	}
	return op
}

func (p *Profile) AddObjectPermissions(objectName string) error {
	for _, f := range p.ObjectPermissions {
		if f.Object.Text == objectName {
			return errors.New("object already exists")
		}
	}

	p.ObjectPermissions = append(p.ObjectPermissions, defaultObjectPermissions(objectName))
	p.ObjectPermissions.Tidy()
	return nil
}

func (p *Profile) DeleteObjectPermissions(objectName string) error {
	found := false
	for i, f := range p.ObjectPermissions {
		if f.Object.Text == objectName {
			p.ObjectPermissions = append(p.ObjectPermissions[:i], p.ObjectPermissions[i+1:]...)
			found = true
		}
	}
	if !found {
		return errors.New("object not found")
	}
	fieldPrefix := objectName + "."
	for i, f := range p.FieldPermissions {
		if strings.HasPrefix(f.Field.Text, fieldPrefix) {
			p.FieldPermissions = append(p.FieldPermissions[:i], p.FieldPermissions[i+1:]...)
			found = true
		}
	}
	return nil
}
