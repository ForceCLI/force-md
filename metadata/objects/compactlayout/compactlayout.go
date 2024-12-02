package compactlayout

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "CompactLayout"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CompactLayoutMetadata struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"CompactLayout"`
	Xmlns   string   `xml:"xmlns,attr"`
	CompactLayout
}

type CompactLayout struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Fields []struct {
		Text string `xml:",chardata"`
	} `xml:"fields"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

func (c *CompactLayoutMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CompactLayoutMetadata) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *CompactLayoutMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CompactLayoutMetadata, error) {
	p := &CompactLayoutMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}
