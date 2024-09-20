package weblink

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type WebLinkMetadata struct {
	Metadata
	XMLName xml.Name `xml:"WebLink"`
	Xmlns   string   `xml:"xmlns,attr"`
	WebLink
}

type WebLink struct {
	FullName     string `xml:"fullName"`
	Availability struct {
		Text string `xml:",chardata"`
	} `xml:"availability"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	DisplayType struct {
		Text string `xml:",chardata"`
	} `xml:"displayType"`
	EncodingKey *struct {
		Text string `xml:",chardata"`
	} `xml:"encodingKey"`
	HasMenubar *struct {
		Text string `xml:",chardata"`
	} `xml:"hasMenubar"`
	HasScrollbars *struct {
		Text string `xml:",chardata"`
	} `xml:"hasScrollbars"`
	HasToolbar *struct {
		Text string `xml:",chardata"`
	} `xml:"hasToolbar"`
	Height *struct {
		Text string `xml:",chardata"`
	} `xml:"height"`
	IsResizable *struct {
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
	Page *struct {
		Text string `xml:",chardata"`
	} `xml:"page"`
	Position *struct {
		Text string `xml:",chardata"`
	} `xml:"position"`
	Protected struct {
		Text string `xml:",chardata"`
	} `xml:"protected"`
	RequireRowSelection *struct {
		Text string `xml:",chardata"`
	} `xml:"requireRowSelection"`
	ShowsLocation *struct {
		Text string `xml:",chardata"`
	} `xml:"showsLocation"`
	ShowsStatus *struct {
		Text string `xml:",chardata"`
	} `xml:"showsStatus"`
	URL *struct {
		Text string `xml:",innerxml"`
	} `xml:"url"`
	Width *struct {
		Text string `xml:",chardata"`
	} `xml:"width"`
}

func (c *WebLinkMetadata) SetMetadata(m Metadata) {
	c.Metadata = m
}

func Open(path string) (*WebLinkMetadata, error) {
	p := &WebLinkMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
