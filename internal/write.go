package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const declaration = `<?xml version="1.0" encoding="UTF-8"?>`

func WriteToFile(t interface{}, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	defer f.Close()
	fmt.Fprintln(f, declaration)
	b, err := xml.MarshalIndent(t, "", "    ")
	if err != nil {
		return errors.Wrap(err, "serializing metadata")
	}
	b = htmlEntities(b)
	if _, err = f.Write(b); err != nil {
		return errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(f, "")
	return nil
}

func htmlEntities(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("&39;"), []byte("&apos;"))
}
