package objectTranslations

import (
	"sort"
)

// Tidy sorts the fields, record types, and validation rules in the CustomObjectTranslation
func (c *CustomObjectTranslation) Tidy() {
	// Sort fields by name
	sort.Slice(c.Fields, func(i, j int) bool {
		return c.Fields[i].Name.Text < c.Fields[j].Name.Text
	})

	// Sort record types by name
	sort.Slice(c.RecordTypes, func(i, j int) bool {
		return c.RecordTypes[i].Name.Text < c.RecordTypes[j].Name.Text
	})

	// Sort validation rules by name
	sort.Slice(c.ValidationRules, func(i, j int) bool {
		return c.ValidationRules[i].Name.Text < c.ValidationRules[j].Name.Text
	})
}
