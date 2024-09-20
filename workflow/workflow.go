package workflow

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type Recipient struct {
	Field     *TextLiteral `xml:"field"`
	Recipient *TextLiteral `xml:"recipient"`
	Type      struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
}

type Alert struct {
	FullName string `xml:"fullName"`
	CcEmails []struct {
		Text string `xml:",chardata"`
	} `xml:"ccEmails"`
	Description   *TextLiteral `xml:"description"`
	Protected     *BooleanText `xml:"protected"`
	Recipients    []Recipient  `xml:"recipients"`
	SenderAddress *TextLiteral `xml:"senderAddress"`
	SenderType    *TextLiteral `xml:"senderType"`
	Template      *TextLiteral `xml:"template"`
}

type FieldUpdate struct {
	FullName           TextLiteral  `xml:"fullName"`
	Description        *TextLiteral `xml:"description"`
	Field              TextLiteral  `xml:"field"`
	Formula            *TextLiteral `xml:"formula"`
	LiteralValue       *TextLiteral `xml:"literalValue"`
	LookupValue        *TextLiteral `xml:"lookupValue"`
	LookupValueType    *TextLiteral `xml:"lookupValueType"`
	Name               TextLiteral  `xml:"name"`
	NotifyAssignee     *BooleanText `xml:"notifyAssignee"`
	Operation          *TextLiteral `xml:"operation"`
	Protected          *BooleanText `xml:"protected"`
	ReevaluateOnChange *BooleanText `xml:"reevaluateOnChange"`
}

type Rule struct {
	FullName string `xml:"fullName"`
	Actions  []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
	} `xml:"actions"`
	Active        BooleanText  `xml:"active"`
	BooleanFilter *TextLiteral `xml:"booleanFilter"`
	CriteriaItems []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Value *TextLiteral `xml:"value"`
	} `xml:"criteriaItems"`
	Description          *TextLiteral `xml:"description"`
	Formula              *TextLiteral `xml:"formula"`
	TriggerType          *TextLiteral `xml:"triggerType"`
	WorkflowTimeTriggers *struct {
		Actions struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"actions"`
		OffsetFromField struct {
			Text string `xml:",chardata"`
		} `xml:"offsetFromField"`
		TimeLength struct {
			Text string `xml:",chardata"`
		} `xml:"timeLength"`
		WorkflowTimeTriggerUnit struct {
			Text string `xml:",chardata"`
		} `xml:"workflowTimeTriggerUnit"`
	} `xml:"workflowTimeTriggers"`
}

type Workflow struct {
	Metadata
	XMLName      xml.Name      `xml:"Workflow"`
	Xmlns        string        `xml:"xmlns,attr"`
	Alerts       []Alert       `xml:"alerts"`
	FieldUpdates []FieldUpdate `xml:"fieldUpdates"`
	Rules        []Rule        `xml:"rules"`
}

func (c *Workflow) SetMetadata(m Metadata) {
	c.Metadata = m
}

func Open(path string) (*Workflow, error) {
	p := &Workflow{}
	return p, internal.ParseMetadataXml(p, path)
}
