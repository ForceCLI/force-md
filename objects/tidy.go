package objects

import (
	"sort"
)

func (p *CustomObject) Tidy() {
	sort.Slice(p.FieldSets, func(i, j int) bool {
		return p.FieldSets[i].FullName.Text < p.FieldSets[j].FullName.Text
	})
	sort.Slice(p.Fields, func(i, j int) bool {
		return p.Fields[i].FullName.Text < p.Fields[j].FullName.Text
	})
	sort.Slice(p.ValidationRules, func(i, j int) bool {
		return p.ValidationRules[i].FullName.Text < p.ValidationRules[j].FullName.Text
	})
	sort.Slice(p.ListViews, func(i, j int) bool {
		return p.ListViews[i].FullName.Text < p.ListViews[j].FullName.Text
	})
}
