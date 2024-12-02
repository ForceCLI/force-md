package objects

type ActionOverrideFilter func(ActionOverride) bool

func (o *CustomObject) GetActionOverrides(filters ...ActionOverrideFilter) []ActionOverride {
	var actions []ActionOverride
ACTIONS:
	for _, a := range o.ActionOverrides {
		for _, filter := range filters {
			if !filter(a) {
				continue ACTIONS
			}
		}
		actions = append(actions, a)
	}
	return actions
}
