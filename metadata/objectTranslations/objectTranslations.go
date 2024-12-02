package objectTranslations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objectTranslations/field"
)

const OBJECT_TRANSLATIONS_NAME = "CustomObjectTranslation"

func init() {
	internal.TypeRegistry.Register(OBJECT_TRANSLATIONS_NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FieldList []field.Field

type CustomObjectTranslation struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"CustomObjectTranslation"`
	Xmlns       string   `xml:"xmlns,attr"`
	RecordTypes struct {
		Description struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
		Label struct {
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"recordTypes"`
	Fields          FieldList `xml:"fields"`
	ValidationRules struct {
		ErrorMessage struct {
		} `xml:"errorMessage"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"validationRules"`
}

func (c *CustomObjectTranslation) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomObjectTranslation) Type() metadata.MetadataType {
	return OBJECT_TRANSLATIONS_NAME
}

func Open(path string) (*CustomObjectTranslation, error) {
	p := &CustomObjectTranslation{}
	return p, metadata.ParseMetadataXml(p, path)
}
