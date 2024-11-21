package notificationTypeConfig

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "NotificationTypeConfig"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type NotificationTypeConfig struct {
	internal.MetadataInfo
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

func (c *NotificationTypeConfig) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *NotificationTypeConfig) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*NotificationTypeConfig, error) {
	p := &NotificationTypeConfig{}
	return p, internal.ParseMetadataXml(p, path)
}
