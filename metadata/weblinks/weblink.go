package weblink

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomPageWebLink"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CustomPageWebLink struct {
	metadata.MetadataInfo
	XMLName      xml.Name `xml:"CustomPageWebLink"`
	Xmlns        string   `xml:"xmlns,attr"`
	Availability struct {
		Text string `xml:",chardata"`
	} `xml:"availability"`
	DisplayType struct {
		Text string `xml:",chardata"`
	} `xml:"displayType"`
	EncodingKey struct {
		Text string `xml:",chardata"`
	} `xml:"encodingKey"`
	HasMenubar struct {
		Text string `xml:",chardata"`
	} `xml:"hasMenubar"`
	HasScrollbars struct {
		Text string `xml:",chardata"`
	} `xml:"hasScrollbars"`
	HasToolbar struct {
		Text string `xml:",chardata"`
	} `xml:"hasToolbar"`
	Height struct {
		Text string `xml:",chardata"`
	} `xml:"height"`
	IsResizable struct {
		Text string `xml:",chardata"`
	} `xml:"isResizable"`
	LinkType struct {
		Text string `xml:",chardata"`
	} `xml:"linkType"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	OpenType struct {
		Text string `xml:",chardata"`
	} `xml:"openType"`
	Position struct {
		Text string `xml:",chardata"`
	} `xml:"position"`
	Protected struct {
		Text string `xml:",chardata"`
	} `xml:"protected"`
	ShowsLocation struct {
		Text string `xml:",chardata"`
	} `xml:"showsLocation"`
	ShowsStatus struct {
		Text string `xml:",chardata"`
	} `xml:"showsStatus"`
	URL struct {
		Text string `xml:",chardata"`
	} `xml:"url"`
}

func (c *CustomPageWebLink) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomPageWebLink) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomPageWebLink, error) {
	p := &CustomPageWebLink{}
	return p, metadata.ParseMetadataXml(p, path)
}
