package profile

import (
	"fmt"
	"sort"
)

func (p *Profile) CloneFieldPermissions(src, dest string) error {
	for _, f := range p.FieldPermissions {
		if f.Field.Text == dest {
			return fmt.Errorf("%s field already exists", dest)
		}
	}
	found := false
	for _, f := range p.FieldPermissions {
		if f.Field.Text == src {
			found = true
			clone := FieldPermissions{}
			clone.Editable.Text = f.Editable.Text
			clone.Readable.Text = f.Readable.Text
			clone.Field.Text = dest
			p.FieldPermissions = append(p.FieldPermissions, clone)
		}
	}
	if !found {
		return fmt.Errorf("source field %s not found", src)
	}
	sort.Slice(p.FieldPermissions, func(i, j int) bool {
		return p.FieldPermissions[i].Field.Text < p.FieldPermissions[j].Field.Text
	})
	return nil
}
