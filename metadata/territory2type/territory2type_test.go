package territory2type

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenParsesTerritory2Type(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "State.territory2Type")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Type xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>State</name>
    <priority>1</priority>
</Territory2Type>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if m.Name.Text != "State" {
		t.Errorf("name = %q, want State", m.Name.Text)
	}
	if m.Priority.Text != "1" {
		t.Errorf("priority = %q, want 1", m.Priority.Text)
	}
	if m.Type() != NAME {
		t.Errorf("Type() = %q, want %q", m.Type(), NAME)
	}
}
