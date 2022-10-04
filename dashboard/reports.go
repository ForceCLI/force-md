package dashboard

func (d *Dashboard) GetReports() []string {
	var reports []string
	if d.DashboardGridLayout != nil {
		for _, c := range d.DashboardGridLayout.DashboardGridComponents {
			reports = append(reports, c.DashboardComponent.Report)
		}
	}
	if d.LeftSection != nil {
		for _, c := range d.LeftSection.Components {
			reports = append(reports, c.Report)
		}
	}
	if d.MiddleSection != nil {
		for _, c := range d.MiddleSection.Components {
			reports = append(reports, c.Report)
		}
	}
	if d.RightSection != nil {
		for _, c := range d.RightSection.Components {
			reports = append(reports, c.Report)
		}
	}
	return reports
}
