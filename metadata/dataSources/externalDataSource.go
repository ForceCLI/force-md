package datasource

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ExternalDataSource"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ExternalDataSource struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ExternalDataSource"`
	Xmlns      string   `xml:"xmlns,attr"`
	IsWritable struct {
		Text string `xml:",chardata"`
	} `xml:"isWritable"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	PrincipalType struct {
		Text string `xml:",chardata"`
	} `xml:"principalType"`
	Protocol struct {
		Text string `xml:",chardata"`
	} `xml:"protocol"`
	DataSourceType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
}

func (c *ExternalDataSource) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ExternalDataSource) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ExternalDataSource, error) {
	p := &ExternalDataSource{}
	return p, metadata.ParseMetadataXml(p, path)
}
