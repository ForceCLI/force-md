package escalationRules

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "EscalationRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type EscalationRules struct {
	metadata.MetadataInfo
	XMLName        xml.Name `xml:"EscalationRules"`
	Xmlns          string   `xml:"xmlns,attr"`
	EscalationRule []struct {
		FullName  TextLiteral  `xml:"fullName"`
		Active    *TextLiteral `xml:"active"`
		RuleEntry []struct {
			BooleanFilter                 *TextLiteral `xml:"booleanFilter"`
			BusinessHours                 *TextLiteral `xml:"businessHours"`
			BusinessHoursSource           *TextLiteral `xml:"businessHoursSource"`
			CriteriaItems                 []FilterItem `xml:"criteriaItems"`
			DisableEscalationWhenModified *TextLiteral `xml:"disableEscalationWhenModified"`
			EscalationAction              []struct {
				AssignedTo          *TextLiteral  `xml:"assignedTo"`
				AssignedToTemplate  *TextLiteral  `xml:"assignedToTemplate"`
				AssignedToType      *TextLiteral  `xml:"assignedToType"`
				MinutesToEscalation *TextLiteral  `xml:"minutesToEscalation"`
				NotifyCaseOwner     *TextLiteral  `xml:"notifyCaseOwner"`
				NotifyEmail         []TextLiteral `xml:"notifyEmail"`
				NotifyTo            *TextLiteral  `xml:"notifyTo"`
				NotifyToTemplate    *TextLiteral  `xml:"notifyToTemplate"`
			} `xml:"escalationAction"`
			EscalationStartTime *TextLiteral `xml:"escalationStartTime"`
			Formula             *TextLiteral `xml:"formula"`
		} `xml:"ruleEntry"`
	} `xml:"escalationRule"`
}

// FilterItem is one criterion of a rule entry
type FilterItem struct {
	Field      TextLiteral  `xml:"field"`
	Operation  TextLiteral  `xml:"operation"`
	Value      *TextLiteral `xml:"value"`
	ValueField *TextLiteral `xml:"valueField"`
}

func (c *EscalationRules) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EscalationRules) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*EscalationRules, error) {
	p := &EscalationRules{}
	return p, metadata.ParseMetadataXml(p, path)
}
