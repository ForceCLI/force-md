package territory2model

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenParsesTerritory2Model(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Geographic.territory2Model")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Model xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>Geographic</name>
    <description>Top level model</description>
</Territory2Model>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if m.Name.Text != "Geographic" {
		t.Errorf("name = %q, want Geographic", m.Name.Text)
	}
	if m.Description == nil || m.Description.Text != "Top level model" {
		t.Errorf("description = %v, want Top level model", m.Description)
	}
	if m.Type() != NAME {
		t.Errorf("Type() = %q, want %q", m.Type(), NAME)
	}
}

func TestOpenParsesTerritory2ModelCustomFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "FY13.territory2Model")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Model xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <customFields>
        <name>AutoNumber__c</name>
        <value xsi:type="xsd:string">M# 000001</value>
    </customFields>
    <name>FY13</name>
</Territory2Model>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(m.CustomFields) != 1 {
		t.Fatalf("customFields length = %d, want 1", len(m.CustomFields))
	}
	field := m.CustomFields[0]
	if field.Name.Text != "AutoNumber__c" || field.Value.Type != "xsd:string" || field.Value.Text != "M# 000001" {
		t.Errorf("customFields[0] = %+v, want AutoNumber__c xsd:string M# 000001", field)
	}
}

func TestOpenParsesTerritory2ModelRecordAccess(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "FY13.territory2Model")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Model xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>FY13</name>
    <recordVisibility>Private</recordVisibility>
    <userRecordAccess>Read</userRecordAccess>
</Territory2Model>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if m.RecordVisibility == nil || m.RecordVisibility.Text != "Private" {
		t.Errorf("recordVisibility = %v, want Private", m.RecordVisibility)
	}
	if m.UserRecordAccess == nil || m.UserRecordAccess.Text != "Read" {
		t.Errorf("userRecordAccess = %v, want Read", m.UserRecordAccess)
	}
}
