package metadata

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type MetadataType = string

type MetadataPointer interface {
	// SetMetadata should have a pointer receiver.  This ensures that functions
	// that take a MetadataPointer receive a pointer.
	SetMetadata(MetadataInfo)
	NameFromPath(path string) MetadataObjectName
}

type RegisterableMetadata interface {
	MetadataPointer
	GetMetadataInfo() MetadataInfo
	Type() MetadataType
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

	meta := MetadataInfo{}
	meta.path = MetadataFilePath(path)
	meta.contents = contents
	name := i.NameFromPath(path)
	meta.name = name
	i.SetMetadata(meta)

	return contents, nil
}

func ParseMetadataXml(i MetadataPointer, path string) error {
	var f *os.File
	var err error
	if path == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(path)
		if err != nil {
			return errors.Wrap(err, "opening file")
		}
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "reading file")
	}
	r := bytes.NewReader(contents)
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = true

	if err := dec.Decode(i); err != nil {
		return errors.Wrap(err, "parsing xml in "+path)
	}

	meta := MetadataInfo{}
	meta.path = MetadataFilePath(path)
	meta.contents = contents
	name := i.NameFromPath(path)
	meta.name = name
	i.SetMetadata(meta)

	return nil
}
