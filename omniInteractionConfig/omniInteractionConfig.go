package omniInteractionConfig

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "OmniInteractionConfig"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type OmniInteractionConfig struct {
	internal.MetadataInfo
	XMLName     xml.Name `xml:"OmniInteractionConfig"`
	Xmlns       string   `xml:"xmlns,attr"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Value struct {
		Text string `xml:",chardata"`
	} `xml:"value"`
}

func (c *OmniInteractionConfig) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *OmniInteractionConfig) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*OmniInteractionConfig, error) {
	p := &OmniInteractionConfig{}
	return p, internal.ParseMetadataXml(p, path)
}
