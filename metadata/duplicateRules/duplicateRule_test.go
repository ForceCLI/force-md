package duplicateRule

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

func assertRoundTrip(t *testing.T, content string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Account.Test.duplicateRule")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	d, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the duplicate rule:\n%s", out)
	}
}

// Nil elements keep the xsi namespace prefix, an empty filter value stays
// open, and operationsOnInsert is written before operationsOnUpdate.
func TestDuplicateRuleRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<DuplicateRule xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <actionOnInsert>Block</actionOnInsert>
    <actionOnUpdate>Allow</actionOnUpdate>
    <alertText xsi:nil="true"/>
    <description xsi:nil="true"/>
    <duplicateRuleFilter>
        <booleanFilter xsi:nil="true"/>
        <duplicateRuleFilterItems>
            <field>MRN__c</field>
            <operation>notEqual</operation>
            <value></value>
            <sortOrder>1</sortOrder>
            <table>Account</table>
        </duplicateRuleFilterItems>
    </duplicateRuleFilter>
    <duplicateRuleMatchRules>
        <matchRuleSObjectType>Account</matchRuleSObjectType>
        <matchingRule>MRN_Matching_Rule</matchingRule>
        <objectMapping xsi:nil="true"/>
    </duplicateRuleMatchRules>
    <isActive>true</isActive>
    <masterLabel>MRN</masterLabel>
    <operationsOnInsert>Alert</operationsOnInsert>
    <operationsOnInsert>Report</operationsOnInsert>
    <operationsOnUpdate>Alert</operationsOnUpdate>
    <securityOption>BypassSharingRules</securityOption>
    <sortOrder>2</sortOrder>
</DuplicateRule>
`)
}

// A rule matching another object maps fields between the objects.
func TestDuplicateRuleObjectMappingRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<DuplicateRule xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <actionOnInsert>Allow</actionOnInsert>
    <actionOnUpdate>Allow</actionOnUpdate>
    <alertText>Use one of these records?</alertText>
    <description>Identify leads with duplicate contacts.</description>
    <duplicateRuleFilter xsi:nil="true"/>
    <duplicateRuleMatchRules>
        <matchRuleSObjectType>Contact</matchRuleSObjectType>
        <matchingRule>Standard_Contact_Match_Rule_v1_1</matchingRule>
        <objectMapping>
            <inputObject>Lead</inputObject>
            <mappingFields>
                <inputField>Email</inputField>
                <outputField>Email</outputField>
            </mappingFields>
            <outputObject>Contact</outputObject>
        </objectMapping>
    </duplicateRuleMatchRules>
    <isActive>true</isActive>
    <masterLabel>Lead to Contact</masterLabel>
    <securityOption>EnforceSharingRules</securityOption>
    <sortOrder>1</sortOrder>
</DuplicateRule>
`)
}
