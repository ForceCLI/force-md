package group

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "Group"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type Group struct {
	internal.MetadataInfo
	XMLName           xml.Name `xml:"Group"`
	Xmlns             string   `xml:"xmlns,attr"`
	DoesIncludeBosses struct {
		Text string `xml:",chardata"`
	} `xml:"doesIncludeBosses"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

func (c *Group) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Group) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Group, error) {
	p := &Group{}
	return p, internal.ParseMetadataXml(p, path)
}
