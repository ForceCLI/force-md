package flowDefinition

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "FlowDefinition"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FlowDefinition struct {
	metadata.MetadataInfo
	XMLName             xml.Name `xml:"FlowDefinition"`
	Xmlns               string   `xml:"xmlns,attr"`
	ActiveVersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"activeVersionNumber"`
}

func (c *FlowDefinition) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *FlowDefinition) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*FlowDefinition, error) {
	p := &FlowDefinition{}
	return p, metadata.ParseMetadataXml(p, path)
}
