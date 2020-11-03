package reporttype

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func (r *ReportType) RemoveField(fieldName string) error {
	pieces := strings.Split(fieldName, ".")
	if len(pieces) != 2 {
		return errors.New(fmt.Sprintf("Invalid field name %s.  Should be in <object>.<field> format", fieldName))
	}
	object := pieces[0]
	field := pieces[1]
	for i, s := range r.Sections {
		for j, c := range s.Columns {
			if c.Field.Text == field && c.Table.Text == object {
				r.Sections[i].Columns = append(s.Columns[:j], s.Columns[j+1:]...)
				return nil
			}
		}
	}
	return errors.New("field not found")
}
