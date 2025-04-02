package platformEventChannels

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "PlatformEventChannel"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type PlatformEventChannel struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"PlatformEventChannel"`
	Xmlns       string   `xml:"xmlns,attr"`
	ChannelType string   `xml:"channelType"`
	Label       string   `xml:"label"`
}

func (c *PlatformEventChannel) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *PlatformEventChannel) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*PlatformEventChannel, error) {
	p := &PlatformEventChannel{}
	return p, metadata.ParseMetadataXml(p, path)
}

func NewPlatformEventChannel(label string) PlatformEventChannel {
	p := PlatformEventChannel{
		Label:       label,
		ChannelType: "event",
		Xmlns:       "http://soap.sforce.com/2006/04/metadata",
	}
	return p
}
