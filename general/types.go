package general

import (
	"encoding/xml"
	"html"
	"strconv"
	"strings"

	"github.com/ForceCLI/force-md/metadata"
)

type Metadata = metadata.MetadataPointer

type TextLiteral struct {
	Text string `xml:",innerxml"`
}

func (t *TextLiteral) String() string {
	if t == nil {
		return ""
	}
	return html.UnescapeString(t.Text)
}

var TrueText = BooleanText{
	Text: "true",
}

var FalseText = BooleanText{
	Text: "false",
}

type BooleanText struct {
	Text string `xml:",chardata"`
}

type IntegerText struct {
	Text string `xml:",chardata"`
}

func (b *BooleanText) ToBool() bool {
	return b != nil && strings.ToLower(b.Text) == "true"
}

func (b *BooleanText) IsTrue() bool {
	return b != nil && strings.ToLower(b.Text) == "true"
}

func (b *BooleanText) IsFalse() bool {
	return b != nil && strings.ToLower(b.Text) == "false"
}

func (b *BooleanText) String() string {
	return strconv.FormatBool(b.ToBool())
}

func (b BooleanText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(b.String(), start)
}
