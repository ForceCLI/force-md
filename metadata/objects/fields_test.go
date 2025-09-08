package objects

import (
	"testing"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/metadata/objects/field"
	"github.com/stretchr/testify/assert"
)

func TestCloneField(t *testing.T) {
	// Create a test object with a field to clone
	obj := &CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		Fields: []field.Field{
			{
				FullName:       "Source_Field__c",
				Type:           &TextLiteral{Text: "Text"},
				Label:          &TextLiteral{Text: "Source Field"},
				Length:         &IntegerText{Text: "255"},
				Required:       &BooleanText{Text: "false"},
				Unique:         &BooleanText{Text: "false"},
				ExternalId:     &BooleanText{Text: "false"},
				Description:    &TextLiteral{Text: "This is a source field"},
				InlineHelpText: &TextLiteral{Text: "Help text for source field"},
			},
			{
				FullName:  "Another_Field__c",
				Type:      &TextLiteral{Text: "Number"},
				Label:     &TextLiteral{Text: "Another Field"},
				Precision: &IntegerText{Text: "18"},
				Scale:     &IntegerText{Text: "2"},
			},
		},
	}

	t.Run("clone_existing_field", func(t *testing.T) {
		// Clone Source_Field__c to Target_Field__c
		err := obj.CloneField("Source_Field__c", "Target_Field__c")
		assert.NoError(t, err)

		// Check that Target_Field__c exists
		targetFieldFound := false
		var targetField field.Field
		for _, f := range obj.Fields {
			if f.FullName == "Target_Field__c" {
				targetFieldFound = true
				targetField = f
				break
			}
		}

		assert.True(t, targetFieldFound, "Target field was not created")

		// Verify the target field has the same properties as source
		assert.NotNil(t, targetField.Type)
		assert.Equal(t, "Text", targetField.Type.Text)

		assert.NotNil(t, targetField.Length)
		assert.Equal(t, "255", targetField.Length.Text)

		assert.NotNil(t, targetField.Description)
		assert.Equal(t, "This is a source field", targetField.Description.Text)

		assert.NotNil(t, targetField.InlineHelpText)
		assert.Equal(t, "Help text for source field", targetField.InlineHelpText.Text)

		// Verify the label was updated
		assert.NotNil(t, targetField.Label)
		assert.Equal(t, "Target Field", targetField.Label.Text)

		// Verify source field still exists
		sourceFieldFound := false
		for _, f := range obj.Fields {
			if f.FullName == "Source_Field__c" {
				sourceFieldFound = true
				break
			}
		}
		assert.True(t, sourceFieldFound, "Source field should still exist")

		// Verify we now have 3 fields
		assert.Equal(t, 3, len(obj.Fields))
	})

	t.Run("clone_nonexistent_field", func(t *testing.T) {
		// Reset object
		obj.Fields = []field.Field{
			{
				FullName: "Source_Field__c",
				Type:     &TextLiteral{Text: "Text"},
			},
			{
				FullName: "Another_Field__c",
				Type:     &TextLiteral{Text: "Number"},
			},
		}

		// Try to clone a field that doesn't exist
		err := obj.CloneField("NonExistent_Field__c", "Target_Field__c")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "source field not found")

		// Verify no new field was created
		assert.Equal(t, 2, len(obj.Fields))
	})

	t.Run("clone_to_existing_field", func(t *testing.T) {
		// Reset object
		obj.Fields = []field.Field{
			{
				FullName: "Source_Field__c",
				Type:     &TextLiteral{Text: "Text"},
			},
			{
				FullName: "Another_Field__c",
				Type:     &TextLiteral{Text: "Number"},
			},
		}

		// Try to clone to a field that already exists
		err := obj.CloneField("Source_Field__c", "Another_Field__c")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "target field already exists")

		// Verify the field count hasn't changed
		assert.Equal(t, 2, len(obj.Fields))

		// Verify Another_Field__c wasn't overwritten
		for _, f := range obj.Fields {
			if f.FullName == "Another_Field__c" {
				assert.NotNil(t, f.Type)
				assert.Equal(t, "Number", f.Type.Text)
				break
			}
		}
	})

	t.Run("clone_field_case_insensitive", func(t *testing.T) {
		// Reset object
		obj.Fields = []field.Field{
			{
				FullName: "Source_Field__c",
				Type:     &TextLiteral{Text: "Text"},
				Label:    &TextLiteral{Text: "Source Field"},
			},
		}

		// Clone with different case
		err := obj.CloneField("source_field__c", "Target_Field__c")
		assert.NoError(t, err)

		// Verify the field was cloned
		assert.Equal(t, 2, len(obj.Fields))

		// Try to clone to existing field with different case
		err = obj.CloneField("Source_Field__c", "target_field__c")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "target field already exists")
	})
}
