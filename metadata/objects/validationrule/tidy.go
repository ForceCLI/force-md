package validationrule

import (
	"fmt"
	"sort"

	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (rules ValidationRuleList) Tidy() {
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].FullName < rules[j].FullName
	})
	for _, f := range rules {
		if f.ErrorConditionFormula != nil {
			if err := f.ErrorConditionFormula.Reformat(formatter.Format); err != nil {
				log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
			}
		}
	}
}

func (f ValidationRule) Tidy() {
	if f.ErrorConditionFormula != nil {
		if err := f.ErrorConditionFormula.Reformat(formatter.Format); err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
		}
	}
}
