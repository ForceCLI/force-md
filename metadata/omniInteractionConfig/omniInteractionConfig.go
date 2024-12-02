package omniInteractionConfig

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "OmniInteractionConfig"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type OmniInteractionConfig struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"OmniInteractionConfig"`
	Xmlns       string   `xml:"xmlns,attr"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Value struct {
		Text string `xml:",chardata"`
	} `xml:"value"`
}

func (c *OmniInteractionConfig) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *OmniInteractionConfig) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*OmniInteractionConfig, error) {
	p := &OmniInteractionConfig{}
	return p, metadata.ParseMetadataXml(p, path)
}
