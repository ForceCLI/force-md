package objects

import (
	"testing"

	"github.com/ForceCLI/force-md/metadata/objects/field"
	"github.com/ForceCLI/force-md/metadata/objects/listview"
	"github.com/ForceCLI/force-md/metadata/objects/recordtype"
	"github.com/ForceCLI/force-md/metadata/objects/validationrule"
	"github.com/ForceCLI/force-md/metadata/objects/weblink"
	"github.com/stretchr/testify/assert"
)

func TestCustomObjectTidy(t *testing.T) {
	obj := &CustomObject{
		Fields: FieldList{
			field.Field{FullName: "ZField__c"},
			field.Field{FullName: "AField__c"},
			field.Field{FullName: "MField__c"},
		},
		ListViews: []listview.ListView{
			{FullName: struct {
				Text string `xml:",chardata"`
			}{Text: "ZView"}},
			{FullName: struct {
				Text string `xml:",chardata"`
			}{Text: "AView"}},
		},
		RecordTypes: []recordtype.RecordType{
			{FullName: "ZType"},
			{FullName: "AType"},
		},
		ValidationRules: validationrule.ValidationRuleList{
			{FullName: "ZRule"},
			{FullName: "ARule"},
		},
		WebLinks: []weblink.WebLink{
			{FullName: "ZLink"},
			{FullName: "ALink"},
		},
	}

	obj.Tidy()

	// Verify fields are sorted
	assert.Equal(t, "AField__c", obj.Fields[0].FullName)
	assert.Equal(t, "MField__c", obj.Fields[1].FullName)
	assert.Equal(t, "ZField__c", obj.Fields[2].FullName)

	// Verify list views are sorted
	assert.Equal(t, "AView", obj.ListViews[0].FullName.Text)
	assert.Equal(t, "ZView", obj.ListViews[1].FullName.Text)

	// Verify record types are sorted
	assert.Equal(t, "AType", obj.RecordTypes[0].FullName)
	assert.Equal(t, "ZType", obj.RecordTypes[1].FullName)

	// Verify validation rules are sorted
	assert.Equal(t, "ARule", obj.ValidationRules[0].FullName)
	assert.Equal(t, "ZRule", obj.ValidationRules[1].FullName)

	// Verify weblinks are sorted
	assert.Equal(t, "ALink", obj.WebLinks[0].FullName)
	assert.Equal(t, "ZLink", obj.WebLinks[1].FullName)
}
