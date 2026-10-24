package workflow

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// Outbound messages, a rule's failed migration tool version, and a task with
// a description and no assignedTo or offsetFromField are written back
// unchanged.
func TestWorkflowRoundTrip(t *testing.T) {
	contents := `<?xml version="1.0" encoding="UTF-8"?>
<Workflow xmlns="http://soap.sforce.com/2006/04/metadata">
    <outboundMessages>
        <fullName>Notify_Service</fullName>
        <apiVersion>59.0</apiVersion>
        <endpointUrl>https://example.com/hook</endpointUrl>
        <fields>Id</fields>
        <fields>Name</fields>
        <includeSessionId>false</includeSessionId>
        <integrationUser>integration@example.com</integrationUser>
        <name>Notify Service</name>
        <protected>false</protected>
        <useDeadLetterQueue>false</useDeadLetterQueue>
    </outboundMessages>
    <rules>
        <fullName>Lead Created</fullName>
        <actions>
            <name>Notify_Service</name>
            <type>OutboundMessage</type>
        </actions>
        <active>false</active>
        <failedMigrationToolVersion>246.6.0</failedMigrationToolVersion>
        <formula>true</formula>
        <triggerType>onCreateOnly</triggerType>
    </rules>
    <tasks>
        <fullName>Call_Lead</fullName>
        <assignedToType>owner</assignedToType>
        <description>Call the lead within a day.</description>
        <dueDateOffset>1</dueDateOffset>
        <notifyAssignee>false</notifyAssignee>
        <priority>Normal</priority>
        <protected>false</protected>
        <status>Not Started</status>
        <subject>Call Lead</subject>
    </tasks>
</Workflow>
`
	path := filepath.Join(t.TempDir(), "Lead.workflow")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	w, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(w)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != contents {
		t.Errorf("round trip changed the workflow:\n%s", out)
	}
}
