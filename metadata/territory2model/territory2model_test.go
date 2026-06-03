package territory2model

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenParsesTerritory2Model(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Geographic.territory2Model")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Model xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>Geographic</name>
    <description>Top level model</description>
</Territory2Model>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if m.Name.Text != "Geographic" {
		t.Errorf("name = %q, want Geographic", m.Name.Text)
	}
	if m.Description == nil || m.Description.Text != "Top level model" {
		t.Errorf("description = %v, want Top level model", m.Description)
	}
	if m.Type() != NAME {
		t.Errorf("Type() = %q, want %q", m.Type(), NAME)
	}
}
