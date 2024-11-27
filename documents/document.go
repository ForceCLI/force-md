package document

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "Document"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type Document struct {
	internal.MetadataInfo
	XMLName         xml.Name `xml:"Document"`
	Xmlns           string   `xml:"xmlns,attr"`
	InternalUseOnly struct {
		Text string `xml:",chardata"`
	} `xml:"internalUseOnly"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Public struct {
		Text string `xml:",chardata"`
	} `xml:"public"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Keywords struct {
		Text string `xml:",chardata"`
	} `xml:"keywords"`
}

func (c *Document) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Document) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Document, error) {
	p := &Document{}
	return p, internal.ParseMetadataXml(p, path)
}