package tab

import (
	"os"
	"path/filepath"
	"testing"
)

// The Home tab's org-default Lightning page assignment is stored as an
// actionOverrides entry on standard-home.
const sampleHomeTab = `<?xml version="1.0" encoding="UTF-8"?>
<CustomTab xmlns="http://soap.sforce.com/2006/04/metadata">
    <actionOverrides>
        <actionName>Tab</actionName>
        <content>Home_Page_Default2</content>
        <formFactor>Large</formFactor>
        <skipRecordTypeSelect>false</skipRecordTypeSelect>
        <type>Flexipage</type>
    </actionOverrides>
</CustomTab>
`

func TestOpenActionOverrides(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "standard-home.tab-meta.xml")
	if err := os.WriteFile(path, []byte(sampleHomeTab), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	c, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(c.ActionOverrides) != 1 {
		t.Fatalf("ActionOverrides count = %d, want 1", len(c.ActionOverrides))
	}
	override := c.ActionOverrides[0]
	if override.ActionName != "Tab" {
		t.Errorf("ActionName = %q, want Tab", override.ActionName)
	}
	if override.Content == nil || *override.Content != "Home_Page_Default2" {
		t.Errorf("Content = %v, want Home_Page_Default2", override.Content)
	}
	if override.FormFactor == nil || *override.FormFactor != "Large" {
		t.Errorf("FormFactor = %v, want Large", override.FormFactor)
	}
	if override.Type != "Flexipage" {
		t.Errorf("Type = %q, want Flexipage", override.Type)
	}
}
