package pathassistant

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	content := `<?xml version="1.0" encoding="UTF-8"?>
<PathAssistant xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <entityName>Widget__c</entityName>
    <fieldName>Stage__c</fieldName>
    <masterLabel>Widget Path</masterLabel>
    <pathAssistantSteps>
        <fieldNames>Name</fieldNames>
        <fieldNames>Owner__c</fieldNames>
        <info>&lt;p&gt;Call the customer.&lt;/p&gt;</info>
        <picklistValueName>Working</picklistValueName>
    </pathAssistantSteps>
    <pathAssistantSteps>
        <picklistValueName>Closed</picklistValueName>
    </pathAssistantSteps>
    <recordTypeName>__MASTER__</recordTypeName>
</PathAssistant>`
	path := filepath.Join(t.TempDir(), "Widget_Path.pathAssistant-meta.xml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	p, err := Open(path)
	assert.NoError(t, err)
	assert.True(t, p.Active.ToBool())
	assert.Equal(t, "Widget__c", p.EntityName.Text)
	assert.Equal(t, "Stage__c", p.FieldName.Text)
	assert.Equal(t, "Widget Path", p.MasterLabel.Text)
	assert.Equal(t, "__MASTER__", p.RecordTypeName.Text)
	assert.Len(t, p.PathAssistantSteps, 2)
	first := p.PathAssistantSteps[0]
	assert.Equal(t, "Working", first.PicklistValueName.Text)
	assert.Len(t, first.FieldNames, 2)
	assert.Equal(t, "Name", first.FieldNames[0].Text)
	assert.Equal(t, "<p>Call the customer.</p>", first.Info.String())
	second := p.PathAssistantSteps[1]
	assert.Equal(t, "Closed", second.PicklistValueName.Text)
	assert.Nil(t, second.Info)
}
