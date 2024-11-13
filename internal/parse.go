package internal

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nbio/xml"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type MetadataInfo struct {
	path     string
	name     string
	contents []byte
}

func (m MetadataInfo) Name() string {
	return m.name
}

func (m MetadataInfo) Path() string {
	return m.path
}

func (m MetadataInfo) Contents() []byte {
	return m.contents
}

func (m MetadataInfo) GetMetadataInfo() MetadataInfo {
	return m
}

type RegisterableMetadata interface {
	MetadataPointer
	GetMetadataInfo() MetadataInfo
	Type() MetadataType
}

type MetadataPointer interface {
	// SetMetadata should have a pointer receiver.  This ensures that functions
	// that take a MetadataPointer receive a pointer.
	SetMetadata(MetadataInfo)
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
	meta.path = path
	meta.contents = contents
	name := strings.TrimSuffix(filepath.Base(path), "-meta.xml")
	meta.name = strings.TrimSuffix(name, filepath.Ext(name))
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
	meta.path = path
	meta.contents = contents
	name := strings.TrimSuffix(filepath.Base(path), "-meta.xml")
	meta.name = strings.TrimSuffix(name, filepath.Ext(name))
	i.SetMetadata(meta)

	return nil
}
