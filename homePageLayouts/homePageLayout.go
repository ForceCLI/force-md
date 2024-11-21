package homePageLayout

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "HomePageLayout"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type HomePageLayout struct {
	internal.MetadataInfo
	XMLName          xml.Name `xml:"HomePageLayout"`
	Xmlns            string   `xml:"xmlns,attr"`
	NarrowComponents []struct {
		Text string `xml:",chardata"`
	} `xml:"narrowComponents"`
	WideComponents []struct {
		Text string `xml:",chardata"`
	} `xml:"wideComponents"`
}

func (c *HomePageLayout) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *HomePageLayout) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*HomePageLayout, error) {
	p := &HomePageLayout{}
	return p, internal.ParseMetadataXml(p, path)
}
