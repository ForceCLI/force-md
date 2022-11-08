package application

import (
	"errors"
)

type ProfileActionOverrideFilter func(ProfileActionOverride) bool

func (o *CustomApplication) GetProfileActionOverrides(filters ...ProfileActionOverrideFilter) []ProfileActionOverride {
	var actions []ProfileActionOverride
ACTIONS:
	for _, a := range o.ProfileActionOverrides {
		for _, filter := range filters {
			if !filter(a) {
				continue ACTIONS
			}
		}
		actions = append(actions, a)
	}
	return actions
}

func (o *CustomApplication) DeleteActionOverrides(filters ...ProfileActionOverrideFilter) error {
	found := false
	newActions := o.ProfileActionOverrides[:0]
ACTIONS:
	for _, a := range o.ProfileActionOverrides {
		for _, filter := range filters {
			if !filter(a) {
				newActions = append(newActions, a)
				continue ACTIONS
			}
		}
		found = true
	}
	if !found {
		return errors.New("no actions found")
	}
	o.ProfileActionOverrides = newActions
	return nil
}
