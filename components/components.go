package components

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "ApexComponent"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ApexComponent struct {
	internal.MetadataInfo
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

func (c *ApexComponent) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexComponent) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*ApexComponent, error) {
	p := &ApexComponent{}
	return p, internal.ParseMetadataXml(p, path)
}
