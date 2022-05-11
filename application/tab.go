package application

import (
	"errors"
	"sort"
	"strings"

	. "github.com/octoberswimmer/force-md/general"
)

var TabExistsError = errors.New("tab already exists")

func (p *CustomApplication) AddTab(tabName string) error {
	for _, t := range p.Tabs {
		if strings.ToLower(t.Text) == strings.ToLower(tabName) {
			return TabExistsError
		}
	}
	p.Tabs = append(p.Tabs, TextLiteral{
		Text: tabName,
	})
	sort.Slice(p.Tabs, func(i, j int) bool {
		return p.Tabs[i].Text < p.Tabs[j].Text
	})
	return nil
}

func (p *CustomApplication) DeleteTab(tabName string) error {
	found := false
	newTabs := p.Tabs[:0]
	for _, f := range p.Tabs {
		if strings.ToLower(f.Text) == strings.ToLower(tabName) {
			found = true
		} else {
			newTabs = append(newTabs, f)
		}
	}
	if !found {
		return errors.New("tab not found")
	}
	p.Tabs = newTabs
	return nil
}

func (p *CustomApplication) GetTabs() []TextLiteral {
	return p.Tabs
}
