package dashboard

import (
	. "github.com/ForceCLI/force-md/general"
)

func (o *Dashboard) UpdateRunningUser(user string) {
	o.RunningUser = NewTextLiteral(user)
}

func (o *Dashboard) UpdateDashboardType(dashboardType string) {
	o.DashboardType = NewTextLiteral(dashboardType)
}
