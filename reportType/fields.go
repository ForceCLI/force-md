package reportType

import (
	"fmt"

	. "github.com/octoberswimmer/force-md/general"
)

func (o *ReportType) GetFields() []Field {
	var fields []Field
	for _, s := range o.Sections {
		for _, f := range s.Columns {
			fields = append(fields, f)
		}
	}
	return fields
}

func (o *ReportType) AddField(sectionName, table, field string) error {
	for s, section := range o.Sections {
		if section.MasterLabel != sectionName {
			continue
		}
		for _, f := range section.Columns {
			if f.Table == table && f.Field == field {
				return fmt.Errorf("field already exists: %s.%s", table, field)
			}
		}
		o.Sections[s].Columns = append(o.Sections[s].Columns, Field{
			CheckedByDefault: FalseText,
			Table:            table,
			Field:            field,
		})
		o.Sections[s].Columns.Tidy()
		return nil
	}
	return fmt.Errorf("section not found: %s", sectionName)
}
