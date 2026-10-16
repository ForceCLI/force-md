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

func TestOpenParsesTerritory2CustomFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "West.territory2")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <accountAccessLevel>Edit</accountAccessLevel>
    <caseAccessLevel>Edit</caseAccessLevel>
    <contactAccessLevel>Edit</contactAccessLevel>
    <customFields>
        <name>Activation_DateTime__c</name>
        <value xsi:type="xsd:dateTime">2014-07-16T05:05:00.000Z</value>
    </customFields>
    <customFields>
        <name>ManagersPhone__c</name>
        <value xsi:nil="true"/>
    </customFields>
    <name>West</name>
    <opportunityAccessLevel>Read</opportunityAccessLevel>
    <territory2Type>State</territory2Type>
</Territory2>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(m.CustomFields) != 2 {
		t.Fatalf("customFields length = %d, want 2", len(m.CustomFields))
	}
	activation := m.CustomFields[0]
	if activation.Name.Text != "Activation_DateTime__c" || activation.Value.Type != "xsd:dateTime" || activation.Value.Text != "2014-07-16T05:05:00.000Z" || activation.Value.Nil != "" {
		t.Errorf("customFields[0] = %+v, want Activation_DateTime__c xsd:dateTime 2014-07-16T05:05:00.000Z", activation)
	}
	phone := m.CustomFields[1]
	if phone.Name.Text != "ManagersPhone__c" || phone.Value.Nil != "true" || phone.Value.Text != "" || phone.Value.Type != "" {
		t.Errorf("customFields[1] = %+v, want ManagersPhone__c nil", phone)
	}
}

func TestOpenParsesTerritory2ObjectAccessLevels(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "USA.territory2")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata">
    <accountAccessLevel>Edit</accountAccessLevel>
    <caseAccessLevel>Edit</caseAccessLevel>
    <contactAccessLevel>Edit</contactAccessLevel>
    <name>USA</name>
    <objectAccessLevels>
        <accessLevel>All</accessLevel>
        <objectType>Lead</objectType>
    </objectAccessLevels>
    <opportunityAccessLevel>Read</opportunityAccessLevel>
    <territory2Type>Geo</territory2Type>
</Territory2>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if len(m.ObjectAccessLevels) != 1 {
		t.Fatalf("objectAccessLevels length = %d, want 1", len(m.ObjectAccessLevels))
	}
	level := m.ObjectAccessLevels[0]
	if level.AccessLevel.Text != "All" || level.ObjectType.Text != "Lead" {
		t.Errorf("objectAccessLevels[0] = %+v, want All on Lead", level)
	}
}

func TestOpenLeavesOmittedAccessLevelsNil(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "West.territory2")
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata">
    <accountAccessLevel>Edit</accountAccessLevel>
    <contactAccessLevel>Read</contactAccessLevel>
    <name>West</name>
    <territory2Type>State</territory2Type>
</Territory2>`
	if err := os.WriteFile(path, []byte(xml), 0644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	m, err := Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if m.AccountAccessLevel == nil || m.AccountAccessLevel.Text != "Edit" {
		t.Errorf("accountAccessLevel = %v, want Edit", m.AccountAccessLevel)
	}
	if m.ContactAccessLevel == nil || m.ContactAccessLevel.Text != "Read" {
		t.Errorf("contactAccessLevel = %v, want Read", m.ContactAccessLevel)
	}
	if m.CaseAccessLevel != nil {
		t.Errorf("caseAccessLevel = %v, want nil", m.CaseAccessLevel)
	}
	if m.OpportunityAccessLevel != nil {
		t.Errorf("opportunityAccessLevel = %v, want nil", m.OpportunityAccessLevel)
	}
}
