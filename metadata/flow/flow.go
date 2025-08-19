package flow

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Flow"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ElementName string

type Start struct {
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	Connector *struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	DoesRequireRecordChangedToMeetCriteria *struct {
		Text string `xml:",chardata"`
	} `xml:"doesRequireRecordChangedToMeetCriteria"`
	FilterLogic   *string      `xml:"filterLogic"`
	FilterFormula *TextLiteral `xml:"filterFormula"`
	Filters       []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operator struct {
			Text string `xml:",chardata"`
		} `xml:"operator"`
		Value *Value `xml:"value"`
	} `xml:"filters"`
	Object            *string `xml:"object"`
	RecordTriggerType *struct {
		Text string `xml:",chardata"`
	} `xml:"recordTriggerType"`
	Schedule *struct {
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
		Name *struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Connector struct {
			IsGoTo *struct {
				Text string `xml:",chardata"`
			} `xml:"isGoTo"`
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		Label *struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		MaxBatchSize *struct {
			Text string `xml:",chardata"`
		} `xml:"maxBatchSize"`
		OffsetNumber *struct {
			Text string `xml:",chardata"`
		} `xml:"offsetNumber"`
		OffsetUnit *struct {
			Text string `xml:",chardata"`
		} `xml:"offsetUnit"`
		PathType    *string `xml:"pathType"`
		RecordField *string `xml:"recordField"`
		TimeSource  *struct {
			Text string `xml:",chardata"`
		} `xml:"timeSource"`
	} `xml:"scheduledPaths"`
	TriggerType *string `xml:"triggerType"`
}

type Rule struct {
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	ConditionLogic string `xml:"conditionLogic"`
	Conditions     []struct {
		LeftValueReference string `xml:"leftValueReference"`
		Operator           string `xml:"operator"`
		RightValue         *Value `xml:"rightValue"`
	} `xml:"conditions"`
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	DoesRequireRecordChangedToMeetCriteria *struct {
		Text string `xml:",chardata"`
	} `xml:"doesRequireRecordChangedToMeetCriteria"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

type Decision struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Description           *TextLiteral           `xml:"description"`
	Name                  ElementName            `xml:"name"`
	Label                 struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LocationX struct {
		Text string `xml:",chardata"`
	} `xml:"locationX"`
	LocationY struct {
		Text string `xml:",chardata"`
	} `xml:"locationY"`
	DefaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"defaultConnector"`
	DefaultConnectorLabel struct {
		Text string `xml:",chardata"`
	} `xml:"defaultConnectorLabel"`
	Rules []Rule `xml:"rules"`
}

type Value struct {
	ElementReference *struct {
		Text string `xml:",chardata"`
	} `xml:"elementReference"`
	StringValue   *TextLiteral `xml:"stringValue"`
	DateTimeValue *TextLiteral `xml:"dateTimeValue"`
	NumberValue   *struct {
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
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
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
	FaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	FilterLogic *string `xml:"filterLogic"`
	Filters     []struct {
		Field    string `xml:"field"`
		Operator string `xml:"operator"`
		Value    *Value `xml:"value"`
	} `xml:"filters"`
	GetFirstRecordOnly *BooleanText `xml:"getFirstRecordOnly"`
	Object             string       `xml:"object"`
	OutputReference    *struct {
		Text string `xml:",chardata"`
	} `xml:"outputReference"`
	QueriedFields []struct {
		Text string `xml:",chardata"`
	} `xml:"queriedFields"`
	OutputAssignments []struct {
		AssignToReference struct {
			Text string `xml:",chardata"`
		} `xml:"assignToReference"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
	} `xml:"outputAssignments"`
	SortField *struct {
		Text string `xml:",chardata"`
	} `xml:"sortField"`
	SortOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"sortOrder"`
	StoreOutputAutomatically *BooleanText `xml:"storeOutputAutomatically"`
}

type RecordRollback struct {
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
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
}

type RecordCreate struct {
	Description *TextLiteral `xml:"description"`
	Name        struct {
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
	AssignRecordIdToReference *struct {
		Text string `xml:",chardata"`
	} `xml:"assignRecordIdToReference"`
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	InputReference *struct {
		Text string `xml:",chardata"`
	} `xml:"inputReference"`
	InputAssignments []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Value struct {
			BooleanValue *struct {
				Text string `xml:",chardata"`
			} `xml:"booleanValue"`
			ElementReference *struct {
				Text string `xml:",chardata"`
			} `xml:"elementReference"`
			StringValue *struct {
				Text string `xml:",chardata"`
			} `xml:"stringValue"`
		} `xml:"value"`
	} `xml:"inputAssignments"`
	Object *struct {
		Text string `xml:",chardata"`
	} `xml:"object"`
	StoreOutputAutomatically *struct {
		Text string `xml:",chardata"`
	} `xml:"storeOutputAutomatically"`
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
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	FilterLogic *string `xml:"filterLogic"`
	Filters     []struct {
		Field    string `xml:"field"`
		Operator string `xml:"operator"`
		Value    *Value `xml:"value"`
	} `xml:"filters"`
	Object         *string `xml:"object"`
	InputReference *struct {
		Text string `xml:",chardata"`
	} `xml:"inputReference"`
}

type RecordUpdate struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Description           *struct {
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
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	FilterLogic *string `xml:"filterLogic"`
	Filters     []struct {
		ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
		Field                 struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operator struct {
			Text string `xml:",chardata"`
		} `xml:"operator"`
		Value *Value `xml:"value"`
	} `xml:"filters"`
	InputAssignments []struct {
		ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
		Field                 struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Value *Value `xml:"value"`
	} `xml:"inputAssignments"`
	Object *struct {
		Text string `xml:",chardata"`
	} `xml:"object"`
	InputReference *struct {
		Text string `xml:",chardata"`
	} `xml:"inputReference"`
}

type Field struct {
	ProcessMetadataValues []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"processMetadataValues"`
	Name                           *string  `xml:"name"`
	ChoiceReferences               []string `xml:"choiceReferences"`
	DataType                       *string  `xml:"dataType"`
	DefaultSelectedChoiceReference *struct {
		Text string `xml:",chardata"`
	} `xml:"defaultSelectedChoiceReference"`
	DefaultValue     *Value `xml:"defaultValue"`
	DataTypeMappings []struct {
		TypeName struct {
			Text string `xml:",chardata"`
		} `xml:"typeName"`
		TypeValue struct {
			Text string `xml:",chardata"`
		} `xml:"typeValue"`
	} `xml:"dataTypeMappings"`
	ExtensionName   *string      `xml:"extensionName"`
	FieldText       *TextLiteral `xml:"fieldText"`
	FieldType       string       `xml:"fieldType"`
	Fields          []Field      `xml:"fields"`
	InputParameters []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"inputParameters"`
	HelpText *struct {
		Text string `xml:",chardata"`
	} `xml:"helpText"`
	InputsOnNextNavToAssocScrn *struct {
		Text string `xml:",chardata"`
	} `xml:"inputsOnNextNavToAssocScrn"`
	IsDisabled *struct {
		ElementReference *struct {
			Text string `xml:",chardata"`
		} `xml:"elementReference"`
		BooleanValue *BooleanText `xml:"booleanValue"`
	} `xml:"isDisabled"`
	IsRequired *struct {
		Text string `xml:",chardata"`
	} `xml:"isRequired"`
	StoreOutputAutomatically *struct {
		Text string `xml:",chardata"`
	} `xml:"storeOutputAutomatically"`
	Scale *struct {
		Text string `xml:",chardata"`
	} `xml:"scale"`
	ValidationRule *struct {
		ErrorMessage struct {
			Text string `xml:",chardata"`
		} `xml:"errorMessage"`
		FormulaExpression TextLiteral `xml:"formulaExpression"`
	} `xml:"validationRule"`
	OutputParameters []struct {
		AssignToReference struct {
			Text string `xml:",chardata"`
		} `xml:"assignToReference"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"outputParameters"`
	RegionContainerType *struct {
		Text string `xml:",chardata"`
	} `xml:"regionContainerType"`
	VisibilityRule *struct {
		ConditionLogic struct {
			Text string `xml:",chardata"`
		} `xml:"conditionLogic"`
		Conditions []struct {
			LeftValueReference struct {
				Text string `xml:",chardata"`
			} `xml:"leftValueReference"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			RightValue *Value `xml:"rightValue"`
		} `xml:"conditions"`
	} `xml:"visibilityRule"`
	ObjectFieldReference *struct {
		Text string `xml:",chardata"`
	} `xml:"objectFieldReference"`
}

type Screen struct {
	Description *TextLiteral `xml:"description"`
	Name        ElementName  `xml:"name"`
	Label       struct {
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
	BackButtonLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"backButtonLabel"`
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	Fields                  []Field      `xml:"fields"`
	HelpText                *TextLiteral `xml:"helpText"`
	NextOrFinishButtonLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"nextOrFinishButtonLabel"`
	PausedText *struct {
		Text string `xml:",chardata"`
	} `xml:"pausedText"`
	ShowFooter struct {
		Text string `xml:",chardata"`
	} `xml:"showFooter"`
	ShowHeader struct {
		Text string `xml:",chardata"`
	} `xml:"showHeader"`
}

type Variable struct {
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Name         string      `xml:"name"`
	DataType     string      `xml:"dataType"`
	IsCollection BooleanText `xml:"isCollection"`
	IsInput      BooleanText `xml:"isInput"`
	IsOutput     BooleanText `xml:"isOutput"`
	ObjectType   *struct {
		Text string `xml:",chardata"`
	} `xml:"objectType"`
	Scale *struct {
		Text string `xml:",chardata"`
	} `xml:"scale"`
	Value *Value `xml:"value"`
}

type Assignment struct {
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
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
		Value             *Value `xml:"value"`
	} `xml:"assignmentItems"`
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
}

type ProcessMetadataValue struct {
	Name  string `xml:"name"`
	Value Value  `xml:"value"`
}

type ActionCall struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Description           *TextLiteral           `xml:"description"`
	Name                  ElementName            `xml:"name"`
	Label                 struct {
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
	ActionType string `xml:"actionType"`
	Connector  *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	FaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	FlowTransactionModel *struct {
		Text string `xml:",chardata"`
	} `xml:"flowTransactionModel"`
	InputParameters []struct {
		ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
		Name                  struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"inputParameters"`
	NameSegment *struct {
		Text string `xml:",chardata"`
	} `xml:"nameSegment"`
	OutputParameters []struct {
		AssignToReference struct {
			Text string `xml:",chardata"`
		} `xml:"assignToReference"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"outputParameters"`
	StoreOutputAutomatically *struct {
		Text string `xml:",chardata"`
	} `xml:"storeOutputAutomatically"`
	VersionSegment *struct {
		Text string `xml:",chardata"`
	} `xml:"versionSegment"`
}

type Flow struct {
	metadata.MetadataInfo
	XMLName     xml.Name     `xml:"Flow"`
	Xmlns       string       `xml:"xmlns,attr"`
	Xsi         string       `xml:"xmlns:xsi,attr,omitempty"`
	ActionCalls []ActionCall `xml:"actionCalls"`
	ApiVersion  *struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Assignments []Assignment `xml:"assignments"`
	Choices     []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		ChoiceText struct {
			Text string `xml:",chardata"`
		} `xml:"choiceText"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Value *Value `xml:"value"`
	} `xml:"choices"`
	CollectionProcessors []struct {
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
			RightValue *Value `xml:"rightValue"`
		} `xml:"conditions"`
		Connector struct {
			IsGoTo *struct {
				Text string `xml:",chardata"`
			} `xml:"isGoTo"`
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
	} `xml:"collectionProcessors"`
	CustomErrors []struct {
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
		CustomErrorMessages struct {
			ErrorMessage struct {
				Text string `xml:",chardata"`
			} `xml:"errorMessage"`
			IsFieldError struct {
				Text string `xml:",chardata"`
			} `xml:"isFieldError"`
		} `xml:"customErrorMessages"`
	} `xml:"customErrors"`
	Constants []struct {
		Description *TextLiteral `xml:"description"`
		Name        struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Value Value `xml:"value"`
	} `xml:"constants"`
	Decisions         []Decision   `xml:"decisions"`
	Description       *TextLiteral `xml:"description"`
	DynamicChoiceSets []struct {
		Description *TextLiteral `xml:"description"`
		Name        struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		CollectionReference *struct {
			Text string `xml:",chardata"`
		} `xml:"collectionReference"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		DisplayField struct {
			Text string `xml:",chardata"`
			Nil  string `xml:"xsi:nil,attr,omitempty"`
		} `xml:"displayField"`
		FilterLogic *string `xml:"filterLogic"`
		Filters     []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operator struct {
				Text string `xml:",chardata"`
			} `xml:"operator"`
			Value *Value `xml:"value"`
		} `xml:"filters"`
		Limit *struct {
			Text string `xml:",chardata"`
		} `xml:"limit"`
		Object struct {
			Text string `xml:",chardata"`
			Nil  string `xml:"xsi:nil,attr,omitempty"`
		} `xml:"object"`
		OutputAssignments *struct {
			AssignToReference struct {
				Text string `xml:",chardata"`
			} `xml:"assignToReference"`
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
		} `xml:"outputAssignments"`
		PicklistField *struct {
			Text string `xml:",chardata"`
		} `xml:"picklistField"`
		PicklistObject *struct {
			Text string `xml:",chardata"`
		} `xml:"picklistObject"`
		SortField *struct {
			Text string `xml:",chardata"`
		} `xml:"sortField"`
		SortOrder *struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
		ValueField *struct {
			Text string `xml:",chardata"`
		} `xml:"valueField"`
	} `xml:"dynamicChoiceSets"`
	Environments *struct {
		Text string `xml:",chardata"`
	} `xml:"environments"`
	Formulas []struct {
		ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
		Description           *TextLiteral           `xml:"description"`
		Name                  struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		DataType struct {
			Text string `xml:",chardata"`
		} `xml:"dataType"`
		Expression *TextLiteral `xml:"expression"`
		Scale      *struct {
			Text string `xml:",chardata"`
		} `xml:"scale"`
	} `xml:"formulas"`
	InterviewLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"interviewLabel"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Loops []struct {
		Description *TextLiteral `xml:"description"`
		Name        struct {
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
		NoMoreValuesConnector *struct {
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"noMoreValuesConnector"`
	} `xml:"loops"`
	ProcessMetadataValues []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"processMetadataValues"`
	ProcessType struct {
		Text string `xml:",chardata"`
	} `xml:"processType"`
	RecordCreates   []RecordCreate   `xml:"recordCreates"`
	RecordDeletes   []RecordDelete   `xml:"recordDeletes"`
	RecordLookups   []RecordLookup   `xml:"recordLookups"`
	RecordRollbacks []RecordRollback `xml:"recordRollbacks"`
	RecordUpdates   []RecordUpdate   `xml:"recordUpdates"`
	RunInMode       *struct {
		Text string `xml:",chardata"`
	} `xml:"runInMode"`
	Screens               []Screen `xml:"screens"`
	Start                 *Start   `xml:"start"`
	StartElementReference *struct {
		Text string `xml:",chardata"`
	} `xml:"startElementReference"`
	Status *struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	Subflows []struct {
		Description *TextLiteral `xml:"description"`
		Name        struct {
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
		Connector *struct {
			IsGoTo *struct {
				Text string `xml:",chardata"`
			} `xml:"isGoTo"`
			TargetReference ElementName `xml:"targetReference"`
		} `xml:"connector"`
		FlowName struct {
			Text string `xml:",chardata"`
		} `xml:"flowName"`
		InputAssignments []struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Value Value `xml:"value"`
		} `xml:"inputAssignments"`
		OutputAssignments []struct {
			AssignToReference struct {
				Text string `xml:",chardata"`
			} `xml:"assignToReference"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
		} `xml:"outputAssignments"`
		StoreOutputAutomatically *struct {
			Text string `xml:",chardata"`
		} `xml:"storeOutputAutomatically"`
	} `xml:"subflows"`
	TextTemplates []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		IsViewedAsPlainText struct {
			Text string `xml:",chardata"`
		} `xml:"isViewedAsPlainText"`
		Text TextLiteral `xml:"text"`
	} `xml:"textTemplates"`
	TriggerOrder *struct {
		Text string `xml:",chardata"`
	} `xml:"triggerOrder"`
	Variables []Variable `xml:"variables"`
	Waits     []struct {
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
				IsGoTo *struct {
					Text string `xml:",chardata"`
				} `xml:"isGoTo"`
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
}

func (c *Flow) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Flow) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Flow, error) {
	p := &Flow{}
	return p, metadata.ParseMetadataXml(p, path)
}
