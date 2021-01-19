package permissionset

import (
	"sort"
)

func (p *PermissionSet) AddTab(tabName string) {
	p.TabSettings = append(p.TabSettings, TabSettings{
		Tab:        tabName,
		Visibility: "Visible",
	})
	sort.Slice(p.TabSettings, func(i, j int) bool {
		return p.TabSettings[i].Tab < p.TabSettings[j].Tab
	})
}
