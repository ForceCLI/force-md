package objectTranslations

import (
	"testing"

	"github.com/ForceCLI/force-md/metadata/objectTranslations/field"
	"github.com/ForceCLI/force-md/metadata/objectTranslations/recordtype"
	"github.com/ForceCLI/force-md/metadata/objectTranslations/validationrule"
	"github.com/stretchr/testify/assert"
)

func TestCustomObjectTranslationTidy(t *testing.T) {
	trans := &CustomObjectTranslation{
		Fields: FieldList{
			field.Field{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "ZField__c"}},
			field.Field{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "AField__c"}},
			field.Field{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "MField__c"}},
		},
		RecordTypes: RecordTypeList{
			recordtype.RecordType{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "ZType"}},
			recordtype.RecordType{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "AType"}},
		},
		ValidationRules: ValidationRuleList{
			validationrule.ValidationRule{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "ZRule"}},
			validationrule.ValidationRule{Name: struct {
				Text string `xml:",chardata"`
			}{Text: "ARule"}},
		},
	}

	trans.Tidy()

	// Verify fields are sorted
	assert.Equal(t, "AField__c", trans.Fields[0].Name.Text)
	assert.Equal(t, "MField__c", trans.Fields[1].Name.Text)
	assert.Equal(t, "ZField__c", trans.Fields[2].Name.Text)

	// Verify record types are sorted
	assert.Equal(t, "AType", trans.RecordTypes[0].Name.Text)
	assert.Equal(t, "ZType", trans.RecordTypes[1].Name.Text)

	// Verify validation rules are sorted
	assert.Equal(t, "ARule", trans.ValidationRules[0].Name.Text)
	assert.Equal(t, "ZRule", trans.ValidationRules[1].Name.Text)
}
