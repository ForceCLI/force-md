package objectTranslations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const OBJECT_TRANSLATIONS_NAME = "CustomObjectTranslation"

func init() {
	internal.TypeRegistry.Register(OBJECT_TRANSLATIONS_NAME, func(path string) (internal.RegisterableMetadata, error) { return OpenObject(path) })
}

type CustomObjectTranslation struct {
	internal.MetadataInfo
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
	ValidationRules struct {
		ErrorMessage struct {
		} `xml:"errorMessage"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"validationRules"`
}

func (c *CustomObjectTranslation) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomObjectTranslation) Type() internal.MetadataType {
	return OBJECT_TRANSLATIONS_NAME
}

func OpenObject(path string) (*CustomObjectTranslation, error) {
	p := &CustomObjectTranslation{}
	return p, internal.ParseMetadataXml(p, path)
}
