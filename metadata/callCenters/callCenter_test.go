package callCenters

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleCallCenter = `<?xml version="1.0" encoding="UTF-8"?>
<CallCenter xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <adapterUrl>https://example.com/adapter</adapterUrl>
    <customSettings>{&quot;reqUseApi&quot;:&quot;true&quot;}</customSettings>
    <displayName>Example Phone</displayName>
    <displayNameLabel>Display Name</displayNameLabel>
    <internalNameLabel>Internal Name</internalNameLabel>
    <sections>
        <items>
            <label>Description</label>
            <name>reqDescription</name>
            <value xsi:nil="true"/>
        </items>
        <items>
            <label>CTI Adapter URL</label>
            <name>reqAdapterUrl</name>
            <value>https://example.com/adapter</value>
        </items>
        <label>General Info</label>
        <name>reqGeneralInfo</name>
    </sections>
    <sections>
        <items>
            <label>Outside Prefix</label>
            <name>reqOutsidePrefix</name>
            <value>9</value>
        </items>
        <label>Dialing Options</label>
        <name>reqDialingOptions</name>
    </sections>
</CallCenter>
`

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ExamplePhone.callCenter")
	if err := os.WriteFile(path, []byte(sampleCallCenter), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	c, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := string(c.Type()); got != "CallCenter" {
		t.Errorf("Type = %q, want CallCenter", got)
	}
	if got := c.DisplayName.String(); got != "Example Phone" {
		t.Errorf("DisplayName = %q, want %q", got, "Example Phone")
	}
	if got := c.AdapterUrl.String(); got != "https://example.com/adapter" {
		t.Errorf("AdapterUrl = %q, want %q", got, "https://example.com/adapter")
	}
	if got := c.CustomSettings.String(); got != `{"reqUseApi":"true"}` {
		t.Errorf("CustomSettings = %q, want %q", got, `{"reqUseApi":"true"}`)
	}
	if got := c.InternalNameLabel.String(); got != "Internal Name" {
		t.Errorf("InternalNameLabel = %q, want %q", got, "Internal Name")
	}
	if got := len(c.Sections); got != 2 {
		t.Fatalf("len(Sections) = %d, want 2", got)
	}

	general := c.Sections[0]
	if got := general.Name.String(); got != "reqGeneralInfo" {
		t.Errorf("section[0].Name = %q, want reqGeneralInfo", got)
	}
	if got := len(general.Items); got != 2 {
		t.Fatalf("len(section[0].Items) = %d, want 2", got)
	}
	description := general.Items[0]
	if description.Value.Nil != "true" || description.Value.Text != "" {
		t.Errorf("reqDescription value = %+v, want nil", description.Value)
	}
	adapter := general.Items[1]
	if got := adapter.Name.String(); got != "reqAdapterUrl" {
		t.Errorf("item[1].Name = %q, want reqAdapterUrl", got)
	}
	if adapter.Value.Nil != "" || adapter.Value.Text != "https://example.com/adapter" {
		t.Errorf("reqAdapterUrl value = %+v, want https://example.com/adapter", adapter.Value)
	}

	if got := c.Sections[1].Items[0].Value.Text; got != "9" {
		t.Errorf("reqOutsidePrefix value = %q, want 9", got)
	}
}
