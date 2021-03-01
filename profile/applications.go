package profile

import (
	"github.com/pkg/errors"
)

type ApplicationFilter func(ApplicationVisibility) bool

var falseBooleanText = BooleanText{
	Text: "false",
}

var trueBooleanText = BooleanText{
	Text: "true",
}

func (p *Profile) DeleteApplicationVisibility(applicationName string) error {
	found := false
	newApps := p.ApplicationVisibilities[:0]
	for _, f := range p.ApplicationVisibilities {
		if f.Application == applicationName {
			found = true
		} else {
			newApps = append(newApps, f)
		}
	}
	if !found {
		return errors.New("application not found")
	}
	p.ApplicationVisibilities = newApps
	return nil
}

func boolToText(v bool) BooleanText {
	if v {
		return trueBooleanText
	}
	return falseBooleanText
}

func (p *Profile) AddApplicationVisibility(appName string, defaultApp bool) error {
	for _, f := range p.ApplicationVisibilities {
		if f.Application == appName {
			return errors.New("application already exists")
		}
	}

	p.ApplicationVisibilities = append(p.ApplicationVisibilities, ApplicationVisibility{
		Application: appName,
		Visible:     trueBooleanText,
		Default:     boolToText(defaultApp),
	})
	p.ApplicationVisibilities.Tidy()
	return nil
}

func (p *Profile) GetApplications(filters ...ApplicationFilter) []ApplicationVisibility {
	var applications []ApplicationVisibility
APPS:
	for _, v := range p.ApplicationVisibilities {
		for _, filter := range filters {
			if !filter(v) {
				continue APPS
			}
		}
		applications = append(applications, v)
	}
	return applications

}
