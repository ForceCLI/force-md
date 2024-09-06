package general

import (
	"fmt"

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

type Nameable interface {
	GetName() string
}

func RemoveDuplicates[T Nameable, S ~[]T](slice *S) {
	if len(*slice) == 0 {
		return
	}

	lastUniqueIndex := 0

	for i := 1; i < len(*slice); i++ {
		// If the current element is not a duplicate, move it to the next position after the last unique element
		if (*slice)[i].GetName() != (*slice)[lastUniqueIndex].GetName() {
			lastUniqueIndex++
			(*slice)[lastUniqueIndex] = (*slice)[i]
		} else {
			fmt.Println("omitting duplicate ", (*slice)[i].GetName())
		}
	}

	// Slice the original slice to the correct length of unique elements
	*slice = (*slice)[:lastUniqueIndex+1]
}
