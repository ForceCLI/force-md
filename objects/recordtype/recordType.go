package recordtype

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects/split"
)

const NAME = "RecordType"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type RecordTypeFilter func(RecordType) bool

type RecordTypeMetadata struct {
	internal.MetadataInfo
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

func (c *RecordTypeMetadata) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *RecordTypeMetadata) NameFromPath(path string) string {
	return split.NameFromPath(path)
}

func (c *RecordTypeMetadata) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*RecordTypeMetadata, error) {
	p := &RecordTypeMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
