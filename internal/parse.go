package internal

import (
	"bytes"
	"io"
	"os"

	"github.com/nbio/xml"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type MetadataPointer interface {
	// MetaCheck should have a pointer receiver.  This ensures that functions
	// that take a MetadataPointer receive a pointer.
	MetaCheck()
}

func ParseMetadataXmlIfPossible(i MetadataPointer, path string) ([]byte, error) {
	var f *os.File
	var err error
	if path == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(path)
		if err != nil {
			return nil, errors.Wrap(err, "opening file")
		}
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "reading file")
	}
	r := bytes.NewReader(contents)
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = true

	if err := dec.Decode(i); err != nil {
		return contents, errors.Wrap(err, "decoding xml")
	}
	return contents, nil
}

func ParseMetadataXml(i MetadataPointer, path string) error {
	var r *os.File
	var err error
	if path == "-" {
		r = os.Stdin
	} else {
		r, err = os.Open(path)
		if err != nil {
			return errors.Wrap(err, "opening file")
		}
	}
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = true

	if err := dec.Decode(i); err != nil {
		return errors.Wrap(err, "parsing xml in "+path)
	}
	return nil
}
