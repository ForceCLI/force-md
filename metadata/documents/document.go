package document

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Document"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Document struct {
	metadata.MetadataInfo
	XMLName         xml.Name `xml:"Document"`
	Xmlns           string   `xml:"xmlns,attr"`
	InternalUseOnly struct {
		Text string `xml:",chardata"`
	} `xml:"metadata"`
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

func (c *Document) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Document) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Document, error) {
	p := &Document{}
	return p, metadata.ParseMetadataXml(p, path)
}
