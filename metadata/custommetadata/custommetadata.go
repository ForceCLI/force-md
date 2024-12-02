package custommetadata

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomMetadata"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type TypedValue struct {
	Text string `xml:",innerxml"`
	Type string `xml:"xsi:type,attr,omitempty"`
	Nil  string `xml:"xsi:nil,attr,omitempty"`
}

type Value struct {
	Field string     `xml:"field"`
	Value TypedValue `xml:"value"`
}

type CustomMetadata struct {
	metadata.MetadataInfo
	XMLName   xml.Name    `xml:"CustomMetadata"`
	Xmlns     string      `xml:"xmlns,attr"`
	Xsi       string      `xml:"xmlns:xsi,attr,omitempty"`
	Xsd       string      `xml:"xmlns:xsd,attr,omitempty"`
	Label     string      `xml:"label"`
	Protected BooleanText `xml:"protected"`
	Values    []Value     `xml:"values"`
}

func (c *CustomMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomMetadata, error) {
	p := &CustomMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}
