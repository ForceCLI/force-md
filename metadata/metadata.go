package metadata

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/nbio/xml"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type MetadataType = string

// Format represents the metadata format (source or metadata)
type Format string

const (
	SourceFormat   Format = "source"   // SFDX source format
	MetadataFormat Format = "metadata" // MDAPI metadata format
)

// FileContents represents a file's path and content
type FileContents struct {
	Path    string
	Content []byte
}

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

// FilesGenerator is an optional interface that metadata types can implement
// to provide custom file generation logic
type FilesGenerator interface {
	// Files returns a map of files that make up this metadata component
	// The map key is the relative file path, the value contains the content
	Files(format Format) (map[string][]byte, error)
}

// DefaultFiles provides a default implementation of the Files method for metadata types
// that consist of a single XML file without associated code files
func DefaultFiles(m RegisterableMetadata, format Format) (map[string][]byte, error) {
	// Get the metadata name
	name := m.GetMetadataInfo().Name()
	if name == "" {
		return nil, fmt.Errorf("metadata name is empty")
	}

	// Get the metadata type to determine directory and suffix
	metadataType := m.Type()

	// Import needed for repo functions
	// This will need to be handled differently - circular import issue
	// For now, we'll require each type to implement its own Files() method
	return nil, fmt.Errorf("Files() method not implemented for %s", metadataType)
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
