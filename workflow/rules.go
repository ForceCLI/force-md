package workflow

type RuleFilter func(Rule) bool

func (w *Workflow) GetRules(filters ...RuleFilter) []Rule {
	var rules []Rule
RULES:
	for _, r := range w.Rules {
		for _, filter := range filters {
			if !filter(r) {
				continue RULES
			}
		}
		rules = append(rules, r)
	}
	return rules
}

func (w *Workflow) GetAlerts() []Alert {
	return w.Alerts
}

func (w *Workflow) GetFieldUpdates() []FieldUpdate {
	return w.FieldUpdates
}
