package platformEventSubscriberConfigs

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "PlatformEventSubscriberConfig"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type PlatformEventSubscriberConfig struct {
	metadata.MetadataInfo
	XMLName   xml.Name `xml:"PlatformEventSubscriberConfig"`
	Xmlns     string   `xml:"xmlns,attr"`
	BatchSize struct {
		Text string `xml:",chardata"`
	} `xml:"batchSize"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	PlatformEventConsumer struct {
		Text string `xml:",chardata"`
	} `xml:"platformEventConsumer"`
	User *TextLiteral `xml:"user"`
}

func (c *PlatformEventSubscriberConfig) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *PlatformEventSubscriberConfig) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*PlatformEventSubscriberConfig, error) {
	p := &PlatformEventSubscriberConfig{}
	return p, metadata.ParseMetadataXml(p, path)
}
