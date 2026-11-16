package general

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	nbioxml "github.com/nbio/xml"
)

type Metadata = metadata.MetadataPointer

// TextLiteral is the text of an element.  It holds the element's character
// data and the inner XML that encodes it.  An element read from metadata is
// written back with the inner XML it was read with, so a CDATA section stays a
// CDATA section; text set with NewTextLiteral or Set is written escaped with
// EscapeText.  The inner XML is always well-formed: it is either read by the
// XML decoder or produced by EscapeText.
//
// UnmarshalXML takes github.com/nbio/xml types because metadata files are
// parsed with that package, while MarshalXML takes encoding/xml types because
// they are written with encoding/xml.  NillableText and TypedText follow the
// same rule.
type TextLiteral struct {
	innerXML string
	value    string
}

// NewTextLiteral returns a TextLiteral whose character data is value.
func NewTextLiteral(value string) *TextLiteral {
	t := &TextLiteral{}
	t.Set(value)
	return t
}

// EscapeText returns s escaped for use as the inner XML of an element, with
// the entity references Salesforce writes in metadata.
func EscapeText(s string) string {
	return internal.FormulaEscaper.Replace(s)
}

// String returns the character data of the element.  A nil TextLiteral has no
// character data.
func (t *TextLiteral) String() string {
	if t == nil {
		return ""
	}
	return t.value
}

// Set replaces the character data of the element with value.
func (t *TextLiteral) Set(value string) {
	t.innerXML = EscapeText(value)
	t.value = value
}

// Reformat replaces the character data of the element with the result of
// applying format to it.  The element is unchanged if format returns an error.
func (t *TextLiteral) Reformat(format func(string) (string, error)) error {
	formatted, err := format(t.value)
	if err != nil {
		return err
	}
	t.Set(formatted)
	return nil
}

// textElement reads an element's inner XML and its character data in one
// pass of the XML decoder.
type textElement struct {
	InnerXML string `xml:",innerxml"`
	Value    string `xml:",chardata"`
}

type innerXMLElement struct {
	InnerXML string `xml:",innerxml"`
}

func (t *TextLiteral) UnmarshalXML(d *nbioxml.Decoder, start nbioxml.StartElement) error {
	var e textElement
	if err := d.DecodeElement(&e, &start); err != nil {
		return err
	}
	t.innerXML = e.InnerXML
	t.value = e.Value
	return nil
}

func (t TextLiteral) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(innerXMLElement{InnerXML: t.innerXML}, start)
}

// NillableText is an element written as <name xsi:nil="true"/> when it has
// no value.  The root element of the metadata declares the xsi namespace.
type NillableText struct {
	TextLiteral
	Nil string
}

func (t *NillableText) UnmarshalXML(d *nbioxml.Decoder, start nbioxml.StartElement) error {
	var e struct {
		textElement
		Nil string `xml:"xsi:nil,attr,omitempty"`
	}
	if err := d.DecodeElement(&e, &start); err != nil {
		return err
	}
	t.innerXML = e.InnerXML
	t.value = e.Value
	t.Nil = e.Nil
	return nil
}

func (t NillableText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		innerXMLElement
		Nil string `xml:"xsi:nil,attr,omitempty"`
	}{innerXMLElement{t.innerXML}, t.Nil}, start)
}

// TypedText is a value written with its XML schema type,
// <value xsi:type="xsd:string">, or as <value xsi:nil="true"/> when it has no
// value.
type TypedText struct {
	TextLiteral
	Type string
	Nil  string
}

func (t *TypedText) UnmarshalXML(d *nbioxml.Decoder, start nbioxml.StartElement) error {
	var e struct {
		textElement
		Type string `xml:"xsi:type,attr,omitempty"`
		Nil  string `xml:"xsi:nil,attr,omitempty"`
	}
	if err := d.DecodeElement(&e, &start); err != nil {
		return err
	}
	t.innerXML = e.InnerXML
	t.value = e.Value
	t.Type = e.Type
	t.Nil = e.Nil
	return nil
}

func (t TypedText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		innerXMLElement
		Type string `xml:"xsi:type,attr,omitempty"`
		Nil  string `xml:"xsi:nil,attr,omitempty"`
	}{innerXMLElement{t.innerXML}, t.Type, t.Nil}, start)
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
