package field

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenReadsPicklistValueTranslations(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "objectTranslations", "Widget__c-fr")
	require.NoError(t, os.MkdirAll(dir, 0o755))
	path := filepath.Join(dir, "Kind__c.fieldTranslation-meta.xml")
	require.NoError(t, os.WriteFile(path, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomFieldTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Genre</label>
    <name>Kind__c</name>
    <picklistValues>
        <masterLabel>Alpha Label</masterLabel>
        <translation>Alpha &amp; FR</translation>
    </picklistValues>
    <picklistValues>
        <masterLabel>Beta Label</masterLabel>
        <translation><!-- Beta Label --></translation>
    </picklistValues>
</CustomFieldTranslation>
`), 0o644))

	trans, err := Open(path)
	require.NoError(t, err)
	require.Len(t, trans.PicklistValues, 2)
	assert.Equal(t, "Alpha Label", trans.PicklistValues[0].MasterLabel.Text)
	assert.Equal(t, "Alpha & FR", trans.PicklistValues[0].Translation.Text)
	assert.Equal(t, "Beta Label", trans.PicklistValues[1].MasterLabel.Text)
	assert.Equal(t, "", trans.PicklistValues[1].Translation.Text, "an untranslated value is a comment")
}
