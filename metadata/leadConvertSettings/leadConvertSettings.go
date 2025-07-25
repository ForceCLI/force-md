package leadConvertSettings

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "LeadConvertSettings"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type LeadConvertSettings struct {
	metadata.MetadataInfo

	XMLName          xml.Name `xml:"LeadConvertSettings"`
	Xmlns            string   `xml:"xmlns,attr"`
	AllowOwnerChange struct {
		Text string `xml:",chardata"`
	} `xml:"allowOwnerChange"`
	ObjectMapping []struct {
		InputObject struct {
			Text string `xml:",chardata"`
		} `xml:"inputObject"`
		MappingFields struct {
			InputField struct {
				Text string `xml:",chardata"`
			} `xml:"inputField"`
			OutputField struct {
				Text string `xml:",chardata"`
			} `xml:"outputField"`
		} `xml:"mappingFields"`
		OutputObject struct {
			Text string `xml:",chardata"`
		} `xml:"outputObject"`
	} `xml:"objectMapping"`
	OpportunityCreationOptions struct {
		Text string `xml:",chardata"`
	} `xml:"opportunityCreationOptions"`
}

func (c *LeadConvertSettings) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*LeadConvertSettings, error) {
	p := &LeadConvertSettings{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *LeadConvertSettings) Type() metadata.MetadataType {
	return NAME
}
