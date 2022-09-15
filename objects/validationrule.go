package objects

import (
	"github.com/pkg/errors"
)

func (p *CustomObject) GetValidationRules() []Rule {
	return p.ValidationRules
}

func (p *CustomObject) DeleteRule(ruleName string) error {
	found := false
	newRules := p.ValidationRules[:0]
	for _, f := range p.ValidationRules {
		if f.FullName == ruleName {
			found = true
		} else {
			newRules = append(newRules, f)
		}
	}
	if !found {
		return errors.New("rule not found")
	}
	p.ValidationRules = newRules
	return nil
}
