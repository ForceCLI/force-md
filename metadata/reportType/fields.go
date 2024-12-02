package reportType

import (
	"fmt"
	"strings"

	. "github.com/ForceCLI/force-md/general"
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

func (o *ReportType) DeleteField(field, table string) error {
	found := false
	for s, section := range o.Sections {
		newFields := section.Columns[:0]
		for _, f := range section.Columns {
			fieldMatch := strings.ToLower(f.Field) == strings.ToLower(field)
			tableMatch := table == "" || strings.ToLower(f.Table) == strings.ToLower(table)
			if !fieldMatch || !tableMatch {
				newFields = append(newFields, f)
			} else {
				found = true
			}
		}
		o.Sections[s].Columns = newFields
	}
	if !found {
		return fmt.Errorf("field not found: %s", field)
	}
	return nil
}
