package objecttranslation

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objecttranslation/field"
)

type FieldList []field.Field

type CustomObjectTranslation struct {
	internal.MetadataInfo
	XMLName     xml.Name `xml:"CustomObjectTranslation"`
	Xmlns       string   `xml:"xmlns,attr"`
	RecordTypes []struct {
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
	ValidationRules []struct {
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

func Open(path string) (*CustomObjectTranslation, error) {
	p := &CustomObjectTranslation{}
	return p, internal.ParseMetadataXml(p, path)
}
