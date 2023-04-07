package permissionset

import (
	"strings"

	"github.com/pkg/errors"

	. "github.com/ForceCLI/force-md/general"
)

var RecordTypeExistsError = errors.New("record type already exists")

func (p *PermissionSet) AddRecordType(recordType string) error {
	for _, r := range p.RecordTypeVisibilities {
		if r.RecordType == recordType {
			return RecordTypeExistsError
		}
	}
	p.RecordTypeVisibilities = append(p.RecordTypeVisibilities, RecordType{
		RecordType: recordType,
		Visible:    TrueText,
	})
	p.RecordTypeVisibilities.Tidy()
	return nil
}

func (p *PermissionSet) GetRecordTypeVisibility() RecordTypeList {
	return p.RecordTypeVisibilities
}

func (p *PermissionSet) DeleteRecordType(recordtype string) error {
	bits := strings.SplitN(recordtype, ".", 2)
	if len(bits) != 2 {
		return errors.New("record type should be ObjectName.RecordTypeName")
	}
	found := false
	newPerms := p.RecordTypeVisibilities[:0]
	for _, f := range p.RecordTypeVisibilities {
		if f.RecordType == recordtype {
			found = true
		} else {
			newPerms = append(newPerms, f)
		}
	}
	if !found {
		return errors.New("record type not found")
	}
	p.RecordTypeVisibilities = newPerms
	return nil
}

func (p *PermissionSet) GetVisibleRecordTypes() []string {
	var recordTypes []string
	for _, r := range p.RecordTypeVisibilities {
		if r.Visible.ToBool() {
			recordTypes = append(recordTypes, r.RecordType)
		}
	}
	return recordTypes
}
