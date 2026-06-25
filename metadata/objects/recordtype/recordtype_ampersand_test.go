package recordtype

import (
	"os"
	"path/filepath"
	"testing"
)

// A record type retrieved from a real org can carry picklist values containing
// bare & characters. Open must parse the whole record type rather than failing
// on the technically-invalid XML.
func TestOpenRecordTypeWithBareAmpersand(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "RnD.recordType-meta.xml")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<RecordType xmlns="http://soap.sforce.com/2006/04/metadata">
    <fullName>RnD</fullName>
    <active>true</active>
    <label>R&D</label>
    <picklistValues>
        <picklist>Category__c</picklist>
        <values>
            <fullName>a & b</fullName>
            <default>false</default>
        </values>
        <values>
            <fullName>R&D</fullName>
            <default>true</default>
        </values>
    </picklistValues>
</RecordType>`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	rt, err := Open(path)
	if err != nil {
		t.Fatalf("Open returned error for bare & metadata: %v", err)
	}
	if rt.Label.Text != "R&D" {
		t.Errorf("label = %q, want %q", rt.Label.Text, "R&D")
	}
	if len(rt.PicklistValues) != 1 {
		t.Fatalf("expected 1 picklist, got %d", len(rt.PicklistValues))
	}
	values := rt.PicklistValues[0].Values
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
