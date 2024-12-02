package fieldset

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "FieldSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FieldSetMetadata struct {
	metadata.MetadataInfo
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

func (c *FieldSetMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *FieldSetMetadata) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *FieldSetMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*FieldSetMetadata, error) {
	p := &FieldSetMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}
