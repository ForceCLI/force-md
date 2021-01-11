package profile

import (
	"github.com/pkg/errors"
)

func (p *Profile) DeleteApplicationVisibility(applicationName string) error {
	found := false
	newApps := p.ApplicationVisibilities[:0]
	for _, f := range p.ApplicationVisibilities {
		if f.Application.Text == applicationName {
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
