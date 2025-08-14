package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata/objects"
	"github.com/ForceCLI/force-md/metadata/objects/field"
)

func TestCopyToSourceFormat(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "sfdx")

	if err := os.MkdirAll(filepath.Join(sourceDir, "objects"), 0755); err != nil {
		t.Fatal(err)
	}

	testObject := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		Fields: []field.Field{
			{
				FullName: "TestField__c",
				Label:    &TextLiteral{Text: "Test Field"},
				Type:     &TextLiteral{Text: "Text"},
				Length:   &IntegerText{Text: "255"},
			},
		},
	}

	// In metadata format, no -meta.xml suffix
	objectPath := filepath.Join(sourceDir, "objects", "TestObject__c.object")
	if err := internal.WriteToFile(testObject, objectPath); err != nil {
		t.Fatal(err)
	}

	if err := copyMetadata(sourceDir, targetDir, "source"); err != nil {
		t.Fatal(err)
	}

	expectedFieldPath := filepath.Join(targetDir, "objects", "TestObject__c", "fields", "TestField__c.field-meta.xml")
	if _, err := os.Stat(expectedFieldPath); err != nil {
		t.Errorf("Expected field file not created: %s", expectedFieldPath)
	}

	expectedObjectPath := filepath.Join(targetDir, "objects", "TestObject__c", "TestObject__c.object-meta.xml")
	if _, err := os.Stat(expectedObjectPath); err != nil {
		t.Errorf("Expected object file not created: %s", expectedObjectPath)
	}
}

func TestCopyToMetadataFormat(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "sfdx")
	targetDir := filepath.Join(tempDir, "src")

	objectDir := filepath.Join(sourceDir, "objects", "TestObject__c")
	fieldsDir := filepath.Join(objectDir, "fields")

	if err := os.MkdirAll(fieldsDir, 0755); err != nil {
		t.Fatal(err)
	}

	emptyObject := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
	}
	objectPath := filepath.Join(objectDir, "TestObject__c.object-meta.xml")
	if err := internal.WriteToFile(emptyObject, objectPath); err != nil {
		t.Fatal(err)
	}

	testField := field.CustomField{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		Field: field.Field{
			FullName: "TestField__c",
			Label:    &TextLiteral{Text: "Test Field"},
			Type:     &TextLiteral{Text: "Text"},
			Length:   &IntegerText{Text: "255"},
		},
	}
	fieldPath := filepath.Join(fieldsDir, "TestField__c.field-meta.xml")
	if err := internal.WriteToFile(testField, fieldPath); err != nil {
		t.Fatal(err)
	}

	if err := copyMetadata(sourceDir, targetDir, "metadata"); err != nil {
		t.Fatal(err)
	}

	// In metadata format, the file should not have -meta.xml suffix
	expectedObjectPath := filepath.Join(targetDir, "objects", "TestObject__c.object")
	if _, err := os.Stat(expectedObjectPath); err != nil {
		t.Errorf("Expected merged object file not created: %s", expectedObjectPath)
	}

	data, err := os.ReadFile(expectedObjectPath)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(data), "TestField__c") {
		t.Logf("Object file contents: %s", string(data))
		t.Error("Merged object file does not contain the field")
	}
}

