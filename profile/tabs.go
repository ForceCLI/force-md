package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) DeleteTabVisibility(tabName string) error {
	found := false
	for i, f := range p.TabVisibilities {
		if f.Tab.Text == tabName {
			p.TabVisibilities = append(p.TabVisibilities[:i], p.TabVisibilities[i+1:]...)
			found = true
		}
	}
	if !found {
		return errors.New("tab not found")
	}
	return nil
}
