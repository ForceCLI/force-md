package components

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

// description is written between apiVersion and label, as Salesforce
// writes it.
func TestApexComponentDescriptionOrder(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<ApexComponent xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>63.0</apiVersion>
    <description>Default site footer component</description>
    <label>SiteFooter</label>
</ApexComponent>
`
	path := filepath.Join(t.TempDir(), "SiteFooter.component-meta.xml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	c, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	out, err := internal.Marshal(c)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != content {
		t.Errorf("round trip changed the component:\n%s", out)
	}
}
