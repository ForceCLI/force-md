package workflow

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Workflow"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

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
	TargetObject       *TextLiteral `xml:"targetObject"`
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

type Task struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	AssignedTo struct {
		Text string `xml:",chardata"`
	} `xml:"assignedTo"`
	AssignedToType struct {
		Text string `xml:",chardata"`
	} `xml:"assignedToType"`
	DueDateOffset struct {
		Text string `xml:",chardata"`
	} `xml:"dueDateOffset"`
	NotifyAssignee struct {
		Text string `xml:",chardata"`
	} `xml:"notifyAssignee"`
	OffsetFromField struct {
		Text string `xml:",chardata"`
	} `xml:"offsetFromField"`
	Priority struct {
		Text string `xml:",chardata"`
	} `xml:"priority"`
	Protected struct {
		Text string `xml:",chardata"`
	} `xml:"protected"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	Subject struct {
		Text string `xml:",chardata"`
	} `xml:"subject"`
}

type Workflow struct {
	metadata.MetadataInfo
	XMLName      xml.Name      `xml:"Workflow"`
	Xmlns        string        `xml:"xmlns,attr"`
	Alerts       []Alert       `xml:"alerts"`
	FieldUpdates []FieldUpdate `xml:"fieldUpdates"`
	Rules        []Rule        `xml:"rules"`
	Tasks        []Task        `xml:"tasks"`
}

func (c *Workflow) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*Workflow, error) {
	p := &Workflow{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *Workflow) Type() metadata.MetadataType {
	return NAME
}
