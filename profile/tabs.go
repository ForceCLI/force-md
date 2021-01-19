package profile

import (
	"sort"

	"github.com/pkg/errors"
)

func (p *Profile) DeleteTabVisibility(tabName string) error {
	found := false
	newTabs := p.TabVisibilities[:0]
	for _, f := range p.TabVisibilities {
		if f.Tab == tabName {
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

func (p *Profile) AddTab(tabName string) {
	p.TabVisibilities = append(p.TabVisibilities, TabVisibility{
		Tab:        tabName,
		Visibility: "DefaultOn",
	})
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab < p.TabVisibilities[j].Tab
	})
}
