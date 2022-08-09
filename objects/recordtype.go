package objects

import "github.com/pkg/errors"

func (o *CustomObject) GetRecordTypes() []RecordType {
	return o.RecordTypes
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
