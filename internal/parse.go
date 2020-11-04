package internal

import (
	"encoding/xml"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

func ParseMetadataXml(i interface{}, r io.Reader) error {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	if err := dec.Decode(i); err != nil {
		return errors.Wrap(err, "parsing xml")
	}
	return nil
}
