package objectTranslations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const FIELD_TRANSLATIONS_NAME = "CustomFieldTranslation"

func init() {
	internal.TypeRegistry.Register(FIELD_TRANSLATIONS_NAME, func(path string) (internal.RegisterableMetadata, error) { return OpenField(path) })
}

type CustomFieldTranslation struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"CustomFieldTranslation"`
	Xmlns      string   `xml:"xmlns,attr"`
	CaseValues struct {
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
		Text string `xml:",chardata"`
	} `xml:"label"`
	PicklistValues []struct {
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		Translation struct {
		} `xml:"translation"`
	} `xml:"picklistValues"`
}

func (c *CustomFieldTranslation) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomFieldTranslation) Type() internal.MetadataType {
	return FIELD_TRANSLATIONS_NAME
}

func OpenField(path string) (*CustomFieldTranslation, error) {
	p := &CustomFieldTranslation{}
	return p, internal.ParseMetadataXml(p, path)
}
