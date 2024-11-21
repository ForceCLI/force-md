package notificationtype

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "CustomNotificationType"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type CustomNotificationType struct {
	internal.MetadataInfo
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

func (c *CustomNotificationType) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomNotificationType) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*CustomNotificationType, error) {
	p := &CustomNotificationType{}
	return p, internal.ParseMetadataXml(p, path)
}
