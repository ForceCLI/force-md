package permissionset

import (
	"errors"

	. "github.com/octoberswimmer/force-md/general"
)

var ApplicationExistsError = errors.New("application already exists")

func (p *PermissionSet) AddApplicationVisibility(appName string) error {
	for _, f := range p.ApplicationVisibilities {
		if f.Application == appName {
			return ApplicationExistsError
		}
	}

	p.ApplicationVisibilities = append(p.ApplicationVisibilities, ApplicationVisibility{
		Application: appName,
		Visible:     TrueText,
	})
	p.ApplicationVisibilities.Tidy()
	return nil
}
