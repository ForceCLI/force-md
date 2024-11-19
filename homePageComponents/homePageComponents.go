package homePageComponents

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "HomePageComponent"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type HomePageComponent struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"HomePageComponent"`
	Xmlns   string   `xml:"xmlns,attr"`
	Height  struct {
		Text string `xml:",chardata"`
	} `xml:"height"`
	Page struct {
		Text string `xml:",chardata"`
	} `xml:"page"`
	PageComponentType struct {
		Text string `xml:",chardata"`
	} `xml:"pageComponentType"`
	ShowLabel struct {
		Text string `xml:",chardata"`
	} `xml:"showLabel"`
	ShowScrollbars struct {
		Text string `xml:",chardata"`
	} `xml:"showScrollbars"`
	Width struct {
		Text string `xml:",chardata"`
	} `xml:"width"`
}

func (c *HomePageComponent) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *HomePageComponent) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*HomePageComponent, error) {
	p := &HomePageComponent{}
	return p, internal.ParseMetadataXml(p, path)
}
