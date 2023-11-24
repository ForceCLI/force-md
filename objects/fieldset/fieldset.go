package fieldset

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type FieldSetMetadata struct {
	XMLName xml.Name `xml:"FieldSet"`
	Xmlns   string   `xml:"xmlns,attr"`
	FieldSet
}

type FieldSet struct {
	FullName        string `xml:"fullName"`
	AvailableFields []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		IsFieldManaged struct {
			Text string `xml:",chardata"`
		} `xml:"isFieldManaged"`
		IsRequired struct {
			Text string `xml:",chardata"`
		} `xml:"isRequired"`
	} `xml:"availableFields"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	DisplayedFields []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		IsFieldManaged struct {
			Text string `xml:",chardata"`
		} `xml:"isFieldManaged"`
		IsRequired struct {
			Text string `xml:",chardata"`
		} `xml:"isRequired"`
	} `xml:"displayedFields"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

func (p *FieldSetMetadata) MetaCheck() {}

func Open(path string) (*FieldSetMetadata, error) {
	p := &FieldSetMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
