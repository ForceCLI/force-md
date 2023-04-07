package permissionset

import (
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	. "github.com/ForceCLI/force-md/general"
)

type ObjectFilter func(ObjectPermissions) bool

var ObjectExistsError = errors.New("object already exists")

func (p *PermissionSet) SetObjectPermissions(objectName string, updates ObjectPermissions) error {
	found := false
	for i, f := range p.ObjectPermissions {
		if f.Object == objectName {
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
	op := ObjectPermissions{
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

func (p *PermissionSet) AddObjectPermissions(objectName string) error {
	for _, f := range p.ObjectPermissions {
		if strings.ToLower(f.Object) == strings.ToLower(objectName) {
			return ObjectExistsError
		}
	}

	p.ObjectPermissions = append(p.ObjectPermissions, defaultObjectPermissions(objectName))
	p.ObjectPermissions.Tidy()
	return nil
}

func (p *PermissionSet) DeleteObjectPermissions(objectName string) {
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
}

func (p *PermissionSet) DeleteObjectFieldPermissions(objectName string) {
	newFieldPerms := p.FieldPermissions[:0]
	fieldPrefix := strings.ToLower(objectName + ".")
	for _, f := range p.FieldPermissions {
		if !strings.HasPrefix(strings.ToLower(f.Field), fieldPrefix) {
			newFieldPerms = append(newFieldPerms, f)
		}
	}
	p.FieldPermissions = newFieldPerms
}

func (p *PermissionSet) DeleteObjectTabVisibility(objectName string) {
	newTabs := p.TabSettings[:0]
	for _, f := range p.TabSettings {
		if f.Tab != objectName {
			newTabs = append(newTabs, f)
		}
	}
	p.TabSettings = newTabs
}

func (p *PermissionSet) GetObjectPermissions(filters ...ObjectFilter) []ObjectPermissions {
	var objectPermissions []ObjectPermissions
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

func (p *PermissionSet) GetGrantedObjectPermissions() []ObjectPermissions {
	var objectPermissions []ObjectPermissions
	for _, o := range p.ObjectPermissions {
		permissionsGranted := false
		objectPermsGranted := ObjectPermissions{
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
