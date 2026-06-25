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
		defer f.Close()
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "reading file")
	}
	r := bytes.NewReader(escapeBareAmpersands(contents))
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
		defer f.Close()
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "reading file")
	}
	r := bytes.NewReader(escapeBareAmpersands(contents))
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

var (
	escapedAmp   = []byte("&amp;")
	cdataOpen    = []byte("<![CDATA[")
	cdataClose   = []byte("]]>")
	commentOpen  = []byte("<!--")
	commentClose = []byte("-->")
)

// predefinedEntities are the five entity references an XML parser must
// recognize even without a DTD.
var predefinedEntities = [][]byte{
	[]byte("&amp;"),
	[]byte("&lt;"),
	[]byte("&gt;"),
	[]byte("&apos;"),
	[]byte("&quot;"),
}

// escapeBareAmpersands rewrites ampersands that do not introduce a valid XML
// entity or character reference into &amp;. Salesforce can emit metadata
// containing bare & characters (e.g. picklist values and labels like "R&D" or
// "a & b"); a real org accepts and deploys such files even though they are not
// strictly well-formed XML. Normalizing the stray ampersands lets the strict
// decoder read them the way Salesforce does — as literal & characters — without
// relaxing any of the decoder's other well-formedness checks. Ampersands inside
// CDATA sections and comments are left untouched, where a bare & is already
// legal and an escape would corrupt the content.
func escapeBareAmpersands(contents []byte) []byte {
	if bytes.IndexByte(contents, '&') < 0 {
		return contents
	}
	out := make([]byte, 0, len(contents)+16)
	for i := 0; i < len(contents); {
		if n := leadingOpaqueRegion(contents[i:]); n > 0 {
			out = append(out, contents[i:i+n]...)
			i += n
			continue
		}
		if contents[i] == '&' && referenceLen(contents[i:]) == 0 {
			out = append(out, escapedAmp...)
			i++
			continue
		}
		out = append(out, contents[i])
		i++
	}
	return out
}

// leadingOpaqueRegion returns the byte length of a CDATA section or comment at
// the start of s, whose contents must be copied through verbatim, or 0 if s
// does not begin with one. An unterminated region extends to the end of s; the
// decoder reports the malformed input.
func leadingOpaqueRegion(s []byte) int {
	switch {
	case bytes.HasPrefix(s, cdataOpen):
		if end := bytes.Index(s[len(cdataOpen):], cdataClose); end >= 0 {
			return len(cdataOpen) + end + len(cdataClose)
		}
	case bytes.HasPrefix(s, commentOpen):
		if end := bytes.Index(s[len(commentOpen):], commentClose); end >= 0 {
			return len(commentOpen) + end + len(commentClose)
		}
	default:
		return 0
	}
	return len(s)
}

// referenceLen returns the byte length of the XML entity or character reference
// at the start of s, or 0 if s does not begin with one. s must begin with '&'.
func referenceLen(s []byte) int {
	for _, e := range predefinedEntities {
		if bytes.HasPrefix(s, e) {
			return len(e)
		}
	}
	// Numeric character reference: &#DDDD; (decimal) or &#xHHHH; (hex).
	if len(s) < 4 || s[1] != '#' {
		return 0
	}
	i := 2
	isDigit := func(b byte) bool { return b >= '0' && b <= '9' }
	if s[i] == 'x' {
		i++
		isDigit = func(b byte) bool {
			return b >= '0' && b <= '9' || b >= 'a' && b <= 'f' || b >= 'A' && b <= 'F'
		}
	}
	start := i
	for i < len(s) && isDigit(s[i]) {
		i++
	}
	if i > start && i < len(s) && s[i] == ';' {
		return i + 1
	}
	return 0
}
