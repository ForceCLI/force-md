package remoteSiteSetting

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "RemoteSiteSetting"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RemoteSiteSetting struct {
	metadata.MetadataInfo
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

func (c *RemoteSiteSetting) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *RemoteSiteSetting) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*RemoteSiteSetting, error) {
	p := &RemoteSiteSetting{}
	return p, metadata.ParseMetadataXml(p, path)
}
