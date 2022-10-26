package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var ConvertNumericXMLEntities = true

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
	b = selfClosing(b)
	if ConvertNumericXMLEntities {
		b = htmlEntities(b)
	}
	if _, err = f.Write(b); err != nil {
		return errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(f, "")
	return nil
}

func htmlEntities(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte("&#39;"), []byte("&apos;"))
	b = bytes.ReplaceAll(b, []byte("&#34;"), []byte("&quot;"))
	return b
}

// Make empty tags self-closing
func selfClosing(b []byte) []byte {
	emptyTag := regexp.MustCompile(`(?:</?(\w+)>){2}`)
	return emptyTag.ReplaceAll(b, []byte(`<$1/>`))
}
