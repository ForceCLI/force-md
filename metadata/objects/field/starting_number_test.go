package field

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// An auto-number field's startingNumber is read, and written back in its
// place.
func TestStartingNumberRoundTrips(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Seq__c.field-meta.xml")
	content := `<?xml version="1.0" encoding="UTF-8"?>
<CustomField xmlns="http://soap.sforce.com/2006/04/metadata">
    <fullName>Seq__c</fullName>
    <displayFormat>S-{0000}</displayFormat>
    <label>Seq</label>
    <startingNumber>100</startingNumber>
    <type>AutoNumber</type>
</CustomField>
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	f, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if f.StartingNumber == nil || f.StartingNumber.Text != "100" {
		t.Fatalf("StartingNumber = %+v, want 100", f.StartingNumber)
	}
	out, err := internal.Marshal(f)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if !strings.Contains(string(out), "    <label>Seq</label>\n    <startingNumber>100</startingNumber>\n    <type>AutoNumber</type>") {
		t.Errorf("startingNumber not written in place:\n%s", out)
	}
}
