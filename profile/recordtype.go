package profile

import (
	"fmt"

	"github.com/imdario/mergo"
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
		Default:    FalseText,
		Visible:    TrueText,
	})
	p.RecordTypeVisibilities.Tidy()
	return nil
}

func (p *Profile) CloneRecordType(src, dest string) error {
	for _, f := range p.RecordTypeVisibilities {
		if f.RecordType == dest {
			return fmt.Errorf("%s record type already exists", dest)
		}
	}
	found := false
	for _, f := range p.RecordTypeVisibilities {
		if f.RecordType == src {
			found = true
			clone := RecordTypeVisibility{}
			clone.Visible.Text = f.Visible.Text
			clone.Default.Text = f.Default.Text
			clone.RecordType = dest
			p.RecordTypeVisibilities = append(p.RecordTypeVisibilities, clone)
		}
	}
	if !found {
		return fmt.Errorf("source record type %s not found", src)
	}
	p.RecordTypeVisibilities.Tidy()
	return nil
}

func (p *Profile) DeleteRecordType(recordtype string) error {
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

func (p *Profile) SetRecordTypeVisibility(fieldName string, updates RecordTypeVisibility) error {
	found := false
	for i, f := range p.RecordTypeVisibilities {
		if f.RecordType == fieldName {
			found = true
			if err := mergo.Merge(&updates, f); err != nil {
				return errors.Wrap(err, "merging permissions")
			}
			p.RecordTypeVisibilities[i] = updates
		}
	}
	if !found {
		return fmt.Errorf("record type not found: %s", fieldName)
	}
	return nil
}

func (p *Profile) GetRecordTypeVisibility() RecordTypeVisibilityList {
	return p.RecordTypeVisibilities
}
