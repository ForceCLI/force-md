package components

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ApexComponent"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexComponent struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ApexComponent"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *ApexComponent) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexComponent) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ApexComponent, error) {
	p := &ApexComponent{}
	return p, metadata.ParseMetadataXml(p, path)
}
