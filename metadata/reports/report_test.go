package report

import (
	"testing"

	. "github.com/ForceCLI/force-md/general"
	"github.com/stretchr/testify/assert"
)

func TestDeleteField(t *testing.T) {
	// Create a test report with field references in multiple places
	testReport := &Report{
		Xmlns:      "http://soap.sforce.com/2006/04/metadata",
		Name:       TextLiteral{Text: "TestReport"},
		Format:     TextLiteral{Text: "Tabular"},
		ReportType: TextLiteral{Text: "AccountList"},
		Columns: []ReportColumn{
			{
				Field: TextLiteral{Text: "ACCOUNT.NAME"},
			},
			{
				Field: TextLiteral{Text: "CRC_Inquiry__c$TestField__c"},
			},
			{
				Field: TextLiteral{Text: "ACCOUNT.TYPE"},
			},
		},
		GroupingsAcross: []ReportGroupingAcross{
			{
				Field:     TextLiteral{Text: "CRC_Inquiry__c$TestField__c"},
				SortOrder: TextLiteral{Text: "Asc"},
			},
			{
				Field:     TextLiteral{Text: "ACCOUNT.TYPE"},
				SortOrder: TextLiteral{Text: "Asc"},
			},
		},
		GroupingsDown: []ReportGroupingDown{
			{
				Field:     TextLiteral{Text: "AnotherField__c"},
				SortOrder: TextLiteral{Text: "Asc"},
			},
		},
		Filter: &ReportFilter{
			CriteriaItems: []ReportFilterCriteriaItem{
				{
					Column:   TextLiteral{Text: "CRC_Inquiry__c$TestField__c"},
					Operator: TextLiteral{Text: "equals"},
					Value:    TextLiteral{Text: "TestValue"},
				},
				{
					Column:   TextLiteral{Text: "ACCOUNT.NAME"},
					Operator: TextLiteral{Text: "notEqual"},
					Value:    TextLiteral{Text: ""},
				},
			},
		},
		Buckets: []ReportBucket{
			{
				DeveloperName:    TextLiteral{Text: "TestBucket"},
				SourceColumnName: TextLiteral{Text: "CRC_Inquiry__c$TestField__c"},
			},
			{
				DeveloperName:    TextLiteral{Text: "AnotherBucket"},
				SourceColumnName: TextLiteral{Text: "ACCOUNT.TYPE"},
			},
		},
	}

	t.Run("delete_existing_field_from_all_locations", func(t *testing.T) {
		// Delete TestField__c which appears in columns, groupingsAcross, filter, and buckets
		err := testReport.DeleteField("TestField__c")
		assert.NoError(t, err)

		// Check that TestField__c (with object prefix) is not in columns
		for _, column := range testReport.Columns {
			if column.Field.Text == "CRC_Inquiry__c$TestField__c" || column.Field.Text == "TestField__c" {
				t.Error("TestField__c should have been deleted from columns")
			}
		}

		// Check that TestField__c (with object prefix) is not in groupingsAcross
		for _, grouping := range testReport.GroupingsAcross {
			if grouping.Field.Text == "CRC_Inquiry__c$TestField__c" || grouping.Field.Text == "TestField__c" {
				t.Error("TestField__c should have been deleted from groupingsAcross")
			}
		}

		// Check that TestField__c (with object prefix) is not in filter criteria
		if testReport.Filter != nil {
			for _, criteria := range testReport.Filter.CriteriaItems {
				if criteria.Column.Text == "CRC_Inquiry__c$TestField__c" || criteria.Column.Text == "TestField__c" {
					t.Error("TestField__c should have been deleted from filter criteria")
				}
			}
		}

		// Check that TestField__c (with object prefix) is not in buckets
		for _, bucket := range testReport.Buckets {
			if bucket.SourceColumnName.Text == "CRC_Inquiry__c$TestField__c" || bucket.SourceColumnName.Text == "TestField__c" {
				t.Error("TestField__c should have been deleted from buckets")
			}
		}

		// Verify other fields are still present
		accountNameFound := false
		accountTypeFound := false
		for _, column := range testReport.Columns {
			if column.Field.Text == "ACCOUNT.NAME" {
				accountNameFound = true
			}
			if column.Field.Text == "ACCOUNT.TYPE" {
				accountTypeFound = true
			}
		}

		assert.True(t, accountNameFound, "ACCOUNT.NAME should still be present in columns")
		assert.True(t, accountTypeFound, "ACCOUNT.TYPE should still be present in columns")

		// Verify other groupings are still present
		accountTypeGroupingFound := false
		for _, grouping := range testReport.GroupingsAcross {
			if grouping.Field.Text == "ACCOUNT.TYPE" {
				accountTypeGroupingFound = true
			}
		}
		assert.True(t, accountTypeGroupingFound, "ACCOUNT.TYPE should still be present in groupingsAcross")

		// Verify other buckets are still present
		anotherBucketFound := false
		for _, bucket := range testReport.Buckets {
			if bucket.SourceColumnName.Text == "ACCOUNT.TYPE" {
				anotherBucketFound = true
			}
		}
		assert.True(t, anotherBucketFound, "ACCOUNT.TYPE bucket should still be present")
	})

	t.Run("delete_nonexistent_field", func(t *testing.T) {
		// Reset the report for this test
		testReport.Columns = []ReportColumn{
			{
				Field: TextLiteral{Text: "ACCOUNT.NAME"},
			},
			{
				Field: TextLiteral{Text: "ACCOUNT.TYPE"},
			},
		}

		// Try to delete a field that doesn't exist
		err := testReport.DeleteField("NonExistentField__c")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'NonExistentField__c' not found")

		// Verify report is unchanged
		assert.Equal(t, 2, len(testReport.Columns))
	})

	t.Run("delete_field_with_object_prefix", func(t *testing.T) {
		// Create a report with object-prefixed field
		report := &Report{
			Columns: []ReportColumn{
				{
					Field: TextLiteral{Text: "Account.Name"},
				},
				{
					Field: TextLiteral{Text: "CustomObject__c$CustomField__c"},
				},
			},
		}

		// Delete using just the field name (without object prefix)
		err := report.DeleteField("CustomField__c")
		assert.NoError(t, err)

		// Verify the field was deleted
		assert.Equal(t, 1, len(report.Columns))
		assert.Equal(t, "Account.Name", report.Columns[0].Field.Text)
	})
}
