package notificationtype

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomNotificationType"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CustomNotificationType struct {
	metadata.MetadataInfo
	XMLName             xml.Name `xml:"CustomNotificationType"`
	Xmlns               string   `xml:"xmlns,attr"`
	CustomNotifTypeName struct {
		Text string `xml:",chardata"`
	} `xml:"customNotifTypeName"`
	Desktop struct {
		Text string `xml:",chardata"`
	} `xml:"desktop"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Mobile struct {
		Text string `xml:",chardata"`
	} `xml:"mobile"`
	Slack struct {
		Text string `xml:",chardata"`
	} `xml:"slack"`
}

func (c *CustomNotificationType) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomNotificationType) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomNotificationType, error) {
	p := &CustomNotificationType{}
	return p, metadata.ParseMetadataXml(p, path)
}
