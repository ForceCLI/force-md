package objects

import (
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

type FieldFilter func(Field) bool

func defaultField(name string) Field {
	f := Field{
		FullName: name,
	}
	return f
}

func (o *CustomObject) GetFields(filters ...FieldFilter) []Field {
	var fields []Field
FIELDS:
	for _, f := range o.Fields {
		for _, filter := range filters {
			if !filter(f) {
				continue FIELDS
			}
		}
		fields = append(fields, f)
	}
	return fields
}

func (o *CustomObject) AddField(fieldName string) error {
	for _, f := range o.Fields {
		if f.FullName == fieldName {
			return errors.New("field already exists")
		}
	}
	f := defaultField(fieldName)
	o.Fields = append(o.Fields, f)
	o.Fields.Tidy()
	return nil
}

func (o *CustomObject) UpdateField(fieldName string, updates Field) error {
	found := false
	for i, f := range o.Fields {
		if strings.ToLower(f.FullName) == strings.ToLower(fieldName) {
			found = true
			if err := mergo.Merge(&updates, f, mergo.WithNoOverrideEmptyStructValues); err != nil {
				return errors.Wrap(err, "merging field updates")
			}
			o.Fields[i] = updates
		}
	}
	if !found {
		return errors.New("field not found")
	}
	return nil
}

func (p *CustomObject) DeleteField(fieldName string) error {
	found := false
	newFields := p.Fields[:0]
	for _, f := range p.Fields {
		if f.FullName == fieldName {
			found = true
		} else {
			newFields = append(newFields, f)
		}
	}
	if !found {
		return errors.New("field not found")
	}
	p.Fields = newFields
	p.DeleteFieldPicklistValues(fieldName)
	p.DeleteFieldFromCompactLayouts(fieldName)
	return nil
}

func (p *CustomObject) DeleteFieldPicklistValues(fieldName string) {
	for i, f := range p.RecordTypes {
		newPicklistValues := f.PicklistValues[:0]
		for _, p := range f.PicklistValues {
			if p.Picklist != fieldName {
				newPicklistValues = append(newPicklistValues, p)
			}
		}
		p.RecordTypes[i].PicklistValues = newPicklistValues
	}
}

func (o *CustomObject) AddFieldPicklistValue(fieldName string, recordType string, picklistValue string) error {
	found := false
	for i, f := range o.RecordTypes {
		if strings.ToLower(f.FullName) != strings.ToLower(recordType) {
			continue
		}
		for j, p := range f.PicklistValues {
			if strings.ToLower(p.Picklist) != strings.ToLower(fieldName) {
				continue
			}
			for _, v := range p.Values {
				if strings.ToLower(v.FullName) == strings.ToLower(picklistValue) {
					return errors.New("value already exists")
				}
			}
			found = true
			o.RecordTypes[i].PicklistValues[j].Values = append(o.RecordTypes[i].PicklistValues[j].Values, ValueSetOption{
				FullName: picklistValue,
				Default:  FalseText,
			})
		}
	}
	if !found {
		return errors.New("field not found")
	}
	return nil
}

func (p *CustomObject) DeleteFieldFromCompactLayouts(fieldName string) {
	for i, f := range p.CompactLayouts {
		newFields := f.Fields[:0]
		for _, p := range f.Fields {
			if strings.ToLower(p.Text) != strings.ToLower(fieldName) {
				newFields = append(newFields, p)
			}
		}
		p.CompactLayouts[i].Fields = newFields
	}
}
