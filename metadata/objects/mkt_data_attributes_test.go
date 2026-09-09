package objects

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMktDataLakeAttributes(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
	<mktDataLakeAttributes>
		<creationType>Custom</creationType>
		<isEnabled>true</isEnabled>
		<objectCategory>Salesforce_SFDCReferenceModel_0_93.Profile</objectCategory>
	</mktDataLakeAttributes>
</CustomObject>
`)
	obj := &CustomObject{}
	require.NoError(t, xml.Unmarshal(src, obj))
	require.NotNil(t, obj.MktDataLakeAttributes)
	assert.Nil(t, obj.MktDataModelAttributes)
	assert.Equal(t, "Custom", obj.MktDataLakeAttributes.CreationType.Text)
	assert.Equal(t, "true", obj.MktDataLakeAttributes.IsEnabled.Text)
	assert.Equal(t, "Salesforce_SFDCReferenceModel_0_93.Profile", obj.MktDataLakeAttributes.ObjectCategory.Text)

	out, err := xml.Marshal(obj)
	require.NoError(t, err)
	assert.Contains(t, string(out), "<mktDataLakeAttributes><creationType>Custom</creationType>")
}

func TestParseMktDataModelAttributes(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
	<mktDataModelAttributes>
		<creationType>Custom</creationType>
		<dataModelTaxonomy>Salesforce_SFDCReferenceModel_0_93</dataModelTaxonomy>
		<dataSpaceName>default</dataSpaceName>
		<isEnabled>true</isEnabled>
		<isSegmentable>true</isSegmentable>
		<isSqlDmo>false</isSqlDmo>
		<isUsedForMetrics>false</isUsedForMetrics>
		<labelOverride>Ad_Leads-14Days</labelOverride>
		<masterLabel>Ad_Leads-14Days</masterLabel>
		<objectCategory>Salesforce_SFDCReferenceModel_0_93.Profile</objectCategory>
	</mktDataModelAttributes>
</CustomObject>
`)
	obj := &CustomObject{}
	require.NoError(t, xml.Unmarshal(src, obj))
	require.NotNil(t, obj.MktDataModelAttributes)
	assert.Nil(t, obj.MktDataLakeAttributes)
	assert.Equal(t, "Custom", obj.MktDataModelAttributes.CreationType.Text)
	assert.Equal(t, "default", obj.MktDataModelAttributes.DataSpaceName.Text)
	assert.Equal(t, "Ad_Leads-14Days", obj.MktDataModelAttributes.MasterLabel.Text)
	assert.Equal(t, "false", obj.MktDataModelAttributes.IsSqlDmo.Text)
}
