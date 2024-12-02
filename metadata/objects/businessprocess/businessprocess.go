package businessprocess

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "BusinessProcess"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type BusinessProcessMetadata struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"BusinessProcess"`
	Xmlns   string   `xml:"xmlns,attr"`
	BusinessProcess
}

type BusinessProcess struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	IsActive struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	Values []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Default struct {
			Text string `xml:",chardata"`
		} `xml:"default"`
	} `xml:"values"`
}

func (c *BusinessProcessMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *BusinessProcessMetadata) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *BusinessProcessMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*BusinessProcessMetadata, error) {
	p := &BusinessProcessMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}
