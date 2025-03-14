package dataweave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DataWeaveResource"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DataWeaveResource struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"DataWeaveResource"`
	Xmlns   string   `xml:"xmlns,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Content struct {
		Nil string `xml:"nil,attr"`
	} `xml:"content"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	IsGlobal struct {
		Text string `xml:",chardata"`
	} `xml:"isGlobal"`
	IsProtected struct {
		Text string `xml:",chardata"`
	} `xml:"isProtected"`
}

func (c *DataWeaveResource) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DataWeaveResource) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DataWeaveResource, error) {
	p := &DataWeaveResource{}
	return p, metadata.ParseMetadataXml(p, path)
}
