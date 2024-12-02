package restrictionrule

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "RestrictionRule"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RestrictionRule struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"RestrictionRule"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	EnforcementType struct {
		Text string `xml:",chardata"`
	} `xml:"enforcementType"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	RecordFilter struct {
		Text string `xml:",chardata"`
	} `xml:"recordFilter"`
	TargetEntity struct {
		Text string `xml:",chardata"`
	} `xml:"targetEntity"`
	UserCriteria struct {
		Text string `xml:",chardata"`
	} `xml:"userCriteria"`
	Version struct {
		Text string `xml:",chardata"`
	} `xml:"version"`
}

func (c *RestrictionRule) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*RestrictionRule, error) {
	p := &RestrictionRule{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *RestrictionRule) Type() metadata.MetadataType {
	return NAME
}
