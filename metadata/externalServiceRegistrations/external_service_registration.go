package externalServiceRegistrations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ExternalServiceRegistration"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// ExternalServiceRegistration mirrors the metadata API ExternalServiceRegistration
// type. The schema body itself (an OpenAPI 2.0 or 3.0 document, JSON or YAML) is
// kept verbatim in the Schema field for the consumer to parse.
type ExternalServiceRegistration struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"ExternalServiceRegistration"`
	Xmlns   string   `xml:"xmlns,attr"`

	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	ExternalServiceProviderType *struct {
		Text string `xml:",chardata"`
	} `xml:"externalServiceProviderType"`
	Label *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	NamedCredential *struct {
		Text string `xml:",chardata"`
	} `xml:"namedCredential"`
	NamedCredentialReference *struct {
		Text string `xml:",chardata"`
	} `xml:"namedCredentialReference"`
	RegistrationUrl *struct {
		Text string `xml:",chardata"`
	} `xml:"registrationUrl"`
	RegistrationProviderType *struct {
		Text string `xml:",chardata"`
	} `xml:"registrationProviderType"`
	Schema *struct {
		Text string `xml:",chardata"`
	} `xml:"schema"`
	SchemaType *struct {
		Text string `xml:",chardata"`
	} `xml:"schemaType"`
	SchemaUploadFileExtension *struct {
		Text string `xml:",chardata"`
	} `xml:"schemaUploadFileExtension"`
	SchemaUploadFileName *struct {
		Text string `xml:",chardata"`
	} `xml:"schemaUploadFileName"`
	Status *struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	SystemVersion *struct {
		Text string `xml:",chardata"`
	} `xml:"systemVersion"`
	Operations []struct {
		Active *struct {
			Text string `xml:",chardata"`
		} `xml:"active"`
		Name *struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"operations"`
}

func (c *ExternalServiceRegistration) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ExternalServiceRegistration) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ExternalServiceRegistration, error) {
	p := &ExternalServiceRegistration{}
	return p, metadata.ParseMetadataXml(p, path)
}
