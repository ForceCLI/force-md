package remoteSiteSetting

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "RemoteSiteSetting"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type RemoteSiteSetting struct {
	internal.MetadataInfo
	XMLName                 xml.Name `xml:"RemoteSiteSetting"`
	Xmlns                   string   `xml:"xmlns,attr"`
	DisableProtocolSecurity struct {
		Text string `xml:",chardata"`
	} `xml:"disableProtocolSecurity"`
	IsActive struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	URL struct {
		Text string `xml:",chardata"`
	} `xml:"url"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *RemoteSiteSetting) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *RemoteSiteSetting) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*RemoteSiteSetting, error) {
	p := &RemoteSiteSetting{}
	return p, internal.ParseMetadataXml(p, path)
}
