package fieldset

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects/split"
)

const NAME = "FieldSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type FieldSetMetadata struct {
	internal.MetadataInfo
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

func (c *FieldSetMetadata) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *FieldSetMetadata) NameFromPath(path string) string {
	return split.NameFromPath(path)
}

func (c *FieldSetMetadata) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*FieldSetMetadata, error) {
	p := &FieldSetMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
