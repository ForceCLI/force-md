package notificationTypeConfig

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "NotificationTypeConfig"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type NotificationTypeConfig struct {
	metadata.MetadataInfo
	XMLName                  xml.Name `xml:"NotificationTypeConfig"`
	Xmlns                    string   `xml:"xmlns,attr"`
	NotificationTypeSettings []struct {
		AppSettings []struct {
			ConnectedAppName struct {
				Text string `xml:",chardata"`
			} `xml:"connectedAppName"`
			Enabled struct {
				Text string `xml:",chardata"`
			} `xml:"enabled"`
		} `xml:"appSettings"`
		NotificationChannels struct {
			DesktopEnabled struct {
				Text string `xml:",chardata"`
			} `xml:"desktopEnabled"`
			MobileEnabled struct {
				Text string `xml:",chardata"`
			} `xml:"mobileEnabled"`
		} `xml:"notificationChannels"`
		NotificationType struct {
			Text string `xml:",chardata"`
		} `xml:"notificationType"`
	} `xml:"notificationTypeSettings"`
}

func (c *NotificationTypeConfig) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *NotificationTypeConfig) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*NotificationTypeConfig, error) {
	p := &NotificationTypeConfig{}
	return p, metadata.ParseMetadataXml(p, path)
}
