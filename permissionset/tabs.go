package permissionset

import (
	"github.com/pkg/errors"
	"sort"
	"strings"
)

var TabExistsError = errors.New("tab already exists")

func (p *PermissionSet) AddTab(tabName string) error {
	for _, t := range p.TabSettings {
		if t.Tab == tabName {
			return TabExistsError
		}
	}
	p.TabSettings = append(p.TabSettings, TabSettings{
		Tab:        tabName,
		Visibility: "Visible",
	})
	sort.Slice(p.TabSettings, func(i, j int) bool {
		return p.TabSettings[i].Tab < p.TabSettings[j].Tab
	})
	return nil
}

func (t *TabSettings) IsVisible() bool {
	return strings.ToLower(t.Visibility) != "hidden"
}
