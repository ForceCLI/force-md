package flow

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type ElementName string

type Start struct {
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	Connector struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FilterLogic struct {
		Text string `xml:",chardata"`
	} `xml:"filterLogic"`
	Filters []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operator struct {
			Text string `xml:",chardata"`
		} `xml:"operator"`
		Value struct {
			StringValue struct {
				Text string `xml:",chardata"`
			} `xml:"stringValue"`
			BooleanValue struct {
				Text string `xml:",chardata"`
			} `xml:"booleanValue"`
			DateTimeValue struct {
				Text string `xml:",chardata"`
			} `xml:"dateTimeValue"`
		} `xml:"value"`
	} `xml:"filters"`
	Object            string `xml:"object"`
	RecordTriggerType struct {
		Text string `xml:",chardata"`
	} `xml:"recordTriggerType"`
	TriggerType string `xml:"triggerType"`
	Schedule    struct {
		Frequency struct {
			Text string `xml:",chardata"`
		} `xml:"frequency"`
		StartDate struct {
			Text string `xml:",chardata"`
		} `xml:"startDate"`
		StartTime struct {
			Text string `xml:",chardata"`
		} `xml:"startTime"`
	} `xml:"schedule"`
	ScheduledPaths []struct {
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		PathType string `xml:"pathType"`
	} `xml:"scheduledPaths"`
	DoesRequireRecordChangedToMeetCriteria struct {
		Text string `xml:",chardata"`
	} `xml:"doesRequireRecordChangedToMeetCriteria"`
}

type Rule struct {
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	ConditionLogic string `xml:"conditionLogic"`
	Conditions     []struct {
		LeftValueReference string `xml:"leftValueReference"`
		Operator           string `xml:"operator"`
		RightValue         Value  `xml:"rightValue"`
	} `xml:"conditions"`
	Connector struct {
		TargetReference ElementName `xml:"targetReference"`
		IsGoTo          struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
	} `xml:"connector"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

type Decision struct {
	Name  ElementName `xml:"name"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	DefaultConnectorLabel struct {
		Text string `xml:",chardata"`
	} `xml:"defaultConnectorLabel"`
	Rules            []Rule `xml:"rules"`
	DefaultConnector struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"defaultConnector"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

type Value struct {
	ElementReference *struct {
		Text string `xml:",chardata"`
	} `xml:"elementReference"`
	StringValue *struct {
		Text string `xml:",chardata"`
	} `xml:"stringValue"`
	NumberValue *struct {
		Text string `xml:",chardata"`
	} `xml:"numberValue"`
	BooleanValue *BooleanText `xml:"booleanValue"`
}

func (v Value) String() string {
	if v.ElementReference != nil {
		return v.ElementReference.Text
	}
	if v.StringValue != nil {
		return v.StringValue.Text
	}
	if v.BooleanValue != nil {
		return v.BooleanValue.String()
	}
	return ""
}

type RecordLookup struct {
	Name  ElementName `xml:"name"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	AssignNullValuesIfNoRecordsFound BooleanText `xml:"assignNullValuesIfNoRecordsFound"`
	Connector                        struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FilterLogic string `xml:"filterLogic"`
	Filters     []struct {
		Field    string `xml:"field"`
		Operator string `xml:"operator"`
		Value    Value  `xml:"value"`
	} `xml:"filters"`
	GetFirstRecordOnly       BooleanText `xml:"getFirstRecordOnly"`
	Object                   string      `xml:"object"`
	StoreOutputAutomatically BooleanText `xml:"storeOutputAutomatically"`
	FaultConnector           struct {
		TargetReference ElementName `xml:"targetReference"`
		IsGoTo          struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
	} `xml:"faultConnector"`
	OutputReference struct {
		Text string `xml:",chardata"`
	} `xml:"outputReference"`
	QueriedFields []struct {
		Text string `xml:",chardata"`
	} `xml:"queriedFields"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	OutputAssignments []struct {
		AssignToReference struct {
			Text string `xml:",chardata"`
		} `xml:"assignToReference"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
	} `xml:"outputAssignments"`
	SortField struct {
		Text string `xml:",chardata"`
	} `xml:"sortField"`
	SortOrder struct {
		Text string `xml:",chardata"`
	} `xml:"sortOrder"`
}

type RecordDelete struct {
	Name  ElementName `xml:"name"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	Connector struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FaultConnector struct {
		IsGoTo struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	InputReference struct {
		Text string `xml:",chardata"`
	} `xml:"inputReference"`
}

type Field struct {
	Name            string  `xml:"name"`
	ExtensionName   string  `xml:"extensionName"`
	FieldType       string  `xml:"fieldType"`
	Fields          []Field `xml:"fields"`
	InputParameters []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value struct {
			ElementReference struct {
				Text string `xml:",chardata"`
			} `xml:"elementReference"`
			StringValue struct {
				Text string `xml:",chardata"`
			} `xml:"stringValue"`
			BooleanValue struct {
				Text string `xml:",chardata"`
			} `xml:"booleanValue"`
		} `xml:"value"`
	} `xml:"inputParameters"`
	InputsOnNextNavToAssocScrn struct {
		Text string `xml:",chardata"`
	} `xml:"inputsOnNextNavToAssocScrn"`
	IsRequired struct {
		Text string `xml:",chardata"`
	} `xml:"isRequired"`
	StoreOutputAutomatically struct {
		Text string `xml:",chardata"`
	} `xml:"storeOutputAutomatically"`
	FieldText    string `xml:"fieldText"`
	DataType     string `xml:"dataType"`
	DefaultValue struct {
		ElementReference struct {
			Text string `xml:",chardata"`
		} `xml:"elementReference"`
		StringValue struct {
			Text string `xml:",chardata"`
		} `xml:"stringValue"`
	} `xml:"defaultValue"`
	ValidationRule struct {
		ErrorMessage struct {
			Text string `xml:",chardata"`
		} `xml:"errorMessage"`
		FormulaExpression struct {
			Text string `xml:",chardata"`
		} `xml:"formulaExpression"`
	} `xml:"validationRule"`
	ChoiceReferences []string `xml:"choiceReferences"`
	OutputParameters []struct {
		AssignToReference struct {
			Text string `xml:",chardata"`
		} `xml:"assignToReference"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"outputParameters"`
	HelpText struct {
		Text string `xml:",chardata"`
	} `xml:"helpText"`
	VisibilityRule struct {
		ConditionLogic struct {
			Text string `xml:",chardata"`
		} `xml:"conditionLogic"`
		Conditions struct {
			LeftValueReference struct {
				Text string `xml:",chardata"`
			} `xml:"leftValueReference"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			RightValue struct {
				BooleanValue struct {
					Text string `xml:",chardata"`
				} `xml:"booleanValue"`
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
			} `xml:"rightValue"`
		} `xml:"conditions"`
	} `xml:"visibilityRule"`
	ObjectFieldReference struct {
		Text string `xml:",chardata"`
	} `xml:"objectFieldReference"`
	Scale struct {
		Text string `xml:",chardata"`
	} `xml:"scale"`
}

type Screen struct {
	Name  ElementName `xml:"name"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	AllowBack struct {
		Text string `xml:",chardata"`
	} `xml:"allowBack"`
	AllowFinish BooleanText `xml:"allowFinish"`
	AllowPause  struct {
		Text string `xml:",chardata"`
	} `xml:"allowPause"`
	Fields     []Field `xml:"fields"`
	ShowFooter struct {
		Text string `xml:",chardata"`
	} `xml:"showFooter"`
	ShowHeader struct {
		Text string `xml:",chardata"`
	} `xml:"showHeader"`
	Connector struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	NextOrFinishButtonLabel struct {
		Text string `xml:",chardata"`
	} `xml:"nextOrFinishButtonLabel"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	HelpText struct {
		Text string `xml:",chardata"`
	} `xml:"helpText"`
	PausedText struct {
		Text string `xml:",chardata"`
	} `xml:"pausedText"`
}

type Variable struct {
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Name         string `xml:"name"`
	DataType     string `xml:"dataType"`
	IsCollection struct {
		Text string `xml:",chardata"`
	} `xml:"isCollection"`
	IsInput    BooleanText `xml:"isInput"`
	IsOutput   BooleanText `xml:"isOutput"`
	ObjectType struct {
		Text string `xml:",chardata"`
	} `xml:"objectType"`
	Scale struct {
		Text string `xml:",chardata"`
	} `xml:"scale"`
	Value struct {
		StringValue struct {
			Text string `xml:",chardata"`
		} `xml:"stringValue"`
		ElementReference struct {
			Text string `xml:",chardata"`
		} `xml:"elementReference"`
		BooleanValue struct {
			Text string `xml:",chardata"`
		} `xml:"booleanValue"`
	} `xml:"value"`
}

type Assignment struct {
	Name  ElementName `xml:"name"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	AssignmentItems []struct {
		AssignToReference string `xml:"assignToReference"`
		Operator          string `xml:"operator"`
		Value             Value  `xml:"value"`
	} `xml:"assignmentItems"`
	Connector struct {
		TargetReference ElementName `xml:"targetReference"`
		IsGoTo          struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
	} `xml:"connector"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

type Flow struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"Flow"`
	Xmlns      string   `xml:"xmlns,attr"`
	Xsi        string   `xml:"xsi,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Environments struct {
		Text string `xml:",chardata"`
	} `xml:"environments"`
	Formulas []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Expression struct {
			Text string `xml:",chardata"`
		} `xml:"expression"`
		Description struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
	} `xml:"formulas"`
	InterviewLabel struct {
		Text string `xml:",chardata"`
	} `xml:"interviewLabel"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	ProcessMetadataValues []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value struct {
			StringValue struct {
				Text string `xml:",chardata"`
			} `xml:"stringValue"`
			BooleanValue struct {
				Text string `xml:",chardata"`
			} `xml:"booleanValue"`
		} `xml:"value"`
	} `xml:"processMetadataValues"`
	ProcessType struct {
		Text string `xml:",chardata"`
	} `xml:"processType"`
	Start  Start `xml:"start"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	Subflows []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		FlowName struct {
			Text string `xml:",chardata"`
		} `xml:"flowName"`
		InputAssignments []struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Value struct {
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
			} `xml:"value"`
		} `xml:"inputAssignments"`
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		OutputAssignments []struct {
			AssignToReference struct {
				Text string `xml:",chardata"`
			} `xml:"assignToReference"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
		} `xml:"outputAssignments"`
	} `xml:"subflows"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Loops struct {
		Description struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		CollectionReference struct {
			Text string `xml:",chardata"`
		} `xml:"collectionReference"`
		IterationOrder struct {
			Text string `xml:",chardata"`
		} `xml:"iterationOrder"`
		NextValueConnector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"nextValueConnector"`
		NoMoreValuesConnector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"noMoreValuesConnector"`
	} `xml:"loops"`
	RecordUpdates []struct {
		Description struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		FilterLogic struct {
			Text string `xml:",chardata"`
		} `xml:"filterLogic"`
		Filters []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			Value struct {
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
				StringValue struct {
					Text string `xml:",chardata"`
				} `xml:"stringValue"`
				BooleanValue struct {
					Text string `xml:",chardata"`
				} `xml:"booleanValue"`
			} `xml:"value"`
		} `xml:"filters"`
		InputAssignments []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Value struct {
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
				BooleanValue struct {
					Text string `xml:",chardata"`
				} `xml:"booleanValue"`
				StringValue struct {
					Text string `xml:",chardata"`
				} `xml:"stringValue"`
			} `xml:"value"`
		} `xml:"inputAssignments"`
		Object struct {
			Text string `xml:",chardata"`
		} `xml:"object"`
		InputReference struct {
			Text string `xml:",chardata"`
		} `xml:"inputReference"`
		FaultConnector struct {
			IsGoTo struct {
				Text string `xml:",chardata"`
			} `xml:"isGoTo"`
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"faultConnector"`
	} `xml:"recordUpdates"`
	Variables     []Variable     `xml:"variables"`
	Decisions     []Decision     `xml:"decisions"`
	Screens       []Screen       `xml:"screens"`
	RecordLookups []RecordLookup `xml:"recordLookups"`
	Assignments   []Assignment   `xml:"assignments"`
	Constants     struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Value struct {
			StringValue struct {
				Text string `xml:",chardata"`
			} `xml:"stringValue"`
		} `xml:"value"`
	} `xml:"constants"`
	DynamicChoiceSets []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		DisplayField struct {
			Text string `xml:",chardata"`
			Nil  string `xml:"nil,attr"`
		} `xml:"displayField"`
		Object struct {
			Text string `xml:",chardata"`
			Nil  string `xml:"nil,attr"`
		} `xml:"object"`
		PicklistField struct {
			Text string `xml:",chardata"`
		} `xml:"picklistField"`
		PicklistObject struct {
			Text string `xml:",chardata"`
		} `xml:"picklistObject"`
		FilterLogic struct {
			Text string `xml:",chardata"`
		} `xml:"filterLogic"`
		Filters []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			Value struct {
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
				StringValue struct {
					Text string `xml:",chardata"`
				} `xml:"stringValue"`
			} `xml:"value"`
		} `xml:"filters"`
		SortField struct {
			Text string `xml:",chardata"`
		} `xml:"sortField"`
		SortOrder struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
		ValueField struct {
			Text string `xml:",chardata"`
		} `xml:"valueField"`
		OutputAssignments struct {
			AssignToReference struct {
				Text string `xml:",chardata"`
			} `xml:"assignToReference"`
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
		} `xml:"outputAssignments"`
		CollectionReference struct {
			Text string `xml:",chardata"`
		} `xml:"collectionReference"`
		Description struct {
			Text string `xml:",chardata"`
		} `xml:"description"`
	} `xml:"dynamicChoiceSets"`
	RecordCreates []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		FaultConnector struct {
			IsGoTo struct {
				Text string `xml:",chardata"`
			} `xml:"isGoTo"`
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"faultConnector"`
		InputReference struct {
			Text string `xml:",chardata"`
		} `xml:"inputReference"`
		InputAssignments []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Value struct {
				BooleanValue struct {
					Text string `xml:",chardata"`
				} `xml:"booleanValue"`
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
				StringValue struct {
					Text string `xml:",chardata"`
				} `xml:"stringValue"`
			} `xml:"value"`
		} `xml:"inputAssignments"`
		Object struct {
			Text string `xml:",chardata"`
		} `xml:"object"`
		StoreOutputAutomatically struct {
			Text string `xml:",chardata"`
		} `xml:"storeOutputAutomatically"`
	} `xml:"recordCreates"`
	RunInMode struct {
		Text string `xml:",chardata"`
	} `xml:"runInMode"`
	Choices []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		ChoiceText struct {
			Text string `xml:",chardata"`
		} `xml:"choiceText"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Value struct {
			StringValue struct {
				Text string `xml:",chardata"`
			} `xml:"stringValue"`
			NumberValue struct {
				Text string `xml:",chardata"`
			} `xml:"numberValue"`
		} `xml:"value"`
	} `xml:"choices"`
	ActionCalls []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		ActionName struct {
			Text string `xml:",chardata"`
		} `xml:"actionName"`
		ActionType struct {
			Text string `xml:",chardata"`
		} `xml:"actionType"`
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		FaultConnector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"faultConnector"`
		FlowTransactionModel struct {
			Text string `xml:",chardata"`
		} `xml:"flowTransactionModel"`
		InputParameters []struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Value struct {
				ElementReference struct {
					Text string `xml:",chardata"`
				} `xml:"elementReference"`
				StringValue struct {
					Text string `xml:",chardata"`
				} `xml:"stringValue"`
				BooleanValue struct {
					Text string `xml:",chardata"`
				} `xml:"booleanValue"`
			} `xml:"value"`
		} `xml:"inputParameters"`
		NameSegment struct {
			Text string `xml:",chardata"`
		} `xml:"nameSegment"`
		OutputParameters struct {
			AssignToReference struct {
				Text string `xml:",chardata"`
			} `xml:"assignToReference"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
		} `xml:"outputParameters"`
		VersionSegment struct {
			Text string `xml:",chardata"`
		} `xml:"versionSegment"`
		StoreOutputAutomatically struct {
			Text string `xml:",chardata"`
		} `xml:"storeOutputAutomatically"`
	} `xml:"actionCalls"`
	RecordDeletes         []RecordDelete `xml:"recordDeletes"`
	StartElementReference struct {
		Text string `xml:",chardata"`
	} `xml:"startElementReference"`
	TextTemplates []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		IsViewedAsPlainText struct {
			Text string `xml:",chardata"`
		} `xml:"isViewedAsPlainText"`
		Text struct {
			Text string `xml:",chardata"`
		} `xml:"text"`
	} `xml:"textTemplates"`
	RecordRollbacks struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
	} `xml:"recordRollbacks"`
	Waits struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		ElementSubtype struct {
			Text string `xml:",chardata"`
		} `xml:"elementSubtype"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		DefaultConnectorLabel struct {
			Text string `xml:",chardata"`
		} `xml:"defaultConnectorLabel"`
		WaitEvents struct {
			ConditionLogic struct {
				Text string `xml:",chardata"`
			} `xml:"conditionLogic"`
			Connector struct {
				TargetReference ElementName `xml:"targetReference"`
			} `xml:"connector"`
			Label struct {
				Text string `xml:",chardata"`
			} `xml:"label"`
			Offset struct {
				Text string `xml:",chardata"`
			} `xml:"offset"`
			OffsetUnit struct {
				Text string `xml:",chardata"`
			} `xml:"offsetUnit"`
		} `xml:"waitEvents"`
	} `xml:"waits"`
	CollectionProcessors struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		ElementSubtype struct {
			Text string `xml:",chardata"`
		} `xml:"elementSubtype"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"locationX"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"locationY"`
		AssignNextValueToReference struct {
			Text string `xml:",chardata"`
		} `xml:"assignNextValueToReference"`
		CollectionProcessorType struct {
			Text string `xml:",chardata"`
		} `xml:"collectionProcessorType"`
		CollectionReference struct {
			Text string `xml:",chardata"`
		} `xml:"collectionReference"`
		ConditionLogic struct {
			Text string `xml:",chardata"`
		} `xml:"conditionLogic"`
		Conditions struct {
			LeftValueReference struct {
				Text string `xml:",chardata"`
			} `xml:"leftValueReference"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			RightValue struct {
				StringValue struct {
					Text string `xml:",chardata"`
				} `xml:"stringValue"`
			} `xml:"rightValue"`
		} `xml:"conditions"`
		Connector struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
	} `xml:"collectionProcessors"`
}

func (c *Flow) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*Flow, error) {
	p := &Flow{}
	return p, internal.ParseMetadataXml(p, path)
}
