package field

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

func assertFieldRoundTrip(t *testing.T, name, contents string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), name+".field-meta.xml")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	f, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(f)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != contents {
		t.Errorf("round trip changed the field:\n%s", out)
	}
}

// A picklist value's color is written after its fullName, and its
// description is kept.
func TestValueSetValueRoundTrip(t *testing.T) {
	assertFieldRoundTrip(t, "Source__c", `<?xml version="1.0" encoding="UTF-8"?>
<CustomField xmlns="http://soap.sforce.com/2006/04/metadata">
    <fullName>Source__c</fullName>
    <label>Source</label>
    <type>Picklist</type>
    <valueSet>
        <valueSetDefinition>
            <sorted>false</sorted>
            <value>
                <fullName>Advertising</fullName>
                <color>#78C953</color>
                <default>false</default>
                <description>Paid placements</description>
                <label>Advertising</label>
            </value>
        </valueSetDefinition>
    </valueSet>
</CustomField>
`)
}

// A master-detail field's elements are written in Salesforce's alphabetical
// order, with relationshipOrder before trackFeedHistory.
func TestMasterDetailFieldRoundTrip(t *testing.T) {
	assertFieldRoundTrip(t, "Account__c", `<?xml version="1.0" encoding="UTF-8"?>
<CustomField xmlns="http://soap.sforce.com/2006/04/metadata">
    <fullName>Account__c</fullName>
    <description>The inquiring account</description>
    <label>Account</label>
    <referenceTo>Account</referenceTo>
    <relationshipLabel>Inquiries</relationshipLabel>
    <relationshipName>Inquiries</relationshipName>
    <relationshipOrder>0</relationshipOrder>
    <reparentableMasterDetail>false</reparentableMasterDetail>
    <trackFeedHistory>false</trackFeedHistory>
    <trackHistory>false</trackHistory>
    <trackTrending>false</trackTrending>
    <type>MasterDetail</type>
    <writeRequiresMasterRead>false</writeRequiresMasterRead>
</CustomField>
`)
}
