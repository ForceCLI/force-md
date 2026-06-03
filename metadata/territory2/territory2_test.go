package territory2

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenParsesTerritory2(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "California.territory2")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata">
    <accountAccessLevel>Edit</accountAccessLevel>
    <caseAccessLevel>Edit</caseAccessLevel>
    <contactAccessLevel>Edit</contactAccessLevel>
    <name>California</name>
    <opportunityAccessLevel>Read</opportunityAccessLevel>
    <ruleAssociations>
        <inherited>true</inherited>
        <ruleName>California</ruleName>
    </ruleAssociations>
    <territory2Type>State</territory2Type>
</Territory2>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if m.Name.Text != "California" {
		t.Errorf("name = %q, want California", m.Name.Text)
	}
	if m.AccountAccessLevel.Text != "Edit" {
		t.Errorf("accountAccessLevel = %q, want Edit", m.AccountAccessLevel.Text)
	}
	if m.OpportunityAccessLevel.Text != "Read" {
		t.Errorf("opportunityAccessLevel = %q, want Read", m.OpportunityAccessLevel.Text)
	}
	if m.Territory2Type.Text != "State" {
		t.Errorf("territory2Type = %q, want State", m.Territory2Type.Text)
	}
	if len(m.RuleAssociations) != 1 || m.RuleAssociations[0].RuleName.Text != "California" {
		t.Errorf("ruleAssociations = %v, want one entry for California", m.RuleAssociations)
	}
	if m.Type() != NAME {
		t.Errorf("Type() = %q, want %q", m.Type(), NAME)
	}
}
