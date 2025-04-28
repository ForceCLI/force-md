package recordtype

import (
	"encoding/xml"
	"testing"
)

// TestRecordTypeMetadataTidy verifies that the Tidy method sorts picklists by name
// and sorts each picklist's values by fullName.
func TestRecordTypeMetadataTidy(t *testing.T) {
	xmlData := []byte(`
<RecordType xmlns="http://soap.sforce.com/2006/04/metadata">
    <picklistValues>
        <picklist>FieldB__c</picklist>
        <values>
            <fullName>Val2</fullName>
        </values>
        <values>
            <fullName>Val1</fullName>
        </values>
    </picklistValues>
    <picklistValues>
        <picklist>FieldA__c</picklist>
        <values>
            <fullName>AVal</fullName>
        </values>
    </picklistValues>
</RecordType>`)

	var rt RecordTypeMetadata
	if err := xml.Unmarshal(xmlData, &rt); err != nil {
		t.Fatalf("failed to unmarshal XML: %v", err)
	}
	rt.Tidy()
	// Expect two picklists
	if len(rt.PicklistValues) != 2 {
		t.Fatalf("expected 2 picklists, got %d", len(rt.PicklistValues))
	}
	// Picklists should be sorted: FieldA__c, FieldB__c
	if rt.PicklistValues[0].Picklist != "FieldA__c" {
		t.Errorf("first picklist: want FieldA__c, got %s", rt.PicklistValues[0].Picklist)
	}
	if rt.PicklistValues[1].Picklist != "FieldB__c" {
		t.Errorf("second picklist: want FieldB__c, got %s", rt.PicklistValues[1].Picklist)
	}
	// Values for FieldB__c should be sorted: Val1, Val2
	vals := rt.PicklistValues[1].Values
	if len(vals) != 2 {
		t.Fatalf("expected 2 values for FieldB__c, got %d", len(vals))
	}
	if vals[0].FullName != "Val1" || vals[1].FullName != "Val2" {
		t.Errorf("values order for FieldB__c: want [Val1,Val2], got [%s,%s]", vals[0].FullName, vals[1].FullName)
	}
}
