package flowDefinition

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "FlowDefinition"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type FlowDefinition struct {
	internal.MetadataInfo
	XMLName             xml.Name `xml:"FlowDefinition"`
	Xmlns               string   `xml:"xmlns,attr"`
	ActiveVersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"activeVersionNumber"`
}

func (c *FlowDefinition) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *FlowDefinition) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*FlowDefinition, error) {
	p := &FlowDefinition{}
	return p, internal.ParseMetadataXml(p, path)
}
