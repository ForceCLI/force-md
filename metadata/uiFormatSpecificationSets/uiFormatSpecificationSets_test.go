package uiFormatSpecificationSets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// A visibility rule without a boolean filter is written back without an
// empty booleanFilter element.
func TestVisibilityRuleWithoutBooleanFilterRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<UiFormatSpecificationSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <field>Assistance_Needed__c</field>
    <formatType>Field</formatType>
    <masterLabel>Assistance Needed</masterLabel>
    <sobjectType>Account</sobjectType>
    <uiFormatSpecifications>
        <formatProperties>{&quot;icon&quot;:&quot;utility:warning&quot;}</formatProperties>
        <formatType>Icon</formatType>
        <order>1</order>
        <visibilityRule>
            <criteria>
                <leftValue>{!Record.Assistance_Needed__c}</leftValue>
                <operator>EQUAL</operator>
                <rightValue>Yes</rightValue>
            </criteria>
        </visibilityRule>
    </uiFormatSpecifications>
</UiFormatSpecificationSet>
`
	path := filepath.Join(t.TempDir(), "Test.uiFormatSpecificationSet")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	u, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(u)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the format specification set:\n%s", out)
	}
}
