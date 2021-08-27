package profile

import (
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var RecordTypeExistsError = errors.New("record type already exists")

func (p *Profile) AddRecordType(recordType string) error {
	for _, r := range p.RecordTypeVisibilities {
		if r.RecordType == recordType {
			return RecordTypeExistsError
		}
	}
	p.RecordTypeVisibilities = append(p.RecordTypeVisibilities, RecordTypeVisibility{
		RecordType: recordType,
		Visible:    TrueText,
	})
	p.RecordTypeVisibilities.Tidy()
	return nil
}
