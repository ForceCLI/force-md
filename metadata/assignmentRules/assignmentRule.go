package assignmentRules

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AssignmentRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type AssignmentRules struct {
	metadata.MetadataInfo
	XMLName        xml.Name `xml:"AssignmentRules"`
	Xmlns          string   `xml:"xmlns,attr"`
	AssignmentRule []struct {
		FullName  TextLiteral `xml:"fullName"`
		Active    TextLiteral `xml:"active"`
		RuleEntry []struct {
			AssignedTo            *TextLiteral  `xml:"assignedTo"`
			AssignedToType        *TextLiteral  `xml:"assignedToType"`
			BooleanFilter         *TextLiteral  `xml:"booleanFilter"`
			CriteriaItems         []FilterItem  `xml:"criteriaItems"`
			Formula               *TextLiteral  `xml:"formula"`
			NotifyCcRecipients    *TextLiteral  `xml:"notifyCcRecipients"`
			OverrideExistingTeams *TextLiteral  `xml:"overrideExistingTeams"`
			Team                  []TextLiteral `xml:"team"`
			Template              *TextLiteral  `xml:"template"`
		} `xml:"ruleEntry"`
	} `xml:"assignmentRule"`
}

// FilterItem is one criterion of a rule entry
type FilterItem struct {
	Field      TextLiteral  `xml:"field"`
	Operation  TextLiteral  `xml:"operation"`
	Value      *TextLiteral `xml:"value"`
	ValueField *TextLiteral `xml:"valueField"`
}

func (c *AssignmentRules) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AssignmentRules) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AssignmentRules, error) {
	p := &AssignmentRules{}
	return p, metadata.ParseMetadataXml(p, path)
}
