package sharingrules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCriteriaRuleUsesField(t *testing.T) {
	rule := CriteriaRule{
		CriteriaItems: []CriteriaItem{
			{
				Field: Field{Text: "TestField__c"},
			},
			{
				Field: Field{Text: "AnotherField__c"},
			},
		},
	}

	t.Run("uses_existing_field", func(t *testing.T) {
		assert.True(t, rule.UsesField("TestField__c"), "Expected rule to use TestField__c")
		assert.True(t, rule.UsesField("AnotherField__c"), "Expected rule to use AnotherField__c")
	})

	t.Run("does_not_use_nonexistent_field", func(t *testing.T) {
		assert.False(t, rule.UsesField("NonExistentField__c"), "Expected rule not to use NonExistentField__c")
	})

	t.Run("case_insensitive_matching", func(t *testing.T) {
		assert.True(t, rule.UsesField("testfield__c"), "Expected case-insensitive field matching")
		assert.True(t, rule.UsesField("TESTFIELD__C"), "Expected case-insensitive field matching")
		assert.True(t, rule.UsesField("anotherField__c"), "Expected case-insensitive field matching")
	})
}
