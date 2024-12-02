package community

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Community"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Community struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Community"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
}

func (c *Community) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Community) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Community, error) {
	p := &Community{}
	return p, metadata.ParseMetadataXml(p, path)
}
