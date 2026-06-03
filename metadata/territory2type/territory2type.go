package territory2type

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Territory2Type"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Territory2Type struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Territory2Type"`
	Xmlns   string   `xml:"xmlns,attr"`
	Name    struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Priority struct {
		Text string `xml:",chardata"`
	} `xml:"priority"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *Territory2Type) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Territory2Type) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Territory2Type, error) {
	p := &Territory2Type{}
	return p, metadata.ParseMetadataXml(p, path)
}
