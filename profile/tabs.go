package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) DeleteTabVisibility(tabName string) error {
	found := false
	newTabs := p.TabVisibilities[:0]
	for _, f := range p.TabVisibilities {
		if f.Tab.Text == tabName {
			found = true
		} else {
			newTabs = append(newTabs, f)
		}
	}
	p.TabVisibilities = newTabs
	if !found {
		return errors.New("tab not found")
	}
	return nil
}
