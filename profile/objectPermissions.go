package profile

import (
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

func (p *Profile) DeleteObjectPermissions(objectName string) {
	found := false
	newObjectPerms := p.ObjectPermissions[:0]
	for _, f := range p.ObjectPermissions {
		if f.Object.Text == objectName {
			found = true
		} else {
			newObjectPerms = append(newObjectPerms, f)
		}
	}
	p.ObjectPermissions = newObjectPerms
	if !found {
		log.Warn(errors.New("object not found"))
	}

	p.DeleteObjectFieldPermissions(objectName)
	p.DeleteObjectTabVisibility(objectName)
	p.DeleteObjectLayoutAssignments(objectName)
}

func (p *Profile) DeleteObjectFieldPermissions(objectName string) {
	newFieldPerms := p.FieldPermissions[:0]
	fieldPrefix := objectName + "."
	for _, f := range p.FieldPermissions {
		if !strings.HasPrefix(f.Field.Text, fieldPrefix) {
			newFieldPerms = append(newFieldPerms, f)
		}
	}
	p.FieldPermissions = newFieldPerms
}

func (p *Profile) DeleteObjectLayoutAssignments(objectName string) {
	layoutPrefix := objectName + "-"
	newLayouts := p.LayoutAssignments[:0]
	for _, f := range p.LayoutAssignments {
		if !strings.HasPrefix(f.Layout.Text, layoutPrefix) {
			newLayouts = append(newLayouts, f)
		}
	}
	p.LayoutAssignments = newLayouts
}

func (p *Profile) DeleteObjectTabVisibility(objectName string) {
	tabPrefix := objectName + "-"
	newTabs := p.TabVisibilities[:0]
	for _, f := range p.TabVisibilities {
		if !strings.HasPrefix(f.Tab, tabPrefix) {
			newTabs = append(newTabs, f)
		}
	}
	p.TabVisibilities = newTabs
}
