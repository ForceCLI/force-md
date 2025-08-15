package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/repo"
)

func TestCopyObjectTranslationToMetadataFormat(t *testing.T) {
	// Create a temporary source directory with CustomObjectTranslation in source format
	sourceDir := t.TempDir()
	targetDir := t.TempDir()

	// Create objectTranslations directory structure (only fields subdirectory needed)
	objTransDir := filepath.Join(sourceDir, "objectTranslations", "Account-en_US")
	os.MkdirAll(objTransDir, 0755)

	// Create base object translation file with recordTypes and validationRules inline (SFDX convention)
	baseContent := `<?xml version="1.0" encoding="UTF-8"?>
<CustomObjectTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <recordTypes>
        <description/>
        <label>Business Account</label>
        <name>Business</name>
    </recordTypes>
    <validationRules>
        <errorMessage>Amount must be positive</errorMessage>
        <name>Amount_Must_Be_Positive</name>
    </validationRules>
</CustomObjectTranslation>`
	os.WriteFile(filepath.Join(objTransDir, "Account-en_US.objectTranslation-meta.xml"), []byte(baseContent), 0644)

	// Create a field translation (fields are the only ones that get decomposed)
	fieldContent := `<?xml version="1.0" encoding="UTF-8"?>
<CustomFieldTranslation xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Test Field</label>
    <name>TestField__c</name>
</CustomFieldTranslation>`
	os.WriteFile(filepath.Join(objTransDir, "TestField__c.fieldTranslation-meta.xml"), []byte(fieldContent), 0644)

	// First, let's check what metadata files we created in source
	t.Log("Source files created:")
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			t.Logf("  %s", path)
		}
		return nil
	})

	// Create a repo and check what it finds
	testRepo := repo.NewRepo()
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			_, err := testRepo.Open(path)
			if err != nil {
				t.Logf("Failed to open %s: %v", path, err)
			} else {
				t.Logf("Successfully opened %s", path)
			}
		}
		return nil
	})

	// Check what's in the repo - should only have CustomObjectTranslation and CustomFieldTranslation
	t.Logf("Repo types: %v", testRepo.Types())
	for _, typeName := range testRepo.Types() {
		items := testRepo.Items(typeName)
		t.Logf("  %s: %d items", typeName, len(items))
		for _, item := range items {
			if reg, ok := item.(metadata.RegisterableMetadata); ok {
				t.Logf("    - Name: %s", reg.GetMetadataInfo().Name())
			}
		}
	}

	// Run the copy command
	err = CopyMetadata(sourceDir, targetDir, "metadata", "")
	if err != nil {
		t.Fatalf("Failed to copy metadata: %v", err)
	}

	// Also check what files were created
	err = filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			t.Logf("Created file: %s", path)
		}
		return nil
	})
	if err != nil {
		t.Logf("Error walking target dir: %v", err)
	}

	// Read the resulting merged file
	mergedPath := filepath.Join(targetDir, "objectTranslations", "Account-en_US.objectTranslation")
	data, err := os.ReadFile(mergedPath)
	if err != nil {
		t.Fatalf("Failed to read merged file: %v", err)
	}

	content := string(data)

	// Debug: print the merged content
	t.Logf("Merged content:\n%s", content)

	// Check that the field translation is included
	if !strings.Contains(content, "TestField__c") {
		t.Error("Merged object translation file does not contain the field translation")
	}

	// Check that the record type translation is included
	if !strings.Contains(content, "Business") {
		t.Error("Merged object translation file does not contain the record type translation")
	}
	if !strings.Contains(content, "Business Account") {
		t.Error("Merged object translation file does not contain the record type label")
	}

	// Check that the validation rule translation is included
	if !strings.Contains(content, "Amount_Must_Be_Positive") {
		t.Error("Merged object translation file does not contain the validation rule translation")
	}
	if !strings.Contains(content, "Amount must be positive") {
		t.Error("Merged object translation file does not contain the validation rule error message")
	}
}
