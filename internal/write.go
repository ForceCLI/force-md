package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var ConvertNumericXMLEntities = true

const declaration = `<?xml version="1.0" encoding="UTF-8"?>`

var emptyRoot = regexp.MustCompile(`^<\w+[^<>]*/>$`)

func DisableIndent(e *xml.Encoder) {
	e.Indent("", "")
}

func EnableIndent(e *xml.Encoder) {
	e.Indent("", "    ")
}

// EmptyElementKeeper is implemented by metadata types that Salesforce writes
// with empty elements as <name></name> instead of <name/>.  Marshal leaves
// their empty elements open, except for those with attributes, such as
// <name xsi:nil="true"/>, which Salesforce still writes self-closed.
type EmptyElementKeeper interface {
	KeepEmptyElements()
}

// SelfClosedElementNamer is implemented by an EmptyElementKeeper whose
// complex elements Salesforce writes self-closed when they have no children,
// such as a flow's <value/>, while it writes its empty text elements open.
type SelfClosedElementNamer interface {
	SelfClosedElements() []string
}

func Marshal(t interface{}) ([]byte, error) {
	var b bytes.Buffer
	w := &b

	fmt.Fprintln(w, declaration)
	m, err := xml.MarshalIndent(t, "", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "serializing metadata")
	}
	if _, ok := t.(EmptyElementKeeper); ok {
		m = SelfClosingWithAttributes(m)
		if n, ok := t.(SelfClosedElementNamer); ok {
			m = SelfClosingNamed(m, n.SelfClosedElements())
		}
	} else {
		m = SelfClosing(m)
	}
	m = literalWhitespace(m)
	if ConvertNumericXMLEntities {
		m = htmlEntities(m)
	}
	if _, err = w.Write(m); err != nil {
		return nil, errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(w, "")
	// Salesforce follows an empty, self-closed root element with a blank line
	if emptyRoot.Match(m) {
		fmt.Fprintln(w, "")
	}
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

// encoding/xml escapes newlines and tabs in character data as &#xA; and
// &#x9;.  Salesforce writes them as literal characters.
func literalWhitespace(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte("&#xA;"), []byte("\n"))
	b = bytes.ReplaceAll(b, []byte("&#x9;"), []byte("\t"))
	return b
}

// Make empty tags self-closing
func SelfClosing(b []byte) []byte {
	emptyTag := regexp.MustCompile(`<(\w+)(\s*[^>]*)>\s*</[^>]+>`)
	return emptyTag.ReplaceAll(b, []byte(`<$1$2/>`))
}

// Make empty tags with the given names self-closing
func SelfClosingNamed(b []byte, names []string) []byte {
	if len(names) == 0 {
		return b
	}
	quoted := make([]string, len(names))
	for i, n := range names {
		quoted[i] = regexp.QuoteMeta(n)
	}
	emptyTag := regexp.MustCompile(`<(` + strings.Join(quoted, "|") + `)>\s*</[^>]+>`)
	return emptyTag.ReplaceAll(b, []byte(`<$1/>`))
}

// Make empty tags that have attributes self-closing
func SelfClosingWithAttributes(b []byte) []byte {
	emptyTag := regexp.MustCompile(`<(\w+)(\s+[^>]*[^/>])>\s*</[^>]+>`)
	return emptyTag.ReplaceAll(b, []byte(`<$1$2/>`))
}
