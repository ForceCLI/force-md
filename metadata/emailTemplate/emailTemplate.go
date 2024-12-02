package emailTemplate

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "EmailTemplate"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type EmailTemplate struct {
	metadata.MetadataInfo
	XMLName   xml.Name `xml:"EmailTemplate"`
	Xmlns     string   `xml:"xmlns,attr"`
	Available struct {
		Text string `xml:",chardata"`
	} `xml:"available"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	EncodingKey struct {
		Text string `xml:",chardata"`
	} `xml:"encodingKey"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Style struct {
		Text string `xml:",chardata"`
	} `xml:"style"`
	Subject struct {
		Text string `xml:",chardata"`
	} `xml:"subject"`
	TemplateType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	TextOnly struct {
		Text string `xml:",chardata"`
	} `xml:"textOnly"`
	UiType struct {
		Text string `xml:",chardata"`
	} `xml:"uiType"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
}

func (c *EmailTemplate) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EmailTemplate) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*EmailTemplate, error) {
	p := &EmailTemplate{}
	return p, metadata.ParseMetadataXml(p, path)
}
