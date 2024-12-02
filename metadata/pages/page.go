package apexPage

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ApexPage"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexPage struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ApexPage"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	AvailableInTouch struct {
		Text string `xml:",chardata"`
	} `xml:"availableInTouch"`
	ConfirmationTokenRequired struct {
		Text string `xml:",chardata"`
	} `xml:"confirmationTokenRequired"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *ApexPage) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexPage) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ApexPage, error) {
	p := &ApexPage{}
	return p, metadata.ParseMetadataXml(p, path)
}
