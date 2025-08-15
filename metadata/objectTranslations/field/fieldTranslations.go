package field

import (
	"encoding/xml"

	"path/filepath"
	"regexp"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const FIELD_TRANSLATIONS_NAME = "CustomFieldTranslation"

var fieldTransPathRegex = regexp.MustCompile(`.*/objectTranslations/([^/]+)/(?:fields/)?([^/]+)\.fieldTranslation-meta\.xml$`)

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
	Label *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	StartsWith *struct {
		Text string `xml:",chardata"`
	} `xml:"startsWith"`
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

func (c *CustomFieldTranslation) NameFromPath(path string) metadata.MetadataObjectName {
	path = filepath.ToSlash(filepath.Clean(path))

	matches := fieldTransPathRegex.FindStringSubmatch(path)

	if len(matches) != 3 {
		// Fallback to default implementation
		return metadata.NameFromPath(path)
	}

	objectTransName := matches[1] // e.g., "Campaign-en_US"
	fieldName := matches[2]       // e.g., "SomeField__c"

	return metadata.MetadataObjectName(objectTransName + "." + fieldName)
}

func Open(path string) (*CustomFieldTranslation, error) {
	p := &CustomFieldTranslation{}
	return p, metadata.ParseMetadataXml(p, path)
}
