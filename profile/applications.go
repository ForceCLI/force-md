package profile

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"

	. "github.com/octoberswimmer/force-md/general"
)

var ApplicationExistsError = errors.New("application already exists")

type ApplicationFilter func(ApplicationVisibility) bool

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
		return TrueText
	}
	return FalseText
}

func (p *Profile) AddApplicationVisibility(appName string, defaultApp bool) error {
	for _, f := range p.ApplicationVisibilities {
		if f.Application == appName {
			return ApplicationExistsError
		}
	}

	p.ApplicationVisibilities = append(p.ApplicationVisibilities, ApplicationVisibility{
		Application: appName,
		Visible:     TrueText,
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

func (p *Profile) SetApplicationVisibility(applicationName string, updates ApplicationVisibility) error {
	found := false
	for i, f := range p.ApplicationVisibilities {
		if strings.ToLower(f.Application) == strings.ToLower(applicationName) {
			found = true
			if err := mergo.Merge(&updates, f); err != nil {
				return errors.Wrap(err, "merging permissions")
			}
			p.ApplicationVisibilities[i] = updates
		}
	}
	if !found {
		return fmt.Errorf("application not found: %s", applicationName)
	}
	return nil
}
