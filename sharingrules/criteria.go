package sharingrules

import (
	"github.com/pkg/errors"
)

func (p *SharingRules) DeleteCriteriaRule(ruleName string) error {
	found := false
	newPerms := p.SharingCriteriaRules[:0]
	for _, f := range p.SharingCriteriaRules {
		if f.FullName.Text == ruleName {
			found = true
		} else {
			newPerms = append(newPerms, f)
		}
	}
	if !found {
		return errors.New("rule not found")
	}
	p.SharingCriteriaRules = newPerms
	return nil
}
