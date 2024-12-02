package objects

import (
	recordtype "github.com/ForceCLI/force-md/metadata/objects/recordtype"
	"github.com/pkg/errors"
)

func (o *CustomObject) GetRecordTypes(filters ...recordtype.RecordTypeFilter) []recordtype.RecordType {
	var recordTypes []recordtype.RecordType
RECORDTYPES:
	for _, v := range o.RecordTypes {
		for _, filter := range filters {
			if !filter(v) {
				continue RECORDTYPES
			}
		}
		recordTypes = append(recordTypes, v)
	}
	return recordTypes
}

func (o *CustomObject) DeleteRecordType(recordType string) error {
	found := false
	newRecordTypes := o.RecordTypes[:0]
	for _, f := range o.RecordTypes {
		if f.FullName == recordType {
			found = true
		} else {
			newRecordTypes = append(newRecordTypes, f)
		}
	}
	if !found {
		return errors.New("record type not found")
	}
	o.RecordTypes = newRecordTypes
	return nil
}
