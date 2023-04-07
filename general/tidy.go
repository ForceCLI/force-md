package general

import (
	"github.com/pkg/errors"

	"github.com/ForceCLI/force-md/internal"
)

type Tidyable interface {
	Tidy()
}

func Tidy(t Tidyable, path string) error {
	t.Tidy()
	if err := internal.WriteToFile(t, path); err != nil {
		return errors.Wrap(err, "writing to file")
	}
	return nil
}
