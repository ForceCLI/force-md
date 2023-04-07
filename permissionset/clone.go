package permissionset

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func (p *PermissionSet) CloneFieldPermissions(src, dest string) error {
	found := false
	for _, f := range p.FieldPermissions {
		if strings.ToLower(dest) == strings.ToLower(f.Field) {
			return errors.New("field already exists")
		}
		if strings.ToLower(f.Field) == strings.ToLower(src) {
			found = true
			clone := FieldPermissions{}
			clone.Editable.Text = f.Editable.Text
			clone.Readable.Text = f.Readable.Text
			clone.Field = dest
			p.FieldPermissions = append(p.FieldPermissions, clone)
		}
	}
	if !found {
		return errors.New("source field not found")
	}
	sort.Slice(p.FieldPermissions, func(i, j int) bool {
		return p.FieldPermissions[i].Field < p.FieldPermissions[j].Field
	})
	return nil
}
