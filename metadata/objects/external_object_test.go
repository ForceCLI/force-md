package objects

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExternalObject(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
	<deploymentStatus>Deployed</deploymentStatus>
	<externalDataSource>Orders_Source</externalDataSource>
	<externalName>orders</externalName>
	<externalRepository>warehouse</externalRepository>
	<fields>
		<fullName>Total__c</fullName>
		<externalDeveloperName>total</externalDeveloperName>
		<externalId>false</externalId>
		<label>Total</label>
		<readOnlyProxy>true</readOnlyProxy>
		<type>Number</type>
	</fields>
	<label>Order</label>
	<pluralLabel>Orders</pluralLabel>
</CustomObject>
`)
	obj := &CustomObject{}
	err := xml.Unmarshal(src, obj)
	assert.NoError(t, err)
	if assert.NotNil(t, obj.ExternalDataSource) {
		assert.Equal(t, "Orders_Source", obj.ExternalDataSource.Text)
	}
	if assert.NotNil(t, obj.ExternalName) {
		assert.Equal(t, "orders", obj.ExternalName.Text)
	}
	if assert.NotNil(t, obj.ExternalRepository) {
		assert.Equal(t, "warehouse", obj.ExternalRepository.Text)
	}
	if assert.Len(t, obj.Fields, 1) {
		if assert.NotNil(t, obj.Fields[0].ExternalDeveloperName) {
			assert.Equal(t, "total", obj.Fields[0].ExternalDeveloperName.Text)
		}
		if assert.NotNil(t, obj.Fields[0].ReadOnlyProxy) {
			assert.Equal(t, "true", obj.Fields[0].ReadOnlyProxy.Text)
		}
	}

	out, err := xml.Marshal(obj)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "<externalDataSource>Orders_Source</externalDataSource><externalName>orders</externalName><externalRepository>warehouse</externalRepository>")
	assert.Contains(t, string(out), "<externalDeveloperName>total</externalDeveloperName><externalId>false</externalId>")
	assert.Contains(t, string(out), "<readOnlyProxy>true</readOnlyProxy>")
}
