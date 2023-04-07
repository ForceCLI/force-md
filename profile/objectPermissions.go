package profile

import (
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/permissionset"
)

type ObjectFilter func(permissionset.ObjectPermissions) bool

var ObjectExistsError = errors.New("object already exists")

func (p *Profile) SetObjectPermissions(objectName string, updates permissionset.ObjectPermissions) error {
	found := false
	for i, f := range p.ObjectPermissions {
		if strings.ToLower(f.Object) == strings.ToLower(objectName) {
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

func defaultObjectPermissions(objectName string) permissionset.ObjectPermissions {
	op := permissionset.ObjectPermissions{
		Object:           objectName,
		AllowCreate:      FalseText,
		AllowDelete:      FalseText,
		AllowEdit:        FalseText,
		AllowRead:        FalseText,
		ModifyAllRecords: FalseText,
		ViewAllRecords:   FalseText,
	}
	return op
}

func (p *Profile) AddObjectPermissions(objectName string) error {
	for _, f := range p.ObjectPermissions {
		if strings.ToLower(f.Object) == strings.ToLower(objectName) {
			return ObjectExistsError
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
		if strings.ToLower(f.Object) == strings.ToLower(objectName) {
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
		if !strings.HasPrefix(f.Layout, layoutPrefix) {
			newLayouts = append(newLayouts, f)
		}
	}
	p.LayoutAssignments = newLayouts
}

func (p *Profile) DeleteObjectRecordTypeLayoutAssignments(objectName, recordType string) {
	layoutPrefix := objectName + "-"
	newLayouts := p.LayoutAssignments[:0]
	for _, f := range p.LayoutAssignments {
		if f.RecordType == nil || !strings.HasPrefix(f.Layout, layoutPrefix) || strings.ToLower(recordType) != strings.ToLower(f.RecordType.Text) {
			newLayouts = append(newLayouts, f)
		}
	}
	p.LayoutAssignments = newLayouts
}

func (p *Profile) DeleteObjectTabVisibility(objectName string) {
	tabName := "standard-" + objectName
	if strings.HasSuffix(objectName, "__c") {
		tabName = objectName
	}
	newTabs := p.TabVisibilities[:0]
	for _, f := range p.TabVisibilities {
		if strings.ToLower(f.Tab) != strings.ToLower(tabName) {
			newTabs = append(newTabs, f)
		}
	}
	p.TabVisibilities = newTabs
}

func (p *Profile) GetObjectPermissions(filters ...ObjectFilter) []permissionset.ObjectPermissions {
	var objectPermissions []permissionset.ObjectPermissions
OBJECTS:
	for _, o := range p.ObjectPermissions {
		for _, filter := range filters {
			if !filter(o) {
				continue OBJECTS
			}
		}
		objectPermissions = append(objectPermissions, o)
	}
	return objectPermissions
}

func (p *Profile) GetGrantedObjectPermissions() []permissionset.ObjectPermissions {
	var objectPermissions []permissionset.ObjectPermissions
	for _, o := range p.ObjectPermissions {
		permissionsGranted := false
		objectPermsGranted := permissionset.ObjectPermissions{
			Object: o.Object,
		}
		if o.AllowCreate.ToBool() {
			objectPermsGranted.AllowCreate = TrueText
			permissionsGranted = true
		}
		if o.AllowRead.ToBool() {
			objectPermsGranted.AllowRead = TrueText
			permissionsGranted = true
		}
		if o.AllowEdit.ToBool() {
			objectPermsGranted.AllowEdit = TrueText
			permissionsGranted = true
		}
		if o.AllowDelete.ToBool() {
			objectPermsGranted.AllowDelete = TrueText
			permissionsGranted = true
		}
		if o.ViewAllRecords.ToBool() {
			objectPermsGranted.ViewAllRecords = TrueText
			permissionsGranted = true
		}
		if o.ModifyAllRecords.ToBool() {
			objectPermsGranted.ModifyAllRecords = TrueText
			permissionsGranted = true
		}
		if permissionsGranted {
			objectPermissions = append(objectPermissions, objectPermsGranted)
		}
	}
	return objectPermissions
}
