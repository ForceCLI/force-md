package internal

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const declaraction = `<?xml version="1.0" encoding="UTF-8"?>`

func WriteToFile(t interface{}, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	defer f.Close()
	fmt.Fprintln(f, declaraction)
	b, err := xml.MarshalIndent(t, "", "    ")
	if err != nil {
		return errors.Wrap(err, "serializing metadata")
	}
	if _, err = f.Write(b); err != nil {
		return errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(f, "")
	return nil
}
