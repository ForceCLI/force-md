package mktDataSources

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

// NAME is the Data Cloud DataSource metadata type. Its files share the
// .dataSource suffix with ExternalDataSource but live in mktDataSources/
// and have a <DataSource> root element.
const NAME = "DataSource"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DataSource struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"DataSource"`
	Xmlns       string   `xml:"xmlns,attr"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Prefix struct {
		Text string `xml:",chardata"`
	} `xml:"prefix"`
}

func (c *DataSource) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DataSource) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DataSource, error) {
	p := &DataSource{}
	return p, metadata.ParseMetadataXml(p, path)
}
