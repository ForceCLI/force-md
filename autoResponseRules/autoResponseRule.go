package autoResponseRules

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "AutoResponseRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type AutoResponseRules struct {
	internal.MetadataInfo
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

func (c *AutoResponseRules) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AutoResponseRules) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*AutoResponseRules, error) {
	p := &AutoResponseRules{}
	return p, internal.ParseMetadataXml(p, path)
}
