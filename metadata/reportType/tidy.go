package reportType

import (
	"sort"
)

func (r *ReportType) Tidy() {
	for i, _ := range r.Sections {
		r.Sections[i].Columns.Tidy()
	}
}

func (c FieldList) Tidy() {
	sort.Slice(c, func(i, j int) bool {
		return c[i].Field < c[j].Field
	})
}
