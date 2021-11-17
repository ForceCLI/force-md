package workflow

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type Recipient struct {
	Field struct {
		Text string `xml:",chardata"`
	} `xml:"field"`
	Type struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	Recipient struct {
		Text string `xml:",chardata"`
	} `xml:"recipient"`
}

type Alert struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Protected struct {
		Text string `xml:",chardata"`
	} `xml:"protected"`
	Recipients []Recipient `xml:"recipients"`
	SenderType struct {
		Text string `xml:",chardata"`
	} `xml:"senderType"`
	Template struct {
		Text string `xml:",chardata"`
	} `xml:"template"`
	CcEmails []struct {
		Text string `xml:",chardata"`
	} `xml:"ccEmails"`
}

type FieldUpdate struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Field struct {
		Text string `xml:",chardata"`
	} `xml:"field"`
	LookupValue struct {
		Text string `xml:",chardata"`
	} `xml:"lookupValue"`
	LookupValueType struct {
		Text string `xml:",chardata"`
	} `xml:"lookupValueType"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	NotifyAssignee struct {
		Text string `xml:",chardata"`
	} `xml:"notifyAssignee"`
	Operation struct {
		Text string `xml:",chardata"`
	} `xml:"operation"`
	Protected struct {
		Text string `xml:",chardata"`
	} `xml:"protected"`
	Formula struct {
		Text string `xml:",chardata"`
	} `xml:"formula"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	LiteralValue struct {
		Text string `xml:",chardata"`
	} `xml:"literalValue"`
}

type Rule struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Actions []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
	} `xml:"actions"`
	Active        BooleanText `xml:"active"`
	CriteriaItems []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Value struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"criteriaItems"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	TriggerType struct {
		Text string `xml:",chardata"`
	} `xml:"triggerType"`
	Formula struct {
		Text string `xml:",chardata"`
	} `xml:"formula"`
	BooleanFilter struct {
		Text string `xml:",chardata"`
	} `xml:"booleanFilter"`
	WorkflowTimeTriggers struct {
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
	XMLName      xml.Name      `xml:"Workflow"`
	Xmlns        string        `xml:"xmlns,attr"`
	FieldUpdates []FieldUpdate `xml:"fieldUpdates"`
	Rules        []Rule        `xml:"rules"`
	Alerts       []Alert       `xml:"alerts"`
}

func (p *Workflow) MetaCheck() {}

func Open(path string) (*Workflow, error) {
	p := &Workflow{}
	return p, internal.ParseMetadataXml(p, path)
}
