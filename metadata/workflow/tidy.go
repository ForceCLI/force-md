package workflow

import (
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (f Workflow) Tidy() {
	for i, rule := range f.Rules {
		if rule.Formula == nil {
			continue
		}
		formatted, err := formatter.FormatFlowFormula(rule.Formula.String())
		if err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", rule.FullName, err.Error()))
		} else {
			f.Rules[i].Formula.Text = internal.FormulaEscaper.Replace(formatted)
		}
	}
}
