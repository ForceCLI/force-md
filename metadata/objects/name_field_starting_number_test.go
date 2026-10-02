package objects

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// An auto-number Name field's startingNumber is read, and written back
// between its label and type.
func TestNameFieldStartingNumberRoundTrips(t *testing.T) {
	path := filepath.Join(t.TempDir(), "Invoice__c.object-meta.xml")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
    <deploymentStatus>Deployed</deploymentStatus>
    <label>Invoice</label>
    <nameField>
        <displayFormat>INV-{0000}</displayFormat>
        <label>Invoice Number</label>
        <startingNumber>0</startingNumber>
        <type>AutoNumber</type>
    </nameField>
    <pluralLabel>Invoices</pluralLabel>
    <sharingModel>ReadWrite</sharingModel>
</CustomObject>
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	o, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if o.NameField == nil || o.NameField.StartingNumber == nil || o.NameField.StartingNumber.Text != "0" {
		t.Fatalf("NameField.StartingNumber = %+v, want 0", o.NameField)
	}
	out, err := internal.Marshal(o)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if !strings.Contains(string(out), "<label>Invoice Number</label>\n        <startingNumber>0</startingNumber>\n        <type>AutoNumber</type>") {
		t.Errorf("startingNumber not written in place:\n%s", out)
	}
}
