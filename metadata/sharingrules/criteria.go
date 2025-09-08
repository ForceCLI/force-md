package sharingrules

import (
	"strings"

	"github.com/pkg/errors"
)

func (p *SharingRules) DeleteCriteriaRule(ruleName string) error {
	found := false
	newPerms := p.SharingCriteriaRules[:0]
	for _, f := range p.SharingCriteriaRules {
		if f.FullName == ruleName {
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

// UsesField checks if a criteria rule uses the specified field (case-insensitive)
func (r CriteriaRule) UsesField(fieldName string) bool {
	for _, item := range r.CriteriaItems {
		if strings.EqualFold(item.Field.Text, fieldName) {
			return true
		}
	}
	return false
}
