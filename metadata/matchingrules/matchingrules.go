package matchingrules

import (
	"encoding/xml"

	"github.com/pkg/errors"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "MatchingRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type MatchingRules struct {
	metadata.MetadataInfo
	XMLName       xml.Name       `xml:"MatchingRules"`
	Xmlns         string         `xml:"xmlns,attr"`
	MatchingRules []MatchingRule `xml:"matchingRules"`
}

type MatchingRule struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	MatchingRuleItems []struct {
		BlankValueBehavior struct {
			Text string `xml:",chardata"`
		} `xml:"blankValueBehavior"`
		FieldName struct {
			Text string `xml:",chardata"`
		} `xml:"fieldName"`
		MatchingMethod struct {
			Text string `xml:",chardata"`
		} `xml:"matchingMethod"`
	} `xml:"matchingRuleItems"`
	RuleStatus struct {
		Text string `xml:",chardata"`
	} `xml:"ruleStatus"`
	BooleanFilter struct {
		Text string `xml:",chardata"`
	} `xml:"booleanFilter"`
}

func (c *MatchingRules) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*MatchingRules, error) {
	p := &MatchingRules{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *MatchingRules) Type() metadata.MetadataType {
	return NAME
}

func (s *MatchingRules) GetMatchingRules() []MatchingRule {
	return s.MatchingRules
}

func (p *MatchingRules) DeleteRule(ruleName string) error {
	found := false
	newRules := p.MatchingRules[:0]
	for _, f := range p.MatchingRules {
		if f.FullName.Text == ruleName {
			found = true
		} else {
			newRules = append(newRules, f)
		}
	}
	if !found {
		return errors.New("rule not found")
	}
	p.MatchingRules = newRules
	return nil
}
