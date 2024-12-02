package field

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const FIELD_TRANSLATIONS_NAME = "CustomFieldTranslation"

func init() {
	internal.TypeRegistry.Register(FIELD_TRANSLATIONS_NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Field struct {
	CaseValues []struct {
		Plural struct {
			Text string `xml:",chardata"`
		} `xml:"plural"`
		Value struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"caseValues"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	StartsWith struct {
		Text string `xml:",chardata"`
	} `xml:"startsWith"`
	Label struct {
	} `xml:"label"`
	PicklistValues []struct {
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		Translation struct {
		} `xml:"translation"`
	} `xml:"picklistValues"`
}

type CustomFieldTranslation struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"CustomFieldTranslation"`
	Xmlns   string   `xml:"xmlns,attr"`
	Field
}

func (c *CustomFieldTranslation) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomFieldTranslation) Type() metadata.MetadataType {
	return FIELD_TRANSLATIONS_NAME
}

func Open(path string) (*CustomFieldTranslation, error) {
	p := &CustomFieldTranslation{}
	return p, metadata.ParseMetadataXml(p, path)
}
