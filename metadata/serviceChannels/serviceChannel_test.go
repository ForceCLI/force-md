package serviceChannels

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleServiceChannel = `<?xml version="1.0" encoding="UTF-8"?>
<ServiceChannel xmlns="http://soap.sforce.com/2006/04/metadata">
    <capacityModel>STATUS_BASED</capacityModel>
    <doesCheckCapOnOwnerChange>false</doesCheckCapOnOwnerChange>
    <doesCheckCapOnStatusChange>true</doesCheckCapOnStatusChange>
    <doesMinimizeWidgetOnAccept>false</doesMinimizeWidgetOnAccept>
    <hasAutoAcceptEnabled>false</hasAutoAcceptEnabled>
    <isInterruptible>false</isInterruptible>
    <label>Leads</label>
    <relatedEntityType>Lead</relatedEntityType>
    <secondaryRoutingPriorityField>CreatedDate</secondaryRoutingPriorityField>
    <serviceChannelFieldPriorities>
        <priority>1</priority>
        <value>Hot</value>
    </serviceChannelFieldPriorities>
    <serviceChannelStatusFieldMappings>
        <type>COMPLETED</type>
        <value>Qualified</value>
    </serviceChannelStatusFieldMappings>
    <serviceChannelStatusFieldMappings>
        <type>IN_PROGRESS</type>
        <value>Working &amp; Contacted</value>
    </serviceChannelStatusFieldMappings>
    <statusField>Status</statusField>
</ServiceChannel>
`

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Leads.serviceChannel")
	if err := os.WriteFile(path, []byte(sampleServiceChannel), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	c, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := string(c.Type()); got != "ServiceChannel" {
		t.Errorf("Type = %q, want ServiceChannel", got)
	}
	if got := c.Label.String(); got != "Leads" {
		t.Errorf("Label = %q, want Leads", got)
	}
	if got := c.RelatedEntityType.String(); got != "Lead" {
		t.Errorf("RelatedEntityType = %q, want Lead", got)
	}
	if got := c.CapacityModel.String(); got != "STATUS_BASED" {
		t.Errorf("CapacityModel = %q, want STATUS_BASED", got)
	}
	if !c.DoesCheckCapOnStatusChange.ToBool() {
		t.Error("DoesCheckCapOnStatusChange = false, want true")
	}
	if c.DoesMinimizeWidgetOnAccept.ToBool() {
		t.Error("DoesMinimizeWidgetOnAccept = true, want false")
	}
	if c.AfterConvoWorkMaxTime != nil {
		t.Errorf("AfterConvoWorkMaxTime = %+v, want nil", c.AfterConvoWorkMaxTime)
	}
	if got := c.SecondaryRoutingPriorityField.String(); got != "CreatedDate" {
		t.Errorf("SecondaryRoutingPriorityField = %q, want CreatedDate", got)
	}
	if got := len(c.ServiceChannelFieldPriorities); got != 1 {
		t.Fatalf("len(ServiceChannelFieldPriorities) = %d, want 1", got)
	}
	if got := c.ServiceChannelFieldPriorities[0].Priority.Text; got != "1" {
		t.Errorf("field priority = %q, want 1", got)
	}
	if got := len(c.ServiceChannelStatusFieldMappings); got != 2 {
		t.Fatalf("len(ServiceChannelStatusFieldMappings) = %d, want 2", got)
	}
	mapping := c.ServiceChannelStatusFieldMappings[1]
	if got := mapping.Type.String(); got != "IN_PROGRESS" {
		t.Errorf("mapping[1].Type = %q, want IN_PROGRESS", got)
	}
	if got := mapping.Value.String(); got != "Working & Contacted" {
		t.Errorf("mapping[1].Value = %q, want %q", got, "Working & Contacted")
	}
	if got := c.StatusField.String(); got != "Status" {
		t.Errorf("StatusField = %q, want Status", got)
	}
}
