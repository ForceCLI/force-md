package flexipage

import (
	"testing"

	. "github.com/ForceCLI/force-md/general"
	"github.com/stretchr/testify/assert"
)

func TestDeleteField(t *testing.T) {
	// Create a test flexipage with field instances
	testFlexiPage := &FlexiPage{
		Xmlns:         "http://soap.sforce.com/2006/04/metadata",
		MasterLabel:   TextLiteral{Text: "Test Page"},
		FlexiPageType: TextLiteral{Text: "RecordPage"},
		FlexiPageRegions: []FlexiPageRegion{
			{
				Name: TextLiteral{Text: "main"},
				Type: TextLiteral{Text: "Region"},
				ItemInstances: []ItemInstance{
					{
						FieldInstance: &FieldInstance{
							FieldItem: TextLiteral{Text: "Account.Name"},
						},
					},
					{
						FieldInstance: &FieldInstance{
							FieldItem: TextLiteral{Text: "Record.Share_With_Program_Text__c"},
						},
					},
					{
						FieldInstance: &FieldInstance{
							FieldItem: TextLiteral{Text: "Account.Type"},
						},
					},
				},
			},
		},
	}

	t.Run("delete_existing_field_with_record_prefix", func(t *testing.T) {
		// Delete Share_With_Program_Text__c (which has Record. prefix)
		err := testFlexiPage.DeleteField("Share_With_Program_Text__c")
		assert.NoError(t, err)

		// Verify the field was deleted
		assert.Equal(t, 2, len(testFlexiPage.FlexiPageRegions[0].ItemInstances))

		// Verify other fields are still present
		accountNameFound := false
		accountTypeFound := false
		for _, instance := range testFlexiPage.FlexiPageRegions[0].ItemInstances {
			if instance.FieldInstance != nil {
				fieldText := instance.FieldInstance.FieldItem.Text
				if fieldText == "Account.Name" {
					accountNameFound = true
				}
				if fieldText == "Account.Type" {
					accountTypeFound = true
				}
			}
		}
		assert.True(t, accountNameFound, "Account.Name should still be present")
		assert.True(t, accountTypeFound, "Account.Type should still be present")
	})

	t.Run("delete_existing_field_without_prefix", func(t *testing.T) {
		// Delete Account.Name (no Record. prefix)
		err := testFlexiPage.DeleteField("Account.Name")
		assert.NoError(t, err)

		// Verify the field was deleted (should now have only 1 field left since we already deleted one)
		assert.Equal(t, 1, len(testFlexiPage.FlexiPageRegions[0].ItemInstances))
	})

	t.Run("delete_nonexistent_field", func(t *testing.T) {
		// Try to delete a field that doesn't exist
		err := testFlexiPage.DeleteField("NonExistentField__c")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'NonExistentField__c' not found")

		// Verify no additional fields were deleted (still 1 from previous test)
		assert.Equal(t, 1, len(testFlexiPage.FlexiPageRegions[0].ItemInstances))
	})
}
