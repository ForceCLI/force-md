package emailTemplate

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "EmailTemplate"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type EmailTemplate struct {
	internal.MetadataInfo
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

func (c *EmailTemplate) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EmailTemplate) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*EmailTemplate, error) {
	p := &EmailTemplate{}
	return p, internal.ParseMetadataXml(p, path)
}
