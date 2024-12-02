package homePageLayout

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "HomePageLayout"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type HomePageLayout struct {
	metadata.MetadataInfo
	XMLName          xml.Name `xml:"HomePageLayout"`
	Xmlns            string   `xml:"xmlns,attr"`
	NarrowComponents []struct {
		Text string `xml:",chardata"`
	} `xml:"narrowComponents"`
	WideComponents []struct {
		Text string `xml:",chardata"`
	} `xml:"wideComponents"`
}

func (c *HomePageLayout) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *HomePageLayout) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*HomePageLayout, error) {
	p := &HomePageLayout{}
	return p, metadata.ParseMetadataXml(p, path)
}
