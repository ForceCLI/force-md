package sharingrules

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/sharingrules"
)

func TestListAllRules(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-sharingrules-test")
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
				SharedTo: sharingrules.SharedTo{
					Group: &sharingrules.Group{Text: "TestGroup"},
				},
				CriteriaItems: []sharingrules.CriteriaItem{
					{
						Field:     sharingrules.Field{Text: "CustomField__c"},
						Operation: sharingrules.Operation{Text: "equals"},
						Value:     sharingrules.Value{Text: "TestValue"},
					},
				},
			},
			{
				FullName:    "TestCriteriaRule2",
				AccessLevel: sharingrules.AccessLevel{Text: "Edit"},
				Label:       sharingrules.Label{Text: "Test Criteria Rule 2"},
				SharedTo: sharingrules.SharedTo{
					Role: &sharingrules.Role{Text: "TestRole"},
				},
				CriteriaItems: []sharingrules.CriteriaItem{
					{
						Field:     sharingrules.Field{Text: "AnotherField__c"},
						Operation: sharingrules.Operation{Text: "contains"},
						Value:     sharingrules.Value{Text: "Test"},
					},
				},
			},
		},
		SharingOwnerRules: []sharingrules.OwnerRule{
			{
				FullName:    "TestOwnerRule1",
				AccessLevel: sharingrules.AccessLevel{Text: "Read"},
				Label:       sharingrules.Label{Text: "Test Owner Rule 1"},
				SharedFrom: sharingrules.SharedFrom{
					Role: &sharingrules.Role{Text: "SourceRole"},
				},
				SharedTo: sharingrules.SharedTo{
					Group: &sharingrules.Group{Text: "TargetGroup"},
				},
			},
		},
	}

	sharingRulesPath := filepath.Join(tempDir, "Account.sharingRules")
	err = internal.WriteToFile(testSharingRules, sharingRulesPath)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("list_all_rules", func(t *testing.T) {
		// Capture output by redirecting stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		listAllRules(sharingRulesPath, "")

		w.Close()
		os.Stdout = oldStdout

		buf := make([]byte, 1024)
		n, _ := r.Read(buf)
		output := string(buf[:n])

		// Check that all rules are listed
		if !strings.Contains(output, "CRITERIA: Account.TestCriteriaRule1") {
			t.Error("Expected TestCriteriaRule1 to be listed")
		}
		if !strings.Contains(output, "CRITERIA: Account.TestCriteriaRule2") {
			t.Error("Expected TestCriteriaRule2 to be listed")
		}
		if !strings.Contains(output, "OWNER: Account.TestOwnerRule1") {
			t.Error("Expected TestOwnerRule1 to be listed")
		}
		if !strings.Contains(output, "Group(TestGroup)") {
			t.Error("Expected Group(TestGroup) to be in output")
		}
		if !strings.Contains(output, "Role(TestRole)") {
			t.Error("Expected Role(TestRole) to be in output")
		}
		if !strings.Contains(output, "CustomField__c equals TestValue") {
			t.Error("Expected criteria information to be in output")
		}
	})

	t.Run("filter_by_criteria_field", func(t *testing.T) {
		// Capture output
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		listAllRules(sharingRulesPath, "CustomField__c")

		w.Close()
		os.Stdout = oldStdout

		buf := make([]byte, 1024)
		n, _ := r.Read(buf)
		output := string(buf[:n])

		// Only TestCriteriaRule1 should be listed (it uses CustomField__c)
		if !strings.Contains(output, "CRITERIA: Account.TestCriteriaRule1") {
			t.Error("Expected TestCriteriaRule1 to be listed when filtering by CustomField__c")
		}
		if strings.Contains(output, "CRITERIA: Account.TestCriteriaRule2") {
			t.Error("TestCriteriaRule2 should not be listed when filtering by CustomField__c")
		}
		if strings.Contains(output, "OWNER: Account.TestOwnerRule1") {
			t.Error("Owner rules should not be listed when filtering by field")
		}
	})

	t.Run("filter_by_nonexistent_field", func(t *testing.T) {
		// Capture output
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		listAllRules(sharingRulesPath, "NonExistentField__c")

		w.Close()
		os.Stdout = oldStdout

		buf := make([]byte, 1024)
		n, _ := r.Read(buf)
		output := string(buf[:n])

		// No rules should be listed
		if strings.Contains(output, "CRITERIA:") {
			t.Error("No criteria rules should be listed when filtering by non-existent field")
		}
		if strings.Contains(output, "OWNER:") {
			t.Error("No owner rules should be listed when filtering by field")
		}
	})
}
