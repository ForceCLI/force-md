package omniUiCard

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "OmniUiCard"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type OmniUiCard struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"OmniUiCard"`
	Xmlns      string   `xml:"xmlns,attr"`
	AuthorName struct {
		Text string `xml:",chardata"`
	} `xml:"authorName"`
	DataSourceConfig struct {
		Text string `xml:",chardata"`
	} `xml:"dataSourceConfig"`
	IsActive struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	OmniUiCardType struct {
		Text string `xml:",chardata"`
	} `xml:"omniUiCardType"`
	PropertySetConfig struct {
		Text string `xml:",chardata"`
	} `xml:"propertySetConfig"`
	SampleDataSourceResponse struct {
		Text string `xml:",chardata"`
	} `xml:"sampleDataSourceResponse"`
	StylingConfiguration *struct {
		Text string `xml:",chardata"`
	} `xml:"stylingConfiguration"`
	VersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"versionNumber"`
}

func (c *OmniUiCard) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *OmniUiCard) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*OmniUiCard, error) {
	p := &OmniUiCard{}
	return p, metadata.ParseMetadataXml(p, path)
}
