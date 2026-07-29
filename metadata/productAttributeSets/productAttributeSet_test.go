package productAttributeSets

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleProductAttributeSet = `<?xml version="1.0" encoding="UTF-8"?>
<ProductAttributeSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <description>Sizing attributes for apparel products</description>
    <developerName>Apparel_Sizing</developerName>
    <masterLabel>Apparel Sizing</masterLabel>
</ProductAttributeSet>
`

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Apparel_Sizing.productAttributeSet-meta.xml")
	if err := os.WriteFile(path, []byte(sampleProductAttributeSet), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	p, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := p.MasterLabel.String(); got != "Apparel Sizing" {
		t.Errorf("MasterLabel = %q, want %q", got, "Apparel Sizing")
	}
	if got := p.DeveloperName.String(); got != "Apparel_Sizing" {
		t.Errorf("DeveloperName = %q, want %q", got, "Apparel_Sizing")
	}
	if got := p.Description.String(); got != "Sizing attributes for apparel products" {
		t.Errorf("Description = %q, want %q", got, "Sizing attributes for apparel products")
	}
	if got := string(p.Type()); got != "ProductAttributeSet" {
		t.Errorf("Type = %q, want ProductAttributeSet", got)
	}
}

func TestOpenWithoutOptionalFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Bike_Components.productAttributeSet-meta.xml")
	minimal := `<?xml version="1.0" encoding="UTF-8"?>
<ProductAttributeSet xmlns="http://soap.sforce.com/2006/04/metadata">
    <masterLabel>Bike Components</masterLabel>
</ProductAttributeSet>
`
	if err := os.WriteFile(path, []byte(minimal), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	p, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if got := p.MasterLabel.String(); got != "Bike Components" {
		t.Errorf("MasterLabel = %q, want %q", got, "Bike Components")
	}
	if p.Description != nil {
		t.Errorf("Description = %v, want nil", p.Description)
	}
	if got := p.Description.String(); got != "" {
		t.Errorf("nil Description String() = %q, want empty", got)
	}
}
