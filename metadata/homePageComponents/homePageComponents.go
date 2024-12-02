package homePageComponents

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "HomePageComponent"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type HomePageComponent struct {
	metadata.MetadataInfo
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

func (c *HomePageComponent) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *HomePageComponent) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*HomePageComponent, error) {
	p := &HomePageComponent{}
	return p, metadata.ParseMetadataXml(p, path)
}
