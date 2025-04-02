package platformEventChannelMembers

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "PlatformEventChannelMember"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FilterExpression struct {
	Text string `xml:",cdata"`
}

type PlatformEventChannelMember struct {
	metadata.MetadataInfo
	XMLName          xml.Name         `xml:"PlatformEventChannelMember"`
	Xmlns            string           `xml:"xmlns,attr"`
	EventChannel     string           `xml:"eventChannel"`
	FilterExpression FilterExpression `xml:"filterExpression"`
	SelectedEntity   string           `xml:"selectedEntity"`
}

func (c *PlatformEventChannelMember) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *PlatformEventChannelMember) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*PlatformEventChannelMember, error) {
	p := &PlatformEventChannelMember{}
	return p, metadata.ParseMetadataXml(p, path)
}

func NewPlatformEventChannelMember(channel, entity, filter string) PlatformEventChannelMember {
	p := PlatformEventChannelMember{
		EventChannel:     channel,
		SelectedEntity:   entity,
		FilterExpression: FilterExpression{Text: filter},
		Xmlns:            "http://soap.sforce.com/2006/04/metadata",
	}
	return p
}
