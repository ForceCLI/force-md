package aura

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "AuraDefinitionBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type AuraDefinitionBundle struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"AuraDefinitionBundle"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
}

func (c *AuraDefinitionBundle) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AuraDefinitionBundle) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*AuraDefinitionBundle, error) {
	p := &AuraDefinitionBundle{}
	return p, internal.ParseMetadataXml(p, path)
}
