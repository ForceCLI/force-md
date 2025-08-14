package objects

import (
	"fmt"
	"sort"

	"github.com/ForceCLI/force-md/internal"
	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (p *CustomObject) Tidy() {
	sort.Slice(p.FieldSets, func(i, j int) bool {
		return p.FieldSets[i].FullName < p.FieldSets[j].FullName
	})
	p.Fields.Tidy()
	p.ValidationRules.Tidy()
	sort.Slice(p.ListViews, func(i, j int) bool {
		return p.ListViews[i].FullName.Text < p.ListViews[j].FullName.Text
	})
	sort.Slice(p.CompactLayouts, func(i, j int) bool {
		return p.CompactLayouts[i].FullName.Text < p.CompactLayouts[j].FullName.Text
	})
	sort.Slice(p.RecordTypes, func(i, j int) bool {
		return p.RecordTypes[i].FullName < p.RecordTypes[j].FullName
	})
	sort.Slice(p.WebLinks, func(i, j int) bool {
		return p.WebLinks[i].FullName < p.WebLinks[j].FullName
	})
	sort.Slice(p.SharingReasons, func(i, j int) bool {
		return p.SharingReasons[i].FullName < p.SharingReasons[j].FullName
	})
	sort.Slice(p.BusinessProcesses, func(i, j int) bool {
		return p.BusinessProcesses[i].FullName.Text < p.BusinessProcesses[j].FullName.Text
	})
}

func (fields FieldList) Tidy() {
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].FullName < fields[j].FullName
	})
	for _, f := range fields {
		if f.Formula != nil {
			formatted, err := formatter.Format(f.Formula.String())
			if err != nil {
				log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
			} else {
				f.Formula.Text = internal.FormulaEscaper.Replace(formatted)
			}
		}
	}
}
