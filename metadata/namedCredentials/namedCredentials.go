package namedCredentials

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "NamedCredential"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type NamedCredential struct {
	metadata.MetadataInfo
	XMLName                xml.Name `xml:"NamedCredential"`
	Xmlns                  string   `xml:"xmlns,attr"`
	AllowMergeFieldsInBody struct {
		Text string `xml:",chardata"`
	} `xml:"allowMergeFieldsInBody"`
	AllowMergeFieldsInHeader struct {
		Text string `xml:",chardata"`
	} `xml:"allowMergeFieldsInHeader"`
	CalloutStatus struct {
		Text string `xml:",chardata"`
	} `xml:"calloutStatus"`
	GenerateAuthorizationHeader struct {
		Text string `xml:",chardata"`
	} `xml:"generateAuthorizationHeader"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	NamedCredentialParameters []struct {
		ParameterName struct {
			Text string `xml:",chardata"`
		} `xml:"parameterName"`
		ParameterType struct {
			Text string `xml:",chardata"`
		} `xml:"parameterType"`
		ParameterValue struct {
			Text string `xml:",chardata"`
		} `xml:"parameterValue"`
		ExternalCredential struct {
			Text string `xml:",chardata"`
		} `xml:"externalCredential"`
	} `xml:"namedCredentialParameters"`
	NamedCredentialType struct {
		Text string `xml:",chardata"`
	} `xml:"namedCredentialType"`
}

func (c *NamedCredential) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *NamedCredential) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*NamedCredential, error) {
	p := &NamedCredential{}
	return p, metadata.ParseMetadataXml(p, path)
}
