package staticresource

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "StaticResource"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type StaticResource struct {
	internal.MetadataInfo
	XMLName      xml.Name `xml:"StaticResource"`
	Xmlns        string   `xml:"xmlns,attr"`
	CacheControl struct {
		Text string `xml:",chardata"`
	} `xml:"cacheControl"`
	ContentType struct {
		Text string `xml:",chardata"`
	} `xml:"contentType"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *StaticResource) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *StaticResource) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*StaticResource, error) {
	p := &StaticResource{}
	return p, internal.ParseMetadataXml(p, path)
}
