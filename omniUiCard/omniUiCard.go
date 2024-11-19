package omniUiCard

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "OmniUiCard"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type OmniUiCard struct {
	internal.MetadataInfo
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
	VersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"versionNumber"`
	StylingConfiguration struct {
		Text string `xml:",chardata"`
	} `xml:"stylingConfiguration"`
}

func (c *OmniUiCard) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *OmniUiCard) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*OmniUiCard, error) {
	p := &OmniUiCard{}
	return p, internal.ParseMetadataXml(p, path)
}
