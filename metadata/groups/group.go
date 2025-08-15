package group

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Group"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Group struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"Group"`
	Xmlns       string   `xml:"xmlns,attr"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	DoesIncludeBosses struct {
		Text string `xml:",chardata"`
	} `xml:"doesIncludeBosses"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

func (c *Group) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Group) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Group, error) {
	p := &Group{}
	return p, metadata.ParseMetadataXml(p, path)
}
