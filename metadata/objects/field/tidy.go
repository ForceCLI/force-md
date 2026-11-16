package field

import (
	"fmt"

	"github.com/octoberswimmer/sformula/formatter"
	log "github.com/sirupsen/logrus"
)

func (f Field) Tidy() {
	if f.Formula != nil {
		if err := f.Formula.Reformat(formatter.Format); err != nil {
			log.Warn(fmt.Sprintf("error formatting %s: %s", f.FullName, err.Error()))
		}
	}
}
