package dashboard

import (
	"errors"
	"strings"
)

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

func (d *Dashboard) DeleteReport(report string) error {
	found := false
	if d.DashboardGridLayout != nil {
		newComponents := d.DashboardGridLayout.DashboardGridComponents[:0]
		for _, f := range d.DashboardGridLayout.DashboardGridComponents {
			if strings.ToLower(f.DashboardComponent.Report) == strings.ToLower(report) {
				found = true
			} else {
				newComponents = append(newComponents, f)
			}
		}
		d.DashboardGridLayout.DashboardGridComponents = newComponents
	}
	if !found {
		return errors.New("report not found")
	}
	return nil
}
