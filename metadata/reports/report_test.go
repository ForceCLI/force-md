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
		Name:       *NewTextLiteral("TestReport"),
		Format:     *NewTextLiteral("Tabular"),
		ReportType: *NewTextLiteral("AccountList"),
		Columns: []ReportColumn{
			{
				Field: *NewTextLiteral("ACCOUNT.NAME"),
			},
			{
				Field: *NewTextLiteral("CRC_Inquiry__c$TestField__c"),
			},
			{
				Field: *NewTextLiteral("ACCOUNT.TYPE"),
			},
		},
		GroupingsAcross: []ReportGroupingAcross{
			{
				Field:     *NewTextLiteral("CRC_Inquiry__c$TestField__c"),
				SortOrder: *NewTextLiteral("Asc"),
			},
			{
				Field:     *NewTextLiteral("ACCOUNT.TYPE"),
				SortOrder: *NewTextLiteral("Asc"),
			},
		},
		GroupingsDown: []ReportGroupingDown{
			{
				Field:     *NewTextLiteral("AnotherField__c"),
				SortOrder: *NewTextLiteral("Asc"),
			},
		},
		Filter: &ReportFilter{
			CriteriaItems: []ReportFilterCriteriaItem{
				{
					Column:   *NewTextLiteral("CRC_Inquiry__c$TestField__c"),
					Operator: *NewTextLiteral("equals"),
					Value:    *NewTextLiteral("TestValue"),
				},
				{
					Column:   *NewTextLiteral("ACCOUNT.NAME"),
					Operator: *NewTextLiteral("notEqual"),
					Value:    *NewTextLiteral(""),
				},
			},
		},
		Buckets: []ReportBucket{
			{
				DeveloperName:    *NewTextLiteral("TestBucket"),
				SourceColumnName: *NewTextLiteral("CRC_Inquiry__c$TestField__c"),
			},
			{
				DeveloperName:    *NewTextLiteral("AnotherBucket"),
				SourceColumnName: *NewTextLiteral("ACCOUNT.TYPE"),
			},
		},
	}

	t.Run("delete_existing_field_from_all_locations", func(t *testing.T) {
		// Delete TestField__c which appears in columns, groupingsAcross, filter, and buckets
		err := testReport.DeleteField("TestField__c")
		assert.NoError(t, err)

		// Check that TestField__c (with object prefix) is not in columns
		for _, column := range testReport.Columns {
			if column.Field.String() == "CRC_Inquiry__c$TestField__c" || column.Field.String() == "TestField__c" {
				t.Error("TestField__c should have been deleted from columns")
			}
		}

		// Check that TestField__c (with object prefix) is not in groupingsAcross
		for _, grouping := range testReport.GroupingsAcross {
			if grouping.Field.String() == "CRC_Inquiry__c$TestField__c" || grouping.Field.String() == "TestField__c" {
				t.Error("TestField__c should have been deleted from groupingsAcross")
			}
		}

		// Check that TestField__c (with object prefix) is not in filter criteria
		if testReport.Filter != nil {
			for _, criteria := range testReport.Filter.CriteriaItems {
				if criteria.Column.String() == "CRC_Inquiry__c$TestField__c" || criteria.Column.String() == "TestField__c" {
					t.Error("TestField__c should have been deleted from filter criteria")
				}
			}
		}

		// Check that TestField__c (with object prefix) is not in buckets
		for _, bucket := range testReport.Buckets {
			if bucket.SourceColumnName.String() == "CRC_Inquiry__c$TestField__c" || bucket.SourceColumnName.String() == "TestField__c" {
				t.Error("TestField__c should have been deleted from buckets")
			}
		}

		// Verify other fields are still present
		accountNameFound := false
		accountTypeFound := false
		for _, column := range testReport.Columns {
			if column.Field.String() == "ACCOUNT.NAME" {
				accountNameFound = true
			}
			if column.Field.String() == "ACCOUNT.TYPE" {
				accountTypeFound = true
			}
		}

		assert.True(t, accountNameFound, "ACCOUNT.NAME should still be present in columns")
		assert.True(t, accountTypeFound, "ACCOUNT.TYPE should still be present in columns")

		// Verify other groupings are still present
		accountTypeGroupingFound := false
		for _, grouping := range testReport.GroupingsAcross {
			if grouping.Field.String() == "ACCOUNT.TYPE" {
				accountTypeGroupingFound = true
			}
		}
		assert.True(t, accountTypeGroupingFound, "ACCOUNT.TYPE should still be present in groupingsAcross")

		// Verify other buckets are still present
		anotherBucketFound := false
		for _, bucket := range testReport.Buckets {
			if bucket.SourceColumnName.String() == "ACCOUNT.TYPE" {
				anotherBucketFound = true
			}
		}
		assert.True(t, anotherBucketFound, "ACCOUNT.TYPE bucket should still be present")
	})

	t.Run("delete_nonexistent_field", func(t *testing.T) {
		// Reset the report for this test
		testReport.Columns = []ReportColumn{
			{
				Field: *NewTextLiteral("ACCOUNT.NAME"),
			},
			{
				Field: *NewTextLiteral("ACCOUNT.TYPE"),
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
					Field: *NewTextLiteral("Account.Name"),
				},
				{
					Field: *NewTextLiteral("CustomObject__c$CustomField__c"),
				},
			},
		}

		// Delete using just the field name (without object prefix)
		err := report.DeleteField("CustomField__c")
		assert.NoError(t, err)

		// Verify the field was deleted
		assert.Equal(t, 1, len(report.Columns))
		assert.Equal(t, "Account.Name", report.Columns[0].Field.String())
	})
}
