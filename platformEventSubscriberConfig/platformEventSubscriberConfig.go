package platformEventSubscriberConfig

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type PlatformEventSubscriberConfig struct {
	internal.MetadataInfo
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

func (c *PlatformEventSubscriberConfig) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*PlatformEventSubscriberConfig, error) {
	p := &PlatformEventSubscriberConfig{}
	return p, internal.ParseMetadataXml(p, path)
}
