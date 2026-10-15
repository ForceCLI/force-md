package callCenters

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CallCenter"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CallCenter struct {
	metadata.MetadataInfo
	XMLName           xml.Name            `xml:"CallCenter"`
	Xmlns             string              `xml:"xmlns,attr"`
	Xsi               string              `xml:"xmlns:xsi,attr,omitempty"`
	AdapterUrl        *TextLiteral        `xml:"adapterUrl"`
	CustomSettings    *TextLiteral        `xml:"customSettings"`
	DisplayName       *TextLiteral        `xml:"displayName"`
	DisplayNameLabel  *TextLiteral        `xml:"displayNameLabel"`
	InternalNameLabel *TextLiteral        `xml:"internalNameLabel"`
	Sections          []CallCenterSection `xml:"sections"`
	Version           *TextLiteral        `xml:"version"`
}

type CallCenterSection struct {
	Items []CallCenterItem `xml:"items"`
	Label *TextLiteral     `xml:"label"`
	Name  *TextLiteral     `xml:"name"`
}

type CallCenterItem struct {
	Label *TextLiteral        `xml:"label"`
	Name  *TextLiteral        `xml:"name"`
	Value CallCenterItemValue `xml:"value"`
}

// CallCenterItemValue is an item's setting; an item with no value is written
// as <value xsi:nil="true"/>.
type CallCenterItemValue struct {
	Text string `xml:",innerxml"`
	Nil  string `xml:"xsi:nil,attr,omitempty"`
}

func (c *CallCenter) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CallCenter) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CallCenter, error) {
	c := &CallCenter{}
	return c, metadata.ParseMetadataXml(c, path)
}
