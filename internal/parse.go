package internal

import (
	"encoding/xml"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type MetadataPointer interface {
	// MetaCheck should have a pointer receiver.  This ensures that functions
	// that take a MetadataPointer receive a pointer.
	MetaCheck()
}

func ParseMetadataXml(i MetadataPointer, r io.Reader) error {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	if err := dec.Decode(i); err != nil {
		return errors.Wrap(err, "parsing xml")
	}
	return nil
}
