package workflow

import (
	"strings"

	"github.com/cwarden/mergo"
	"github.com/pkg/errors"
)

func (o *Workflow) UpdateAlert(alertName string, updates Alert) error {
	found := false
	for i, f := range o.Alerts {
		if strings.ToLower(f.FullName) == strings.ToLower(alertName) {
			found = true
			if err := mergo.Merge(&updates, f, mergo.WithNoOverrideEmptyStructValues); err != nil {
				return errors.Wrap(err, "merging field updates")
			}
			o.Alerts[i] = updates
		}
	}
	if !found {
		return errors.New("alert not found")
	}
	return nil
}

func (o *Workflow) DeleteAlert(alertName string) error {
	found := false
	newAlerts := o.Alerts[:0]
	for _, a := range o.Alerts {
		if strings.ToLower(a.FullName) == strings.ToLower(alertName) {
			found = true
		} else {
			newAlerts = append(newAlerts, a)
		}
	}
	if !found {
		return errors.New("alert not found")
	}
	o.Alerts = newAlerts
	return nil
}
