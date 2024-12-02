package workflow

import (
	"strings"

	"github.com/pkg/errors"
)

type RuleFilter func(Rule) bool
type AlertFilter func(Alert) bool

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

func (w *Workflow) GetAlerts(filters ...AlertFilter) []Alert {
	var alerts []Alert
ALERTS:
	for _, a := range w.Alerts {
		for _, filter := range filters {
			if !filter(a) {
				continue ALERTS
			}
		}
		alerts = append(alerts, a)
	}
	return alerts
}

func (w *Workflow) GetFieldUpdates() []FieldUpdate {
	return w.FieldUpdates
}

func (o *Workflow) DeleteRule(ruleName string) error {
	found := false
	newRules := o.Rules[:0]
	for _, a := range o.Rules {
		if strings.ToLower(a.FullName) == strings.ToLower(ruleName) {
			found = true
		} else {
			newRules = append(newRules, a)
		}
	}
	if !found {
		return errors.New("rule not found")
	}
	o.Rules = newRules
	return nil
}
