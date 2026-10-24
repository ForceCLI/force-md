package sharingrules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// Rules that share to a role and its internal subordinates, or to all
// internal users, are written back as Salesforce writes them.
func TestSharingRulesRoundTrip(t *testing.T) {
	contents := `<?xml version="1.0" encoding="UTF-8"?>
<SharingRules xmlns="http://soap.sforce.com/2006/04/metadata">
    <sharingCriteriaRules>
        <fullName>Admissions_Director</fullName>
        <accessLevel>Edit</accessLevel>
        <label>Admissions Director</label>
        <sharedTo>
            <roleAndSubordinatesInternal>Admissions_Director</roleAndSubordinatesInternal>
        </sharedTo>
        <criteriaItems>
            <field>Type</field>
            <operation>equals</operation>
            <value>Prospect</value>
        </criteriaItems>
        <includeRecordsOwnedByAll>true</includeRecordsOwnedByAll>
    </sharingCriteriaRules>
    <sharingOwnerRules>
        <fullName>All_Internal</fullName>
        <accessLevel>Read</accessLevel>
        <label>All Internal</label>
        <sharedTo>
            <allInternalUsers></allInternalUsers>
        </sharedTo>
        <sharedFrom>
            <roleAndSubordinatesInternal>Outreach_Manager</roleAndSubordinatesInternal>
        </sharedFrom>
    </sharingOwnerRules>
</SharingRules>
`
	path := filepath.Join(t.TempDir(), "Account.sharingRules")
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	r, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != contents {
		t.Errorf("round trip changed the sharing rules:\n%s", out)
	}
}
