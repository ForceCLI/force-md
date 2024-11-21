package restrictionrule

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "RestrictionRule"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type RestrictionRule struct {
	internal.MetadataInfo
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

func (c *RestrictionRule) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*RestrictionRule, error) {
	p := &RestrictionRule{}
	return p, internal.ParseMetadataXml(p, path)
}

func (c *RestrictionRule) Type() internal.MetadataType {
	return NAME
}
