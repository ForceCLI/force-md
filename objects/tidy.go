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
		return p.ValidationRules[i].FullName < p.ValidationRules[j].FullName
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

func (values ValueSetOptionList) Tidy() {
	sort.Slice(values, func(i, j int) bool {
		return values[i].FullName < values[j].FullName
	})
}

func (picklists PicklistList) Tidy() {
	sort.Slice(picklists, func(i, j int) bool {
		return picklists[i].Picklist < picklists[j].Picklist
	})
}
