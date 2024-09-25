package custommetadata

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type TypedValue struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr,omitempty"`
	Nil  string `xml:"xsi:nil,attr,omitempty"`
}

type Value struct {
	Field string     `xml:"field"`
	Value TypedValue `xml:"value"`
}

type CustomMetadata struct {
	internal.MetadataInfo
	XMLName   xml.Name    `xml:"CustomMetadata"`
	Xmlns     string      `xml:"xmlns,attr"`
	Xsi       string      `xml:"xmlns:xsi,attr"`
	Xsd       string      `xml:"xmlns:xsd,attr"`
	Label     string      `xml:"label"`
	Protected BooleanText `xml:"protected"`
	Values    []Value     `xml:"values"`
}

func (c *CustomMetadata) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*CustomMetadata, error) {
	p := &CustomMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
