package sharingrules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/sharingrules"
)

func TestDeleteRule(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-sharingrules-delete-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test sharing rules file with both criteria and owner rules
	testSharingRules := &sharingrules.SharingRules{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		SharingCriteriaRules: []sharingrules.CriteriaRule{
			{
				FullName:    "TestCriteriaRule1",
				AccessLevel: sharingrules.AccessLevel{Text: "Read"},
				Label:       sharingrules.Label{Text: "Test Criteria Rule 1"},
			},
			{
				FullName:    "TestCriteriaRule2",
				AccessLevel: sharingrules.AccessLevel{Text: "Edit"},
				Label:       sharingrules.Label{Text: "Test Criteria Rule 2"},
			},
		},
		SharingOwnerRules: []sharingrules.OwnerRule{
			{
				FullName:    "TestOwnerRule1",
				AccessLevel: sharingrules.AccessLevel{Text: "Read"},
				Label:       sharingrules.Label{Text: "Test Owner Rule 1"},
			},
			{
				FullName:    "TestOwnerRule2",
				AccessLevel: sharingrules.AccessLevel{Text: "Edit"},
				Label:       sharingrules.Label{Text: "Test Owner Rule 2"},
			},
		},
	}

	sharingRulesPath := filepath.Join(tempDir, "Account.sharingRules")
	err = internal.WriteToFile(testSharingRules, sharingRulesPath)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("delete_criteria_rule_by_type", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		deleteRule(sharingRulesPath, "TestCriteriaRule1", "criteria")

		// Verify the rule was deleted
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Should have 1 criteria rule left
		if len(updatedRules.SharingCriteriaRules) != 1 {
			t.Errorf("Expected 1 criteria rule, got %d", len(updatedRules.SharingCriteriaRules))
		}

		// The remaining rule should be TestCriteriaRule2
		if len(updatedRules.SharingCriteriaRules) > 0 && updatedRules.SharingCriteriaRules[0].FullName != "TestCriteriaRule2" {
			t.Errorf("Expected remaining rule to be TestCriteriaRule2, got %s", updatedRules.SharingCriteriaRules[0].FullName)
		}

		// Owner rules should be unchanged
		if len(updatedRules.SharingOwnerRules) != 2 {
			t.Errorf("Expected 2 owner rules, got %d", len(updatedRules.SharingOwnerRules))
		}
	})

	t.Run("delete_owner_rule_by_type", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		deleteRule(sharingRulesPath, "TestOwnerRule2", "owner")

		// Verify the rule was deleted
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Should have 1 owner rule left
		if len(updatedRules.SharingOwnerRules) != 1 {
			t.Errorf("Expected 1 owner rule, got %d", len(updatedRules.SharingOwnerRules))
		}

		// The remaining rule should be TestOwnerRule1
		if len(updatedRules.SharingOwnerRules) > 0 && updatedRules.SharingOwnerRules[0].FullName != "TestOwnerRule1" {
			t.Errorf("Expected remaining rule to be TestOwnerRule1, got %s", updatedRules.SharingOwnerRules[0].FullName)
		}

		// Criteria rules should be unchanged
		if len(updatedRules.SharingCriteriaRules) != 2 {
			t.Errorf("Expected 2 criteria rules, got %d", len(updatedRules.SharingCriteriaRules))
		}
	})

	t.Run("delete_criteria_rule_auto_detect", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Delete without specifying type - should find criteria rule
		deleteRule(sharingRulesPath, "TestCriteriaRule1", "")

		// Verify the rule was deleted
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Should have 1 criteria rule left
		if len(updatedRules.SharingCriteriaRules) != 1 {
			t.Errorf("Expected 1 criteria rule, got %d", len(updatedRules.SharingCriteriaRules))
		}

		// The remaining rule should be TestCriteriaRule2
		if len(updatedRules.SharingCriteriaRules) > 0 && updatedRules.SharingCriteriaRules[0].FullName != "TestCriteriaRule2" {
			t.Errorf("Expected remaining rule to be TestCriteriaRule2, got %s", updatedRules.SharingCriteriaRules[0].FullName)
		}

		// Owner rules should be unchanged
		if len(updatedRules.SharingOwnerRules) != 2 {
			t.Errorf("Expected 2 owner rules, got %d", len(updatedRules.SharingOwnerRules))
		}
	})

	t.Run("delete_owner_rule_auto_detect", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Delete without specifying type - should find owner rule if criteria doesn't exist
		deleteRule(sharingRulesPath, "TestOwnerRule1", "")

		// Verify the rule was deleted
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Should have 1 owner rule left
		if len(updatedRules.SharingOwnerRules) != 1 {
			t.Errorf("Expected 1 owner rule, got %d", len(updatedRules.SharingOwnerRules))
		}

		// The remaining rule should be TestOwnerRule2
		if len(updatedRules.SharingOwnerRules) > 0 && updatedRules.SharingOwnerRules[0].FullName != "TestOwnerRule2" {
			t.Errorf("Expected remaining rule to be TestOwnerRule2, got %s", updatedRules.SharingOwnerRules[0].FullName)
		}

		// Criteria rules should be unchanged
		if len(updatedRules.SharingCriteriaRules) != 2 {
			t.Errorf("Expected 2 criteria rules, got %d", len(updatedRules.SharingCriteriaRules))
		}
	})

	t.Run("delete_nonexistent_rule", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		originalRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Try to delete a rule that doesn't exist
		deleteRule(sharingRulesPath, "NonExistentRule", "")

		// Verify nothing changed
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		if len(updatedRules.SharingCriteriaRules) != len(originalRules.SharingCriteriaRules) {
			t.Error("Criteria rules should be unchanged when deleting non-existent rule")
		}

		if len(updatedRules.SharingOwnerRules) != len(originalRules.SharingOwnerRules) {
			t.Error("Owner rules should be unchanged when deleting non-existent rule")
		}
	})

	t.Run("delete_with_object_prefix", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Delete with object prefix (should be stripped)
		deleteRule(sharingRulesPath, "Account.TestCriteriaRule1", "criteria")

		// Verify the rule was deleted
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Should have 1 criteria rule left
		if len(updatedRules.SharingCriteriaRules) != 1 {
			t.Errorf("Expected 1 criteria rule, got %d", len(updatedRules.SharingCriteriaRules))
		}
	})

	t.Run("delete_with_invalid_type", func(t *testing.T) {
		// Reset the test file
		err = internal.WriteToFile(testSharingRules, sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		originalRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		// Try to delete with invalid type
		deleteRule(sharingRulesPath, "TestCriteriaRule1", "invalid")

		// Verify nothing changed
		updatedRules, err := sharingrules.Open(sharingRulesPath)
		if err != nil {
			t.Fatal(err)
		}

		if len(updatedRules.SharingCriteriaRules) != len(originalRules.SharingCriteriaRules) {
			t.Error("Criteria rules should be unchanged when using invalid type")
		}

		if len(updatedRules.SharingOwnerRules) != len(originalRules.SharingOwnerRules) {
			t.Error("Owner rules should be unchanged when using invalid type")
		}
	})
}
