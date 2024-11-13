package validationrule

import (
	"fmt"
	"sort"

	"github.com/ForceCLI/force-md/internal"
	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (rules ValidationRuleList) Tidy() {
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].FullName < rules[j].FullName
	})
	for _, f := range rules {
		formatted, err := formatter.Format(f.ErrorConditionFormula.String())
		if err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
		} else {
			f.ErrorConditionFormula.Text = internal.FormulaEscaper.Replace(formatted)
		}
	}
}

func (f ValidationRule) Tidy() {
	formatted, err := formatter.Format(f.ErrorConditionFormula.String())
	if err != nil {
		log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
	} else {
		f.ErrorConditionFormula.Text = internal.FormulaEscaper.Replace(formatted)
	}
}
