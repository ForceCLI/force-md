package objects

import (
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

type FieldFilter func(Field) bool

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

func (o *CustomObject) UpdateField(fieldName string, updates Field) error {
	found := false
	for i, f := range o.Fields {
		if f.FullName.Text == fieldName {
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
