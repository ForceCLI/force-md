package validationrule

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type ValidationRule struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"ValidationRule"`
	Xmlns   string   `xml:"xmlns,attr"`
	Rule
}

type Rule struct {
	FullName string `xml:"fullName"`
	Active   struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	Description *struct {
		Text string `xml:",innerxml"`
	} `xml:"description"`
	ErrorConditionFormula struct {
		Text string `xml:",innerxml"`
	} `xml:"errorConditionFormula"`
	ErrorDisplayField *struct {
		Text string `xml:",chardata"`
	} `xml:"errorDisplayField"`
	ErrorMessage struct {
		Text string `xml:",innerxml"`
	} `xml:"errorMessage"`
}

func (c *ValidationRule) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*ValidationRule, error) {
	p := &ValidationRule{}
	return p, internal.ParseMetadataXml(p, path)
}
