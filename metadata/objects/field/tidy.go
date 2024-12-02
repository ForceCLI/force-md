package field

import (
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (f Field) Tidy() {
	if f.Formula != nil {
		formatted, err := formatter.Format(f.Formula.String())
		if err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
		} else {
			f.Formula.Text = internal.FormulaEscaper.Replace(formatted)
		}
	}
}
