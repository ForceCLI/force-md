package weblink

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects/split"
)

const NAME = "WebLink"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type WebLinkMetadata struct {
	internal.MetadataInfo
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

func (c *WebLinkMetadata) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WebLinkMetadata) NameFromPath(path string) string {
	return split.NameFromPath(path)
}

func (c *WebLinkMetadata) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*WebLinkMetadata, error) {
	p := &WebLinkMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
