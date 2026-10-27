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
	ConditionLogic string      `xml:"conditionLogic"`
	Conditions     []Condition `xml:"conditions"`
	Connector      *struct {
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

// Condition is one condition an element evaluates. Decision rules and
// Collection Filters share the shape.
type Condition struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	LeftValueReference    string                 `xml:"leftValueReference"`
	Operator              string                 `xml:"operator"`
	RightValue            *Value                 `xml:"rightValue"`
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
	InputConfiguratorMode *TextLiteral `xml:"inputConfiguratorMode"`
	StringValue           *TextLiteral `xml:"stringValue"`
	DateValue             *TextLiteral `xml:"dateValue"`
	DateTimeValue         *TextLiteral `xml:"dateTimeValue"`
	NumberValue           *struct {
		Text string `xml:",chardata"`
	} `xml:"numberValue"`
	BooleanValue *BooleanText `xml:"booleanValue"`
	// An action input that takes a collection lists its members as
	// collectionElements, each a literal or an element reference.
	CollectionElements []Value `xml:"collectionElements"`
	// A Transform maps its output through a formula rather than a literal or
	// a reference, carrying the expression and its result type in the value.
	FormulaDataType   *TextLiteral `xml:"formulaDataType"`
	FormulaExpression *TextLiteral `xml:"formulaExpression"`
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

// CollectionProcessor is a Collection Filter or Collection Sort element,
// distinguished by CollectionProcessorType. A filter keeps the members of
// CollectionReference that satisfy Conditions (or Formula); a sort orders them
// by SortOptions and keeps at most Limit of them.
type CollectionProcessor struct {
	Description *TextLiteral `xml:"description"`
	Name        struct {
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
	Conditions  []Condition  `xml:"conditions"`
	Formula     *TextLiteral `xml:"formula"`
	Limit       *IntegerText `xml:"limit"`
	SortOptions []SortOption `xml:"sortOptions"`
	Connector   struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
}

// SortOption is one level of a Collection Sort's ordering, applied in the order
// the options appear.
type SortOption struct {
	DoesPutEmptyStringAndNullFirst BooleanText `xml:"doesPutEmptyStringAndNullFirst"`
	SortField                      struct {
		Text string `xml:",chardata"`
	} `xml:"sortField"`
	SortOrder struct {
		Text string `xml:",chardata"`
	} `xml:"sortOrder"`
}

// CustomError is a Custom Error element. It ends the path it sits on and
// blocks the save, reporting each of its messages against either the record or
// a named field.
type CustomError struct {
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
	CustomErrorMessages []CustomErrorMessage `xml:"customErrorMessages"`
}

// CustomErrorMessage is one message a Custom Error element reports.
// FieldSelection names the field the message attaches to and is set only when
// IsFieldError is true.
type CustomErrorMessage struct {
	ErrorMessage struct {
		Text string `xml:",chardata"`
	} `xml:"errorMessage"`
	FieldSelection *struct {
		Text string `xml:",chardata"`
	} `xml:"fieldSelection"`
	IsFieldError BooleanText `xml:"isFieldError"`
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
		ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
		Field                 struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Value *Value `xml:"value"`
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
	LocationX NillableText `xml:"locationX"`
	LocationY NillableText `xml:"locationY"`
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
	Name             *string  `xml:"name"`
	ChoiceReferences []string `xml:"choiceReferences"`
	DataType         *string  `xml:"dataType"`
	DataTypeMappings []struct {
		TypeName struct {
			Text string `xml:",chardata"`
		} `xml:"typeName"`
		TypeValue struct {
			Text string `xml:",chardata"`
		} `xml:"typeValue"`
	} `xml:"dataTypeMappings"`
	DefaultSelectedChoiceReference *struct {
		Text string `xml:",chardata"`
	} `xml:"defaultSelectedChoiceReference"`
	DefaultValue  *Value       `xml:"defaultValue"`
	ExtensionName *string      `xml:"extensionName"`
	FieldText     *TextLiteral `xml:"fieldText"`
	FieldType     string       `xml:"fieldType"`
	Fields        []Field      `xml:"fields"`
	HelpText      *struct {
		Text string `xml:",chardata"`
	} `xml:"helpText"`
	InputParameters []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"inputParameters"`
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
	ObjectFieldReference *struct {
		Text string `xml:",chardata"`
	} `xml:"objectFieldReference"`
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
	Scale *struct {
		Text string `xml:",chardata"`
	} `xml:"scale"`
	StoreOutputAutomatically *struct {
		Text string `xml:",chardata"`
	} `xml:"storeOutputAutomatically"`
	StyleProperties *struct {
		VerticalAlignment *Value `xml:"verticalAlignment"`
		Width             *Value `xml:"width"`
	} `xml:"styleProperties"`
	ValidationRule *struct {
		ErrorMessage struct {
			Text string `xml:",chardata"`
		} `xml:"errorMessage"`
		FormulaExpression TextLiteral `xml:"formulaExpression"`
	} `xml:"validationRule"`
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
	Actions   []ScreenAction `xml:"actions"`
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
	StageReference *struct {
		ElementReference struct {
			Text string `xml:",chardata"`
		} `xml:"elementReference"`
	} `xml:"stageReference"`
	StyleSettings []struct {
		PropertyName  TextLiteral `xml:"propertyName"`
		PropertyValue Value       `xml:"propertyValue"`
		Scope         TextLiteral `xml:"scope"`
	} `xml:"styleSettings"`
	Triggers []ScreenTrigger `xml:"triggers"`
}

// ScreenAction is an action a screen runs while it is displayed, such as a
// subflow, without navigating away from the screen. The screen's triggers
// name the events that run it.
type ScreenAction struct {
	Name       ElementName `xml:"name"`
	ActionName struct {
		Text string `xml:",chardata"`
	} `xml:"actionName"`
	ActionType      string `xml:"actionType"`
	InputParameters []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"inputParameters"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	NameSegment *struct {
		Text string `xml:",chardata"`
	} `xml:"nameSegment"`
	VersionSegment *struct {
		Text string `xml:",chardata"`
	} `xml:"versionSegment"`
}

// ScreenTrigger runs screen actions when an event, such as the screen
// initializing or a field's value changing, fires on its event source.
type ScreenTrigger struct {
	EventName   TextLiteral `xml:"eventName"`
	EventSource TextLiteral `xml:"eventSource"`
	Handlers    []struct {
		ScreenActionName TextLiteral `xml:"screenActionName"`
	} `xml:"handlers"`
}

// Stage is a flow stage resource: one step of the flow's screen progress
// indicator, referenced by screens through stageReference.
type Stage struct {
	Description *TextLiteral `xml:"description"`
	Name        ElementName  `xml:"name"`
	IsActive    BooleanText  `xml:"isActive"`
	Label       struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	StageOrder struct {
		Text string `xml:",chardata"`
	} `xml:"stageOrder"`
}

// OrchestratedStage is a stage of a flow orchestration or flow approval
// process (processType Orchestrator or ApprovalWorkflow). A stage holds the
// steps that run while the orchestration is in it and connects to the next
// element once it completes.
type OrchestratedStage struct {
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
	Connector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	ExitActionInputParameters  []StageStepParameter       `xml:"exitActionInputParameters"`
	ExitActionName             *TextLiteral               `xml:"exitActionName"`
	ExitActionOutputParameters []StageStepOutputParameter `xml:"exitActionOutputParameters"`
	ExitActionType             *TextLiteral               `xml:"exitActionType"`
	ExitConditionLogic         *TextLiteral               `xml:"exitConditionLogic"`
	ExitConditions             []StageStepCondition       `xml:"exitConditions"`
	RunAsUser                  *BooleanText               `xml:"runAsUser"`
	StageSteps                 []StageStep                `xml:"stageSteps"`
}

// StageStep is one step of an OrchestratedStage: a background step runs an
// autolaunched flow, an approval step assigns an approval work item, and an
// interactive step assigns a screen flow to a user, group, or queue.
type StageStep struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Description           *TextLiteral           `xml:"description"`
	Name                  ElementName            `xml:"name"`
	ActionName            struct {
		Text string `xml:",chardata"`
	} `xml:"actionName"`
	ActionType string `xml:"actionType"`
	Assignees  []struct {
		Assignee     *Value `xml:"assignee"`
		AssigneeType struct {
			Text string `xml:",chardata"`
		} `xml:"assigneeType"`
	} `xml:"assignees"`
	CanAssigneeEdit             *BooleanText               `xml:"canAssigneeEdit"`
	DebugSimulateStep           *BooleanText               `xml:"debugSimulateStep"`
	EntryActionInputParameters  []StageStepParameter       `xml:"entryActionInputParameters"`
	EntryActionName             *TextLiteral               `xml:"entryActionName"`
	EntryActionOutputParameters []StageStepOutputParameter `xml:"entryActionOutputParameters"`
	EntryActionType             *TextLiteral               `xml:"entryActionType"`
	EntryConditionLogic         *TextLiteral               `xml:"entryConditionLogic"`
	EntryConditions             []StageStepCondition       `xml:"entryConditions"`
	ExitActionInputParameters   []StageStepParameter       `xml:"exitActionInputParameters"`
	ExitActionName              *TextLiteral               `xml:"exitActionName"`
	ExitActionOutputParameters  []StageStepOutputParameter `xml:"exitActionOutputParameters"`
	ExitActionType              *TextLiteral               `xml:"exitActionType"`
	ExitConditionLogic          *TextLiteral               `xml:"exitConditionLogic"`
	ExitConditions              []StageStepCondition       `xml:"exitConditions"`
	InputParameters             []StageStepParameter       `xml:"inputParameters"`
	Label                       struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	OutputConfigParams []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"outputConfigParams"`
	OutputParameters        []StageStepOutputParameter `xml:"outputParameters"`
	RequiresAsyncProcessing *BooleanText               `xml:"requiresAsyncProcessing"`
	RunAsUser               *BooleanText               `xml:"runAsUser"`
	ShouldLock              *BooleanText               `xml:"shouldLock"`
	StepSubtype             *TextLiteral               `xml:"stepSubtype"`
}

// StageStepParameter is an input parameter of a stage step or of a stage or
// step entry/exit evaluation flow.
type StageStepParameter struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Name                  struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Value *Value `xml:"value"`
}

// StageStepOutputParameter maps an output of a stage step's flow onto a flow
// resource.
type StageStepOutputParameter struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	AssignToReference     struct {
		Text string `xml:",chardata"`
	} `xml:"assignToReference"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

// StageStepCondition is an entry or exit condition of a stage or step. It
// carries a conditionType the decision Condition type does not.
type StageStepCondition struct {
	ConditionType      *TextLiteral `xml:"conditionType"`
	LeftValueReference string       `xml:"leftValueReference"`
	Operator           string       `xml:"operator"`
	RightValue         *Value       `xml:"rightValue"`
}

type Variable struct {
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Name      string `xml:"name"`
	ApexClass *struct {
		Text string `xml:",chardata"`
	} `xml:"apexClass"`
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
	Value *Value `xml:"value"`
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
	Offset           *TextLiteral `xml:"offset"`
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
	VersionString *TextLiteral `xml:"versionString"`
}

// Transform is a Transform element. It either maps a source collection onto a
// new collection, one output member per source member, or reduces a collection
// to a single value; IsCollection tells the two apart. The output is stored
// automatically under the element's own name.
type Transform struct {
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
	ElementSubtype *TextLiteral `xml:"elementSubtype"`
	ApexClass      *TextLiteral `xml:"apexClass"`
	Connector      *struct {
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
	// DataType and ObjectType describe what the transform produces: an
	// ObjectType is present only when DataType is SObject.
	DataType                 *TextLiteral     `xml:"dataType"`
	ObjectType               *TextLiteral     `xml:"objectType"`
	IsCollection             *BooleanText     `xml:"isCollection"`
	Scale                    *TextLiteral     `xml:"scale"`
	StoreOutputAutomatically *BooleanText     `xml:"storeOutputAutomatically"`
	TransformValues          []TransformValue `xml:"transformValues"`
}

// TransformValue groups the actions that produce one output value of a
// Transform.
type TransformValue struct {
	Name                  *TextLiteral           `xml:"name"`
	Label                 *TextLiteral           `xml:"label"`
	TransformValueActions []TransformValueAction `xml:"transformValueActions"`
}

// TransformValueAction produces one output of a Transform. A Map action names
// the expression it maps from in Value, and the field it lands on in
// OutputFieldApiName when the output is a record. An aggregate action (Count,
// Sum) carries no Value: it names the collection it reduces through an
// InputParameters entry called aggregationValues.
type TransformValueAction struct {
	// OutputFieldApiName is empty when the transform produces scalars rather
	// than records.
	OutputFieldApiName *TextLiteral `xml:"outputFieldApiName"`
	TransformType      *TextLiteral `xml:"transformType"`
	Value              *Value       `xml:"value"`
	InputParameters    []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"inputParameters"`
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
	AreMetricsLoggedToDataCloud *TextLiteral `xml:"areMetricsLoggedToDataCloud"`
	Assignments                 []Assignment `xml:"assignments"`
	Choices                     []struct {
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
	CollectionProcessors []CollectionProcessor `xml:"collectionProcessors"`
	CustomErrors         []CustomError         `xml:"customErrors"`
	CustomProperties     []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"customProperties"`
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
		OutputAssignments []struct {
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
		AssignNextValueToReference *struct {
			Text string `xml:",chardata"`
		} `xml:"assignNextValueToReference"`
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
	MigratedFromWorkflowRuleName *TextLiteral `xml:"migratedFromWorkflowRuleName"`
	ProcessMetadataValues        []struct {
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
	Screens               []Screen            `xml:"screens"`
	OrchestratedStages    []OrchestratedStage `xml:"orchestratedStages"`
	Stages                []Stage             `xml:"stages"`
	Start                 *Start              `xml:"start"`
	StartElementReference *struct {
		Text string `xml:",chardata"`
	} `xml:"startElementReference"`
	Status *struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	Steps    []Step `xml:"steps"`
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
		Description *TextLiteral `xml:"description"`
		Name        struct {
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
	Transforms []Transform `xml:"transforms"`
	Variables  []Variable  `xml:"variables"`
	Waits      []Wait      `xml:"waits"`
}

// Step is a step element from the legacy Cloud Flow Designer.
type Step struct {
	Description *TextLiteral `xml:"description"`
	Name        ElementName  `xml:"name"`
	Label       TextLiteral  `xml:"label"`
	LocationX   TextLiteral  `xml:"locationX"`
	LocationY   TextLiteral  `xml:"locationY"`
	Connectors  []struct {
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connectors"`
}

// Wait is a Wait element. Each wait event whose conditions are met is armed
// and pauses the interview until its event (e.g. a time-based alarm) fires;
// when no event arms, the default connector runs immediately.
type Wait struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Name                  struct {
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
	DefaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"defaultConnector"`
	DefaultConnectorLabel struct {
		Text string `xml:",chardata"`
	} `xml:"defaultConnectorLabel"`
	FaultConnector *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"faultConnector"`
	TimeZoneId *struct {
		Text string `xml:",chardata"`
	} `xml:"timeZoneId"`
	WaitEvents []WaitEvent `xml:"waitEvents"`
}

// WaitEvent is one event a Wait element can pause on. Conditions (combined
// per ConditionLogic) decide whether the event arms; EventType names the
// event kind (AlarmEvent, DateRefAlarmEvent, platform event API name);
// InputParameters carry the alarm's offset definition.
type WaitEvent struct {
	ProcessMetadataValues []ProcessMetadataValue `xml:"processMetadataValues"`
	Name                  struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	ConditionLogic struct {
		Text string `xml:",chardata"`
	} `xml:"conditionLogic"`
	Conditions []Condition `xml:"conditions"`
	Connector  *struct {
		IsGoTo *struct {
			Text string `xml:",chardata"`
		} `xml:"isGoTo"`
		TargetReference ElementName `xml:"targetReference"`
	} `xml:"connector"`
	EventType struct {
		Text string `xml:",chardata"`
	} `xml:"eventType"`
	InputParameters []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Value *Value `xml:"value"`
	} `xml:"inputParameters"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Offset struct {
		Text string `xml:",chardata"`
	} `xml:"offset"`
	OffsetUnit struct {
		Text string `xml:",chardata"`
	} `xml:"offsetUnit"`
}

// KeepEmptyElements marks flows as written with empty text elements open, as
// Salesforce writes them.
func (c *Flow) KeepEmptyElements() {}

// SelfClosedElements lists the flow elements of type Value, which Salesforce
// writes self-closed when they are empty.
func (c *Flow) SelfClosedElements() []string {
	return []string{"collectionElements", "defaultValue", "propertyValue", "rightValue", "value", "verticalAlignment", "width"}
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
