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

func DisableIndent(e *xml.Encoder) {
	e.Indent("", "")
}

func EnableIndent(e *xml.Encoder) {
	e.Indent("", "    ")
}

func Marshal(t interface{}) ([]byte, error) {
	var b bytes.Buffer
	w := &b

	fmt.Fprintln(w, declaration)
	m, err := xml.MarshalIndent(t, "", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "serializing metadata")
	}
	m = SelfClosing(m)
	if ConvertNumericXMLEntities {
		m = htmlEntities(m)
	}
	if _, err = w.Write(m); err != nil {
		return nil, errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(w, "")
	return b.Bytes(), nil
}

func WriteToFile(t interface{}, fileName string) error {
	var f *os.File
	var err error
	if fileName == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(fileName)
		if err != nil {
			return errors.Wrap(err, "opening file")
		}
	}
	defer f.Close()
	b, err := Marshal(t)
	if err != nil {
		return errors.Wrap(err, "marshalling")
	}
	if _, err = f.Write(b); err != nil {
		return errors.Wrap(err, "writing xml")
	}
	return nil
}

func htmlEntities(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte("&#39;"), []byte("&apos;"))
	b = bytes.ReplaceAll(b, []byte("&#34;"), []byte("&quot;"))
	return b
}

// Make empty tags self-closing
func SelfClosing(b []byte) []byte {
	emptyTag := regexp.MustCompile(`<(\w+)(\s*[^>]*)>\s*</[^>]+>`)
	return emptyTag.ReplaceAll(b, []byte(`<$1$2/>`))
}
