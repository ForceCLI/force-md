package objects

import (
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
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
		if f.FullName == fieldName {
			found = true
			if err := mergo.Merge(&updates, f); err != nil {
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
	return nil
}
