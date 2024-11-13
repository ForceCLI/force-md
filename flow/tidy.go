package flow

import (
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (f Flow) Tidy() {
	for i, formula := range f.Formulas {
		formatted, err := formatter.FormatFlowFormula(formula.Expression.String())
		if err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", formula.Name, err.Error()))
		} else {
			f.Formulas[i].Expression.Text = internal.FormulaEscaper.Replace(formatted)
		}
	}
	if f.Start != nil && f.Start.FilterFormula != nil {
		formatted, err := formatter.FormatFlowFormula(f.Start.FilterFormula.String())
		if err != nil {
			log.Warn(fmt.Sprintf("error formatting Start filter formula: %s", err.Error()))
		} else {
			f.Start.FilterFormula.Text = internal.FormulaEscaper.Replace(formatted)
		}
	}
}
