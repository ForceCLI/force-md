package field

import (
	"os"
	"path/filepath"
	"testing"
)

// Salesforce can retrieve picklist metadata containing bare & characters that
// are not valid XML. A real org still deploys such files, so Open must parse
// them, treating the bare ampersand as a literal character.
func TestOpenPicklistWithBareAmpersand(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Category__c.field-meta.xml")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomField xmlns="http://soap.sforce.com/2006/04/metadata">
    <fullName>Category__c</fullName>
    <label>Category</label>
    <type>Picklist</type>
    <valueSet>
        <valueSetDefinition>
            <sorted>false</sorted>
            <value>
                <fullName>a & b</fullName>
                <default>false</default>
                <label>a & b</label>
            </value>
            <value>
                <fullName>R&D</fullName>
                <default>false</default>
                <label>R&D</label>
            </value>
        </valueSetDefinition>
    </valueSet>
</CustomField>`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	cf, err := Open(path)
	if err != nil {
		t.Fatalf("Open returned error for bare & metadata: %v", err)
	}
	if cf.ValueSet == nil || cf.ValueSet.ValueSetDefinition == nil {
		t.Fatal("valueSetDefinition was not parsed")
	}
	values := cf.ValueSet.ValueSetDefinition.Value
	if len(values) != 2 {
		t.Fatalf("expected 2 picklist values, got %d", len(values))
	}
	if values[0].FullName != "a & b" {
		t.Errorf("value[0].FullName = %q, want %q", values[0].FullName, "a & b")
	}
	if values[1].FullName != "R&D" {
		t.Errorf("value[1].FullName = %q, want %q", values[1].FullName, "R&D")
	}
}
