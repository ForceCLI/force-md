package tab

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "CustomTab"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type CustomTab struct {
	internal.MetadataInfo
	XMLName      xml.Name `xml:"CustomTab"`
	Xmlns        string   `xml:"xmlns,attr"`
	CustomObject struct {
		Text string `xml:",chardata"`
	} `xml:"customObject"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Motif struct {
		Text string `xml:",chardata"`
	} `xml:"motif"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Page struct {
		Text string `xml:",chardata"`
	} `xml:"page"`
	FlexiPage struct {
		Text string `xml:",chardata"`
	} `xml:"flexiPage"`
	Icon struct {
		Text string `xml:",chardata"`
	} `xml:"icon"`
	LwcComponent struct {
		Text string `xml:",chardata"`
	} `xml:"lwcComponent"`
	FrameHeight struct {
		Text string `xml:",chardata"`
	} `xml:"frameHeight"`
	HasSidebar struct {
		Text string `xml:",chardata"`
	} `xml:"hasSidebar"`
	URL struct {
		Text string `xml:",chardata"`
	} `xml:"url"`
	UrlEncodingKey struct {
		Text string `xml:",chardata"`
	} `xml:"urlEncodingKey"`
}

func (c *CustomTab) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomTab) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*CustomTab, error) {
	p := &CustomTab{}
	return p, internal.ParseMetadataXml(p, path)
}
