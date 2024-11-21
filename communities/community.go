package community

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "Community"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type Community struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"Community"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
}

func (c *Community) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Community) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Community, error) {
	p := &Community{}
	return p, internal.ParseMetadataXml(p, path)
}
