package permissionset

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
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
	p.CustomPermissions.Tidy()
	return nil
}
