package matchingrules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

func assertRoundTrip(t *testing.T, content string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Account.matchingRule")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(m)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the matching rules:\n%q", out)
	}
}

// A matching rule without a description is written back without an empty
// description element.
func TestMatchingRuleWithoutDescriptionRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<MatchingRules xmlns="http://soap.sforce.com/2006/04/metadata">
    <matchingRules>
        <fullName>MRN_Matching_Rule</fullName>
        <label>MRN Matching Rule</label>
        <matchingRuleItems>
            <blankValueBehavior>NullNotAllowed</blankValueBehavior>
            <fieldName>MRN__c</fieldName>
            <matchingMethod>Exact</matchingMethod>
        </matchingRuleItems>
        <ruleStatus>Active</ruleStatus>
    </matchingRules>
</MatchingRules>
`)
}

// A file without matching rules keeps the blank line Salesforce writes after
// the empty root element.
func TestEmptyMatchingRulesRoundTrip(t *testing.T) {
	assertRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<MatchingRules xmlns="http://soap.sforce.com/2006/04/metadata"/>

`)
}
