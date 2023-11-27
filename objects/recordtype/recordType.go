package recordtype

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type RecordTypeFilter func(RecordType) bool

type RecordTypeMetadata struct {
	XMLName xml.Name `xml:"RecordType"`
	Xmlns   string   `xml:"xmlns,attr"`
	RecordType
}

type Picklist struct {
	Picklist string             `xml:"picklist"`
	Values   ValueSetOptionList `xml:"values"`
}

type PicklistList []Picklist

type ValueSetOption struct {
	FullName string      `xml:"fullName"`
	Default  BooleanText `xml:"default"`
}

type ValueSetOptionList []ValueSetOption

type RecordType struct {
	FullName        string      `xml:"fullName"`
	Active          BooleanText `xml:"active"`
	BusinessProcess *struct {
		Text string `xml:",chardata"`
	} `xml:"businessProcess"`
	CompactLayoutAssignment *struct {
		Text string `xml:",chardata"`
	} `xml:"compactLayoutAssignment"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	PicklistValues PicklistList `xml:"picklistValues"`
}

func (p *RecordTypeMetadata) MetaCheck() {}

func Open(path string) (*RecordTypeMetadata, error) {
	p := &RecordTypeMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
