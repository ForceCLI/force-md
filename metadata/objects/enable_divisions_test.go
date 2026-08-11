package objects

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnableDivisions(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
	<enableActivities>true</enableActivities>
	<enableDivisions>true</enableDivisions>
	<enableHistory>false</enableHistory>
</CustomObject>
`)
	obj := &CustomObject{}
	err := xml.Unmarshal(src, obj)
	assert.NoError(t, err)
	if assert.NotNil(t, obj.EnableDivisions) {
		assert.Equal(t, "true", obj.EnableDivisions.Text)
	}

	out, err := xml.Marshal(obj)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "<enableDivisions>true</enableDivisions>")
}

func TestParseWithoutEnableDivisions(t *testing.T) {
	src := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<CustomObject xmlns="http://soap.sforce.com/2006/04/metadata">
	<enableActivities>true</enableActivities>
</CustomObject>
`)
	obj := &CustomObject{}
	err := xml.Unmarshal(src, obj)
	assert.NoError(t, err)
	assert.Nil(t, obj.EnableDivisions)
}
