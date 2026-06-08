package datacategorygroup

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenParsesDataCategoryGroup(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Regions.datacategorygroup-meta.xml")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<DataCategoryGroup xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <dataCategory>
        <label>All Regions</label>
        <name>All</name>
        <dataCategory>
            <label>Americas</label>
            <name>Americas</name>
            <dataCategory>
                <label>United States</label>
                <name>USA</name>
            </dataCategory>
        </dataCategory>
    </dataCategory>
    <label>Regions</label>
    <objectUsage>
        <object>KnowledgeArticleVersion</object>
    </objectUsage>
</DataCategoryGroup>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := string(m.Name()); got != "Regions" {
		t.Errorf("name = %q, want Regions", got)
	}
	if m.Label != "Regions" {
		t.Errorf("label = %q, want Regions", m.Label)
	}
	if !m.Active.ToBool() {
		t.Errorf("active = false, want true")
	}
	if m.DataCategory.Name != "All" {
		t.Errorf("root category name = %q, want All", m.DataCategory.Name)
	}
	if len(m.DataCategory.DataCategory) != 1 {
		t.Fatalf("root child count = %d, want 1", len(m.DataCategory.DataCategory))
	}
	americas := m.DataCategory.DataCategory[0]
	if americas.Name != "Americas" {
		t.Errorf("child name = %q, want Americas", americas.Name)
	}
	if len(americas.DataCategory) != 1 || americas.DataCategory[0].Name != "USA" {
		t.Errorf("Americas children = %+v, want one USA", americas.DataCategory)
	}
	if m.ObjectUsage == nil || len(m.ObjectUsage.Object) != 1 || m.ObjectUsage.Object[0] != "KnowledgeArticleVersion" {
		t.Errorf("objectUsage = %+v, want [KnowledgeArticleVersion]", m.ObjectUsage)
	}
}
