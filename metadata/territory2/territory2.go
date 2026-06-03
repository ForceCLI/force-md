package territory2

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Territory2"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RuleAssociation struct {
	RuleName struct {
		Text string `xml:",chardata"`
	} `xml:"ruleName"`
	Inherited struct {
		Text string `xml:",chardata"`
	} `xml:"inherited"`
}

type Territory2 struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Territory2"`
	Xmlns   string   `xml:"xmlns,attr"`
	Name    struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	AccountAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"accountAccessLevel"`
	CaseAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"caseAccessLevel"`
	ContactAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"contactAccessLevel"`
	OpportunityAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"opportunityAccessLevel"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	ParentTerritory *struct {
		Text string `xml:",chardata"`
	} `xml:"parentTerritory"`
	Territory2Type struct {
		Text string `xml:",chardata"`
	} `xml:"territory2Type"`
	RuleAssociations []RuleAssociation `xml:"ruleAssociations"`
}

func (c *Territory2) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Territory2) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Territory2, error) {
	p := &Territory2{}
	return p, metadata.ParseMetadataXml(p, path)
}
