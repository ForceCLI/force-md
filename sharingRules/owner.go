package sharingRules

import (
	"github.com/pkg/errors"
)

func (p *SharingRules) DeleteOwnerRule(ruleName string) error {
	found := false
	newPerms := p.SharingOwnerRules[:0]
	for _, f := range p.SharingOwnerRules {
		if f.FullName.Text == ruleName {
			found = true
		} else {
			newPerms = append(newPerms, f)
		}
	}
	if !found {
		return errors.New("rule not found")
	}
	p.SharingOwnerRules = newPerms
	return nil
}
