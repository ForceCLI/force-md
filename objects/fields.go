package objects

import ()

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
