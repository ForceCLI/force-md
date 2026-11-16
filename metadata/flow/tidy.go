package flow

import (
	"fmt"

	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (f Flow) Tidy() {
	// Salesforce saves a flow whose formulas fail validation with status
	// InvalidDraft, so a formula in such a flow that does not parse is
	// expected and is logged at debug level instead of as a warning.
	logFormatError := log.Warn
	if f.Status != nil && f.Status.Text == "InvalidDraft" {
		logFormatError = log.Debug
	}
	for i, formula := range f.Formulas {
		if err := f.Formulas[i].Expression.Reformat(formatter.FormatFlowFormula); err != nil {
			logFormatError(fmt.Sprintf("error formatting %s: %s", formula.Name, err.Error()))
		}
	}
	if f.Start != nil && f.Start.FilterFormula != nil {
		if err := f.Start.FilterFormula.Reformat(formatter.FormatFlowFormula); err != nil {
			logFormatError(fmt.Sprintf("error formatting Start filter formula: %s", err.Error()))
		}
	}
}
