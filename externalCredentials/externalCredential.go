package externalCredential

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "ExternalCredential"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ExternalCredential struct {
	internal.MetadataInfo
	XMLName                xml.Name `xml:"ExternalCredential"`
	Xmlns                  string   `xml:"xmlns,attr"`
	AuthenticationProtocol struct {
		Text string `xml:",chardata"`
	} `xml:"authenticationProtocol"`
	ExternalCredentialParameters []struct {
		ParameterName struct {
			Text string `xml:",chardata"`
		} `xml:"parameterName"`
		ParameterType struct {
			Text string `xml:",chardata"`
		} `xml:"parameterType"`
		SequenceNumber struct {
			Text string `xml:",chardata"`
		} `xml:"sequenceNumber"`
		ParameterValue struct {
			Text string `xml:",chardata"`
		} `xml:"parameterValue"`
	} `xml:"externalCredentialParameters"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

func (c *ExternalCredential) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ExternalCredential) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*ExternalCredential, error) {
	p := &ExternalCredential{}
	return p, internal.ParseMetadataXml(p, path)
}
