package recordtype

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "RecordType"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RecordTypeFilter func(RecordType) bool

type RecordTypeMetadata struct {
	metadata.MetadataInfo
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

func (c *RecordTypeMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *RecordTypeMetadata) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *RecordTypeMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*RecordTypeMetadata, error) {
	p := &RecordTypeMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}
