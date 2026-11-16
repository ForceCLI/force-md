package workflow

import (
	"fmt"

	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (f Workflow) Tidy() {
	for i, rule := range f.Rules {
		if rule.Formula == nil {
			continue
		}
		if err := f.Rules[i].Formula.Reformat(formatter.FormatFlowFormula); err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", rule.FullName, err.Error()))
		}
	}
}
