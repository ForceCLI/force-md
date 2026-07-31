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
