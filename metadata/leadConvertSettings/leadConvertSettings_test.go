package leadConvertSettings

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// Every field mapping of an object mapping is kept, not only the last.
func TestLeadConvertSettingsMappingFieldsRoundTrip(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<LeadConvertSettings xmlns="http://soap.sforce.com/2006/04/metadata">
    <allowOwnerChange>true</allowOwnerChange>
    <objectMapping>
        <inputObject>Lead</inputObject>
        <mappingFields>
            <inputField>X18_Digit_Id__c</inputField>
            <outputField>Lead_Id__c</outputField>
        </mappingFields>
        <mappingFields>
            <inputField>Account_Classification__c</inputField>
            <outputField>Account_Classification__c</outputField>
        </mappingFields>
        <outputObject>Account</outputObject>
    </objectMapping>
    <objectMapping>
        <inputObject>Lead</inputObject>
        <outputObject>Opportunity</outputObject>
    </objectMapping>
    <opportunityCreationOptions>VisibleOptional</opportunityCreationOptions>
</LeadConvertSettings>
`
	path := filepath.Join(t.TempDir(), "LeadConvertSettings.LeadConvertSetting")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	l, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(l)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the lead convert settings:\n%s", out)
	}
}
