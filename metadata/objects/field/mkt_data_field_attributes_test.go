package field

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMktDataLakeFieldAttributes(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomField xmlns="http://soap.sforce.com/2006/04/metadata">
	<fullName>sfdc_app_track_record_id__c</fullName>
	<label>sfdc_app_track_record_id</label>
	<length>60</length>
	<mktDataLakeFieldAttributes>
		<definitionCreationType>Custom</definitionCreationType>
		<externalName>sfdc_app_track_record_id</externalName>
		<isEventDate>false</isEventDate>
		<isInternalOrganization>false</isInternalOrganization>
		<isRecordModified>false</isRecordModified>
		<primaryIndexOrder>1</primaryIndexOrder>
		<usageTag>NONE</usageTag>
	</mktDataLakeFieldAttributes>
	<required>false</required>
	<type>Text</type>
</CustomField>
`)
	f := &Field{}
	require.NoError(t, xml.Unmarshal(src, f))
	require.NotNil(t, f.MktDataLakeFieldAttributes)
	assert.Nil(t, f.MktDataModelFieldAttributes)
	assert.Equal(t, "Custom", f.MktDataLakeFieldAttributes.DefinitionCreationType.Text)
	assert.Equal(t, "sfdc_app_track_record_id", f.MktDataLakeFieldAttributes.ExternalName.Text)
	assert.Equal(t, "1", f.MktDataLakeFieldAttributes.PrimaryIndexOrder.Text)
	assert.Equal(t, "NONE", f.MktDataLakeFieldAttributes.UsageTag.Text)
}

func TestParseMktDataModelFieldAttributes(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomField xmlns="http://soap.sforce.com/2006/04/metadata">
	<fullName>KQ_id__c</fullName>
	<label>Key Qualifier Id</label>
	<length>60</length>
	<mktDataModelFieldAttributes>
		<definitionCreationType>System</definitionCreationType>
		<isDynamicLookup>false</isDynamicLookup>
		<labelOverride>Key Qualifier Id</labelOverride>
		<masterLabel>Key Qualifier Id</masterLabel>
		<refAttrDeveloperName>KQ_id</refAttrDeveloperName>
		<usageTag>KeyQualifier</usageTag>
	</mktDataModelFieldAttributes>
	<required>false</required>
	<type>Text</type>
</CustomField>
`)
	f := &Field{}
	require.NoError(t, xml.Unmarshal(src, f))
	require.NotNil(t, f.MktDataModelFieldAttributes)
	assert.Nil(t, f.MktDataLakeFieldAttributes)
	assert.Equal(t, "System", f.MktDataModelFieldAttributes.DefinitionCreationType.Text)
	assert.Equal(t, "KQ_id", f.MktDataModelFieldAttributes.RefAttrDeveloperName.Text)
	assert.Equal(t, "KeyQualifier", f.MktDataModelFieldAttributes.UsageTag.Text)
	assert.Nil(t, f.MktDataModelFieldAttributes.InvalidMergeActionType)
}
