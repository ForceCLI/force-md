package flow

import (
	"os"
	"path/filepath"
	"testing"
)

func openFlow(t *testing.T, contents string) *Flow {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Test_Flow.flow-meta.xml")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	f, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	return f
}

const customErrorFlow = `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <customErrors>
        <description>Blocks the save outright</description>
        <name>Record_Error</name>
        <label>Record Error</label>
        <locationX>50</locationX>
        <locationY>300</locationY>
        <customErrorMessages>
            <errorMessage>The pending request must be resolved first.</errorMessage>
            <isFieldError>false</isFieldError>
        </customErrorMessages>
    </customErrors>
    <customErrors>
        <name>Field_Error</name>
        <label>Field Error</label>
        <locationX>200</locationX>
        <locationY>300</locationY>
        <customErrorMessages>
            <errorMessage>Pick a rating.</errorMessage>
            <fieldSelection>Rating</fieldSelection>
            <isFieldError>true</isFieldError>
        </customErrorMessages>
        <customErrorMessages>
            <errorMessage>And a site.</errorMessage>
            <fieldSelection>Site</fieldSelection>
            <isFieldError>true</isFieldError>
        </customErrorMessages>
    </customErrors>
    <label>Test Flow</label>
    <processType>AutoLaunchedFlow</processType>
    <status>Active</status>
</Flow>
`

func TestCustomErrors(t *testing.T) {
	f := openFlow(t, customErrorFlow)

	if got := len(f.CustomErrors); got != 2 {
		t.Fatalf("CustomErrors = %d, want 2", got)
	}

	recordError := f.CustomErrors[0]
	if got := recordError.Name.Text; got != "Record_Error" {
		t.Errorf("Name = %q, want Record_Error", got)
	}
	if got := recordError.Label.Text; got != "Record Error" {
		t.Errorf("Label = %q, want %q", got, "Record Error")
	}
	if got := recordError.Description.String(); got != "Blocks the save outright" {
		t.Errorf("Description = %q, want %q", got, "Blocks the save outright")
	}
	if got := len(recordError.CustomErrorMessages); got != 1 {
		t.Fatalf("Record_Error messages = %d, want 1", got)
	}
	if got := recordError.CustomErrorMessages[0].ErrorMessage.Text; got != "The pending request must be resolved first." {
		t.Errorf("ErrorMessage = %q, want %q", got, "The pending request must be resolved first.")
	}
	if recordError.CustomErrorMessages[0].IsFieldError.ToBool() {
		t.Error("IsFieldError = true, want false")
	}
	if recordError.CustomErrorMessages[0].FieldSelection != nil {
		t.Errorf("FieldSelection = %v, want nil for a record-level message", recordError.CustomErrorMessages[0].FieldSelection)
	}

	// Salesforce allows an element to report more than one message, so they
	// have to survive as separate entries rather than collapsing onto one.
	fieldError := f.CustomErrors[1]
	if got := len(fieldError.CustomErrorMessages); got != 2 {
		t.Fatalf("Field_Error messages = %d, want 2", got)
	}
	if fieldError.Description != nil {
		t.Errorf("Description = %v, want nil", fieldError.Description)
	}
	for i, want := range []struct {
		message string
		field   string
	}{
		{"Pick a rating.", "Rating"},
		{"And a site.", "Site"},
	} {
		message := fieldError.CustomErrorMessages[i]
		if got := message.ErrorMessage.Text; got != want.message {
			t.Errorf("message %d ErrorMessage = %q, want %q", i, got, want.message)
		}
		if !message.IsFieldError.ToBool() {
			t.Errorf("message %d IsFieldError = false, want true", i)
		}
		if message.FieldSelection == nil {
			t.Fatalf("message %d FieldSelection = nil, want %q", i, want.field)
		}
		if got := message.FieldSelection.Text; got != want.field {
			t.Errorf("message %d FieldSelection = %q, want %q", i, got, want.field)
		}
	}
}

func TestCustomErrorsAbsent(t *testing.T) {
	f := openFlow(t, `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Test Flow</label>
    <processType>AutoLaunchedFlow</processType>
    <status>Active</status>
</Flow>
`)

	if got := len(f.CustomErrors); got != 0 {
		t.Errorf("CustomErrors = %d, want 0", got)
	}
}

const collectionProcessorFlow = `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <collectionProcessors>
        <description>Keeps the open ones</description>
        <name>Open_Jobs</name>
        <elementSubtype>FilterCollectionProcessor</elementSubtype>
        <label>Open Jobs</label>
        <locationX>1106</locationX>
        <locationY>395</locationY>
        <assignNextValueToReference>currentItem_Open_Jobs</assignNextValueToReference>
        <collectionProcessorType>FilterCollectionProcessor</collectionProcessorType>
        <collectionReference>Get_Jobs</collectionReference>
        <conditionLogic>and</conditionLogic>
        <conditions>
            <leftValueReference>currentItem_Open_Jobs.Status__c</leftValueReference>
            <operator>EqualTo</operator>
            <rightValue>
                <stringValue>Open</stringValue>
            </rightValue>
        </conditions>
        <conditions>
            <leftValueReference>currentItem_Open_Jobs.Is_Paid__c</leftValueReference>
            <operator>EqualTo</operator>
            <rightValue>
                <booleanValue>false</booleanValue>
            </rightValue>
        </conditions>
        <connector>
            <targetReference>Next_Element</targetReference>
        </connector>
    </collectionProcessors>
    <collectionProcessors>
        <name>Sort_Jobs</name>
        <elementSubtype>SortCollectionProcessor</elementSubtype>
        <label>Sort Jobs</label>
        <locationX>842</locationX>
        <locationY>1250</locationY>
        <collectionProcessorType>SortCollectionProcessor</collectionProcessorType>
        <collectionReference>Get_Jobs</collectionReference>
        <limit>5</limit>
        <sortOptions>
            <doesPutEmptyStringAndNullFirst>false</doesPutEmptyStringAndNullFirst>
            <sortField>Is_GC__c</sortField>
            <sortOrder>Desc</sortOrder>
        </sortOptions>
        <sortOptions>
            <doesPutEmptyStringAndNullFirst>true</doesPutEmptyStringAndNullFirst>
            <sortField>Name</sortField>
            <sortOrder>Asc</sortOrder>
        </sortOptions>
    </collectionProcessors>
    <label>Test Flow</label>
    <processType>AutoLaunchedFlow</processType>
    <status>Active</status>
</Flow>
`

func TestCollectionProcessors(t *testing.T) {
	f := openFlow(t, collectionProcessorFlow)

	if got := len(f.CollectionProcessors); got != 2 {
		t.Fatalf("CollectionProcessors = %d, want 2", got)
	}

	filter := f.CollectionProcessors[0]
	if got := filter.Name.Text; got != "Open_Jobs" {
		t.Errorf("Name = %q, want Open_Jobs", got)
	}
	if got := filter.CollectionProcessorType.Text; got != "FilterCollectionProcessor" {
		t.Errorf("CollectionProcessorType = %q, want FilterCollectionProcessor", got)
	}
	if got := filter.CollectionReference.Text; got != "Get_Jobs" {
		t.Errorf("CollectionReference = %q, want Get_Jobs", got)
	}
	if got := filter.AssignNextValueToReference.Text; got != "currentItem_Open_Jobs" {
		t.Errorf("AssignNextValueToReference = %q, want currentItem_Open_Jobs", got)
	}
	if got := filter.ConditionLogic.Text; got != "and" {
		t.Errorf("ConditionLogic = %q, want and", got)
	}
	if got := filter.Description.String(); got != "Keeps the open ones" {
		t.Errorf("Description = %q, want %q", got, "Keeps the open ones")
	}
	if got := string(filter.Connector.TargetReference); got != "Next_Element" {
		t.Errorf("Connector.TargetReference = %q, want Next_Element", got)
	}

	// A filter routinely applies more than one condition, so they have to
	// survive as separate entries rather than collapsing onto one.
	if got := len(filter.Conditions); got != 2 {
		t.Fatalf("Conditions = %d, want 2", got)
	}
	if got := filter.Conditions[0].LeftValueReference; got != "currentItem_Open_Jobs.Status__c" {
		t.Errorf("condition 0 LeftValueReference = %q, want currentItem_Open_Jobs.Status__c", got)
	}
	if got := filter.Conditions[0].Operator; got != "EqualTo" {
		t.Errorf("condition 0 Operator = %q, want EqualTo", got)
	}
	if got := filter.Conditions[0].RightValue.String(); got != "Open" {
		t.Errorf("condition 0 RightValue = %q, want Open", got)
	}
	if got := filter.Conditions[1].LeftValueReference; got != "currentItem_Open_Jobs.Is_Paid__c" {
		t.Errorf("condition 1 LeftValueReference = %q, want currentItem_Open_Jobs.Is_Paid__c", got)
	}
	if got := filter.Conditions[1].RightValue.BooleanValue.ToBool(); got {
		t.Error("condition 1 RightValue = true, want false")
	}
	if filter.Limit != nil {
		t.Errorf("Limit = %v, want nil for a filter", filter.Limit)
	}
	if got := len(filter.SortOptions); got != 0 {
		t.Errorf("SortOptions = %d, want 0 for a filter", got)
	}

	sort := f.CollectionProcessors[1]
	if got := sort.CollectionProcessorType.Text; got != "SortCollectionProcessor" {
		t.Errorf("CollectionProcessorType = %q, want SortCollectionProcessor", got)
	}
	if sort.Limit == nil {
		t.Fatal("Limit = nil, want 5")
	}
	if got := sort.Limit.Text; got != "5" {
		t.Errorf("Limit = %q, want 5", got)
	}
	if sort.Description != nil {
		t.Errorf("Description = %v, want nil", sort.Description)
	}
	if got := len(sort.Conditions); got != 0 {
		t.Errorf("Conditions = %d, want 0 for a sort", got)
	}

	// Sort levels apply in the order they appear, so order and per-level
	// options both matter.
	if got := len(sort.SortOptions); got != 2 {
		t.Fatalf("SortOptions = %d, want 2", got)
	}
	for i, want := range []struct {
		field     string
		order     string
		nullFirst bool
	}{
		{"Is_GC__c", "Desc", false},
		{"Name", "Asc", true},
	} {
		option := sort.SortOptions[i]
		if got := option.SortField.Text; got != want.field {
			t.Errorf("sort option %d SortField = %q, want %q", i, got, want.field)
		}
		if got := option.SortOrder.Text; got != want.order {
			t.Errorf("sort option %d SortOrder = %q, want %q", i, got, want.order)
		}
		if got := option.DoesPutEmptyStringAndNullFirst.ToBool(); got != want.nullFirst {
			t.Errorf("sort option %d DoesPutEmptyStringAndNullFirst = %v, want %v", i, got, want.nullFirst)
		}
	}
}

const waitFlow = `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <waits>
        <name>myWait_myRule_1</name>
        <label>myWait_myRule_1</label>
        <locationX>0</locationX>
        <locationY>0</locationY>
        <defaultConnector>
            <targetReference>After_Wait</targetReference>
        </defaultConnector>
        <defaultConnectorLabel>defaultLabel</defaultConnectorLabel>
        <faultConnector>
            <targetReference>On_Fault</targetReference>
        </faultConnector>
        <timeZoneId>America/Los_Angeles</timeZoneId>
        <waitEvents>
            <name>myWaitEvent_event_0</name>
            <conditionLogic>and</conditionLogic>
            <conditions>
                <leftValueReference>postActionExecutionVariable</leftValueReference>
                <operator>EqualTo</operator>
                <rightValue>
                    <booleanValue>false</booleanValue>
                </rightValue>
            </conditions>
            <connector>
                <targetReference>Post_Wait_Decision</targetReference>
            </connector>
            <eventType>DateRefAlarmEvent</eventType>
            <inputParameters>
                <name>TimeTableColumnEnumOrId</name>
                <value>
                    <stringValue>Showing__c</stringValue>
                </value>
            </inputParameters>
            <inputParameters>
                <name>TimeOffset</name>
                <value>
                    <numberValue>24.0</numberValue>
                </value>
            </inputParameters>
            <label>myWaitEvent_event_0</label>
        </waitEvents>
        <waitEvents>
            <name>myWaitEvent_event_1</name>
            <conditionLogic>or</conditionLogic>
            <connector>
                <targetReference>Other_Path</targetReference>
            </connector>
            <eventType>AlarmEvent</eventType>
            <offset>2</offset>
            <offsetUnit>Hours</offsetUnit>
            <label>myWaitEvent_event_1</label>
        </waitEvents>
    </waits>
</Flow>`

func TestWaits(t *testing.T) {
	f := openFlow(t, waitFlow)

	if len(f.Waits) != 1 {
		t.Fatalf("Waits = %d, want 1", len(f.Waits))
	}
	w := f.Waits[0]
	if w.Name.Text != "myWait_myRule_1" {
		t.Errorf("Name = %q, want %q", w.Name.Text, "myWait_myRule_1")
	}
	if w.DefaultConnector == nil || string(w.DefaultConnector.TargetReference) != "After_Wait" {
		t.Errorf("DefaultConnector = %+v, want target After_Wait", w.DefaultConnector)
	}
	if w.FaultConnector == nil || string(w.FaultConnector.TargetReference) != "On_Fault" {
		t.Errorf("FaultConnector = %+v, want target On_Fault", w.FaultConnector)
	}
	if w.TimeZoneId == nil || w.TimeZoneId.Text != "America/Los_Angeles" {
		t.Errorf("TimeZoneId = %+v, want America/Los_Angeles", w.TimeZoneId)
	}

	if len(w.WaitEvents) != 2 {
		t.Fatalf("WaitEvents = %d, want 2", len(w.WaitEvents))
	}

	first := w.WaitEvents[0]
	if first.Name.Text != "myWaitEvent_event_0" {
		t.Errorf("event 0 Name = %q", first.Name.Text)
	}
	if first.ConditionLogic.Text != "and" {
		t.Errorf("event 0 ConditionLogic = %q, want and", first.ConditionLogic.Text)
	}
	if len(first.Conditions) != 1 {
		t.Fatalf("event 0 Conditions = %d, want 1", len(first.Conditions))
	}
	cond := first.Conditions[0]
	if cond.LeftValueReference != "postActionExecutionVariable" || cond.Operator != "EqualTo" {
		t.Errorf("event 0 condition = %+v", cond)
	}
	if cond.RightValue == nil || cond.RightValue.BooleanValue == nil || cond.RightValue.BooleanValue.ToBool() {
		t.Errorf("event 0 condition RightValue = %+v, want booleanValue false", cond.RightValue)
	}
	if first.Connector == nil || string(first.Connector.TargetReference) != "Post_Wait_Decision" {
		t.Errorf("event 0 Connector = %+v, want target Post_Wait_Decision", first.Connector)
	}
	if first.EventType.Text != "DateRefAlarmEvent" {
		t.Errorf("event 0 EventType = %q, want DateRefAlarmEvent", first.EventType.Text)
	}
	if len(first.InputParameters) != 2 {
		t.Fatalf("event 0 InputParameters = %d, want 2", len(first.InputParameters))
	}
	if first.InputParameters[0].Name.Text != "TimeTableColumnEnumOrId" ||
		first.InputParameters[0].Value == nil ||
		first.InputParameters[0].Value.StringValue == nil ||
		first.InputParameters[0].Value.StringValue.Text != "Showing__c" {
		t.Errorf("event 0 input parameter 0 = %+v", first.InputParameters[0])
	}
	if first.InputParameters[1].Name.Text != "TimeOffset" ||
		first.InputParameters[1].Value == nil ||
		first.InputParameters[1].Value.NumberValue == nil ||
		first.InputParameters[1].Value.NumberValue.Text != "24.0" {
		t.Errorf("event 0 input parameter 1 = %+v", first.InputParameters[1])
	}

	second := w.WaitEvents[1]
	if second.ConditionLogic.Text != "or" {
		t.Errorf("event 1 ConditionLogic = %q, want or", second.ConditionLogic.Text)
	}
	if len(second.Conditions) != 0 {
		t.Errorf("event 1 Conditions = %d, want 0", len(second.Conditions))
	}
	if second.EventType.Text != "AlarmEvent" {
		t.Errorf("event 1 EventType = %q, want AlarmEvent", second.EventType.Text)
	}
	if second.Offset.Text != "2" || second.OffsetUnit.Text != "Hours" {
		t.Errorf("event 1 offset = %q %q, want 2 Hours", second.Offset.Text, second.OffsetUnit.Text)
	}
}

func TestWaitsAbsent(t *testing.T) {
	f := openFlow(t, `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata"></Flow>`)
	if len(f.Waits) != 0 {
		t.Errorf("Waits = %d, want 0", len(f.Waits))
	}
}
