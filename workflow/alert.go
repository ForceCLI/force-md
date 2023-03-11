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
