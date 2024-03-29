package dashboard

import (
	. "github.com/ForceCLI/force-md/general"
)

func (o *Dashboard) UpdateRunningUser(user string) {
	o.RunningUser = &TextLiteral{
		Text: user,
	}
}

func (o *Dashboard) UpdateDashboardType(dashboardType string) {
	o.DashboardType = &TextLiteral{
		Text: dashboardType,
	}
}
