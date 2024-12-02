package autoResponseRules

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AutoResponseRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type AutoResponseRules struct {
	metadata.MetadataInfo
	XMLName          xml.Name `xml:"AutoResponseRules"`
	Xmlns            string   `xml:"xmlns,attr"`
	AutoResponseRule []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Active struct {
			Text string `xml:",chardata"`
		} `xml:"active"`
	} `xml:"autoResponseRule"`
}

func (c *AutoResponseRules) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AutoResponseRules) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AutoResponseRules, error) {
	p := &AutoResponseRules{}
	return p, metadata.ParseMetadataXml(p, path)
}
