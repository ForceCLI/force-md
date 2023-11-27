package field

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type CustomFieldTranslation struct {
	XMLName xml.Name `xml:"CustomFieldTranslation"`
	Xmlns   string   `xml:"xmlns,attr"`
	Field
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

func (p *CustomFieldTranslation) MetaCheck() {}

func Open(path string) (*CustomFieldTranslation, error) {
	p := &CustomFieldTranslation{}
	return p, internal.ParseMetadataXml(p, path)
}
