package territory2rule

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenParsesTerritory2Rule(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "California.territory2Rule")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Rule xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <name>California</name>
    <objectType>Account</objectType>
    <ruleItems>
        <field>Account.BillingState</field>
        <operation>equals</operation>
        <value>California</value>
    </ruleItems>
</Territory2Rule>`
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
	if m.Active.Text != "true" {
		t.Errorf("active = %q, want true", m.Active.Text)
	}
	if m.ObjectType.Text != "Account" {
		t.Errorf("objectType = %q, want Account", m.ObjectType.Text)
	}
	if len(m.RuleItems) != 1 {
		t.Fatalf("ruleItems length = %d, want 1", len(m.RuleItems))
	}
	item := m.RuleItems[0]
	if item.Field.Text != "Account.BillingState" || item.Operation.Text != "equals" || item.Value.Text != "California" {
		t.Errorf("ruleItem = %+v, want Account.BillingState/equals/California", item)
	}
	if m.Type() != NAME {
		t.Errorf("Type() = %q, want %q", m.Type(), NAME)
	}
}

func TestOpenLeavesOmittedRuleItemValueNil(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "No_City.territory2Rule")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Rule xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <name>No City</name>
    <objectType>Account</objectType>
    <ruleItems>
        <field>Account.BillingCity</field>
        <operation>equals</operation>
    </ruleItems>
</Territory2Rule>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(m.RuleItems) != 1 {
		t.Fatalf("ruleItems length = %d, want 1", len(m.RuleItems))
	}
	if m.RuleItems[0].Value != nil {
		t.Errorf("ruleItems[0].Value = %v, want nil", m.RuleItems[0].Value)
	}
}
