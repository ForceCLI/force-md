package objects

import (
	"sort"
)

func (p *CustomObject) Tidy() {
	sort.Slice(p.FieldSets, func(i, j int) bool {
		return p.FieldSets[i].FullName < p.FieldSets[j].FullName
	})
	p.Fields.Tidy()
	sort.Slice(p.ValidationRules, func(i, j int) bool {
		return p.ValidationRules[i].FullName.Text < p.ValidationRules[j].FullName.Text
	})
	sort.Slice(p.ListViews, func(i, j int) bool {
		return p.ListViews[i].FullName.Text < p.ListViews[j].FullName.Text
	})
}

func (fields FieldList) Tidy() {
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].FullName < fields[j].FullName
	})
}
