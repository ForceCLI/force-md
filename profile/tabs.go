package profile

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

var TabExistsError = errors.New("tab already exists")

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

func (p *Profile) AddTab(tabName string) error {
	for _, t := range p.TabVisibilities {
		if t.Tab == tabName {
			return TabExistsError
		}
	}
	p.TabVisibilities = append(p.TabVisibilities, TabVisibility{
		Tab:        tabName,
		Visibility: "DefaultOn",
	})
	sort.Slice(p.TabVisibilities, func(i, j int) bool {
		return p.TabVisibilities[i].Tab < p.TabVisibilities[j].Tab
	})
	return nil
}

func (p *Profile) SetTabVisibility(tabName string, visibility string) error {
	found := false
	for i, f := range p.TabVisibilities {
		if f.Tab == tabName {
			found = true
			p.TabVisibilities[i].Visibility = visibility
		}
	}
	if !found {
		return fmt.Errorf("tab not found: %s", tabName)
	}
	return nil
}

func (p *Profile) GetTabs() TabVisibilityList {
	return p.TabVisibilities
}
