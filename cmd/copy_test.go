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
	"github.com/ForceCLI/force-md/metadata/pkg"
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

	if err := CopyMetadata(sourceDir, targetDir, "source", ""); err != nil {
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

	if err := CopyMetadata(sourceDir, targetDir, "metadata", ""); err != nil {
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

func TestCopyWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "filtered")

	// Create source directory structure
	objectsDir := filepath.Join(sourceDir, "objects")
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create TestObject1 with field
	testObject1 := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		Fields: []field.Field{
			{
				FullName: "Field1__c",
				Label:    &TextLiteral{Text: "Field 1"},
				Type:     &TextLiteral{Text: "Text"},
				Length:   &IntegerText{Text: "100"},
			},
		},
	}
	object1Path := filepath.Join(objectsDir, "TestObject1__c.object")
	if err := internal.WriteToFile(testObject1, object1Path); err != nil {
		t.Fatal(err)
	}

	// Create TestObject2 with field
	testObject2 := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
		Fields: []field.Field{
			{
				FullName: "Field2__c",
				Label:    &TextLiteral{Text: "Field 2"},
				Type:     &TextLiteral{Text: "Number"},
			},
		},
	}
	object2Path := filepath.Join(objectsDir, "TestObject2__c.object")
	if err := internal.WriteToFile(testObject2, object2Path); err != nil {
		t.Fatal(err)
	}

	// Create package.xml that only includes TestObject1
	packageContent := pkg.Package{
		Xmlns:   "http://soap.sforce.com/2006/04/metadata",
		Version: "58.0",
		Types: []pkg.MetadataItems{
			{
				Name:    "CustomObject",
				Members: []pkg.Member{"TestObject1__c"},
			},
		},
	}
	packagePath := filepath.Join(tempDir, "package.xml")
	if err := internal.WriteToFile(packageContent, packagePath); err != nil {
		t.Fatal(err)
	}

	// Copy with package filter
	if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
		t.Fatal(err)
	}

	// Check that only TestObject1 was copied
	expectedObject1Path := filepath.Join(targetDir, "objects", "TestObject1__c", "TestObject1__c.object-meta.xml")
	if _, err := os.Stat(expectedObject1Path); err != nil {
		t.Errorf("Expected TestObject1 file not created: %s", expectedObject1Path)
	}

	expectedObject2Path := filepath.Join(targetDir, "objects", "TestObject2__c", "TestObject2__c.object-meta.xml")
	if _, err := os.Stat(expectedObject2Path); err == nil {
		t.Errorf("TestObject2 should not have been copied: %s", expectedObject2Path)
	}
}

func TestCopyWithPackageWildcard(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "filtered")

	// Create source directory structure
	objectsDir := filepath.Join(sourceDir, "objects")
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create multiple objects
	for i := 1; i <= 3; i++ {
		objectName := "TestObject" + string(rune('0'+i)) + "__c"
		testObject := objects.CustomObject{
			Xmlns: "http://soap.sforce.com/2006/04/metadata",
		}
		objectPath := filepath.Join(objectsDir, objectName+".object")
		if err := internal.WriteToFile(testObject, objectPath); err != nil {
			t.Fatal(err)
		}
	}

	// Create package.xml with wildcard
	packageContent := pkg.Package{
		Xmlns:   "http://soap.sforce.com/2006/04/metadata",
		Version: "58.0",
		Types: []pkg.MetadataItems{
			{
				Name:    "CustomObject",
				Members: []pkg.Member{"*"},
			},
		},
	}
	packagePath := filepath.Join(tempDir, "package.xml")
	if err := internal.WriteToFile(packageContent, packagePath); err != nil {
		t.Fatal(err)
	}

	// Copy with package filter
	if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
		t.Fatal(err)
	}

	// Check that all objects were copied
	for i := 1; i <= 3; i++ {
		objectName := "TestObject" + string(rune('0'+i)) + "__c"
		expectedPath := filepath.Join(targetDir, "objects", objectName, objectName+".object-meta.xml")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("Expected %s file not created: %s", objectName, expectedPath)
		}
	}
}

func TestCopyWithPackageWildcardAndNamed(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "filtered")

	// Create source directory structure
	objectsDir := filepath.Join(sourceDir, "objects")
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create TestObject1 only
	testObject1 := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
	}
	object1Path := filepath.Join(objectsDir, "TestObject1__c.object")
	if err := internal.WriteToFile(testObject1, object1Path); err != nil {
		t.Fatal(err)
	}

	// Create package.xml with wildcard AND a named component that doesn't exist
	packageContent := pkg.Package{
		Xmlns:   "http://soap.sforce.com/2006/04/metadata",
		Version: "58.0",
		Types: []pkg.MetadataItems{
			{
				Name: "CustomObject",
				Members: []pkg.Member{
					"*",
					"NonExistentObject__c",
				},
			},
		},
	}
	packagePath := filepath.Join(tempDir, "package.xml")
	if err := internal.WriteToFile(packageContent, packagePath); err != nil {
		t.Fatal(err)
	}

	// Copy with package filter - should fail
	err = CopyMetadata(sourceDir, targetDir, "source", packagePath)
	if err == nil {
		t.Error("Expected error for missing named component with wildcard")
	}
	if err != nil && !strings.Contains(err.Error(), "missing required components") {
		t.Errorf("Expected error about missing components, got: %v", err)
	}
}

func TestCopyFolderSubfolderWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with nested dashboard folders
	// dashboards/DashboardCBSHomepages/CBSContractingHomepage.dashboardFolder-meta.xml
	parentDir := filepath.Join(sourceDir, "dashboards", "DashboardCBSHomepages")
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a DashboardFolder metadata file in the nested folder
	folderContent := `<?xml version="1.0" encoding="UTF-8"?>
<DashboardFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>CBSContractingHomepage</name>
</DashboardFolder>`
	// In source format, this should be at dashboards/DashboardCBSHomepages/CBSContractingHomepage.dashboardFolder-meta.xml
	folderPath := filepath.Join(parentDir, "CBSContractingHomepage.dashboardFolder-meta.xml")
	if err := os.WriteFile(folderPath, []byte(folderContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create another folder in the root for comparison
	rootFolderContent := `<?xml version="1.0" encoding="UTF-8"?>
<DashboardFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>RootFolder</name>
</DashboardFolder>`
	rootFolderPath := filepath.Join(sourceDir, "dashboards", "RootFolder.dashboardFolder-meta.xml")
	if err := os.WriteFile(rootFolderPath, []byte(rootFolderContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("with_nested_folder_in_package", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "target_nested")

		// Create package.xml with nested folder reference
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "DashboardFolder",
					Members: []pkg.Member{"DashboardCBSHomepages/CBSContractingHomepage"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_nested.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDir, "metadata", packagePath); err != nil {
			t.Fatalf("Failed to copy with package filter: %v", err)
		}

		// Check that the nested folder was copied with correct structure
		expectedFolderPath := filepath.Join(targetDir, "dashboards", "DashboardCBSHomepages", "CBSContractingHomepage-meta.xml")
		if _, err := os.Stat(expectedFolderPath); err != nil {
			t.Errorf("Expected nested dashboard folder file not created at %s: %v", expectedFolderPath, err)
		}

		// Check that the root folder was NOT copied
		notExpectedPath := filepath.Join(targetDir, "dashboards", "RootFolder-meta.xml")
		if _, err := os.Stat(notExpectedPath); err == nil {
			t.Errorf("Root folder should not have been copied: %s", notExpectedPath)
		}
	})

	t.Run("with_wildcard", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "target_wildcard")

		// Create package.xml with wildcard
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "DashboardFolder",
					Members: []pkg.Member{"*"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_wildcard.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDir, "metadata", packagePath); err != nil {
			t.Fatalf("Failed to copy with package filter: %v", err)
		}

		// Check that both folders were copied with correct structure
		expectedNestedPath := filepath.Join(targetDir, "dashboards", "DashboardCBSHomepages", "CBSContractingHomepage-meta.xml")
		if _, err := os.Stat(expectedNestedPath); err != nil {
			t.Errorf("Expected nested dashboard folder file not created at %s: %v", expectedNestedPath, err)
		}

		expectedRootPath := filepath.Join(targetDir, "dashboards", "RootFolder-meta.xml")
		if _, err := os.Stat(expectedRootPath); err != nil {
			t.Errorf("Expected root dashboard folder file not created at %s: %v", expectedRootPath, err)
		}
	})
}

func TestCopyFolderSubfolderPreservation(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with nested dashboard folders
	// dashboards/ParentFolder/SubFolder/
	nestedDashboardsDir := filepath.Join(sourceDir, "dashboards", "ParentFolder", "SubFolder")
	if err := os.MkdirAll(nestedDashboardsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a DashboardFolder metadata file in the nested folder
	folderContent := `<?xml version="1.0" encoding="UTF-8"?>
<DashboardFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>SubFolder</name>
</DashboardFolder>`
	// In source format, this should be at dashboards/ParentFolder/SubFolder.dashboardFolder-meta.xml
	folderPath := filepath.Join(sourceDir, "dashboards", "ParentFolder", "SubFolder.dashboardFolder-meta.xml")
	if err := os.WriteFile(folderPath, []byte(folderContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a dashboard in the subfolder
	dashboardContent := `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata">
    <title>Nested Dashboard</title>
    <isGridLayout>false</isGridLayout>
</Dashboard>`
	dashboardPath := filepath.Join(nestedDashboardsDir, "NestedDashboard.dashboard-meta.xml")
	if err := os.WriteFile(dashboardPath, []byte(dashboardContent), 0644); err != nil {
		t.Fatal(err)
	}

	targetDir := filepath.Join(tempDir, "target")

	// Copy without package filter to test basic folder preservation
	if err := CopyMetadata(sourceDir, targetDir, "metadata", ""); err != nil {
		t.Fatalf("Failed to copy metadata: %v", err)
	}

	// Check that the folder structure is preserved
	expectedFolderPath := filepath.Join(targetDir, "dashboards", "ParentFolder", "SubFolder-meta.xml")
	if _, err := os.Stat(expectedFolderPath); err != nil {
		t.Errorf("Expected dashboard folder file not created at %s: %v", expectedFolderPath, err)
	}

	// Check that the dashboard is also in the correct location
	expectedDashboardPath := filepath.Join(targetDir, "dashboards", "ParentFolder", "SubFolder", "NestedDashboard.dashboard")
	if _, err := os.Stat(expectedDashboardPath); err != nil {
		t.Errorf("Expected dashboard file not created at %s: %v", expectedDashboardPath, err)
	}
}

func TestCopyDashboardAndFolderWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with dashboard folders
	dashboardsDir := filepath.Join(sourceDir, "dashboards", "TestFolder")
	if err := os.MkdirAll(dashboardsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a DashboardFolder metadata file
	folderContent := `<?xml version="1.0" encoding="UTF-8"?>
<DashboardFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>TestFolder</name>
</DashboardFolder>`
	// In metadata format, DashboardFolder files are in the root of dashboards dir
	folderPath := filepath.Join(sourceDir, "dashboards", "TestFolder-meta.xml")
	if err := os.WriteFile(folderPath, []byte(folderContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a dashboard in the folder
	dashboardContent := `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata">
    <title>Test Dashboard</title>
    <isGridLayout>false</isGridLayout>
</Dashboard>`
	dashboardPath := filepath.Join(dashboardsDir, "TestDashboard.dashboard")
	if err := os.WriteFile(dashboardPath, []byte(dashboardContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("both_dashboard_and_folder", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "both_filtered")

		// Create package.xml with both Dashboard and DashboardFolder
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "Dashboard",
					Members: []pkg.Member{"*"},
				},
				{
					Name:    "DashboardFolder",
					Members: []pkg.Member{"*"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "both_package.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy with package filter: %v", err)
		}

		// Check that the dashboard was copied
		expectedDashboardPath := filepath.Join(targetDir, "dashboards", "TestFolder", "TestDashboard.dashboard-meta.xml")
		if _, err := os.Stat(expectedDashboardPath); err != nil {
			t.Errorf("Expected dashboard file not created at %s: %v", expectedDashboardPath, err)
		}

		// Check that the folder was also copied
		expectedFolderPath := filepath.Join(targetDir, "dashboards", "TestFolder.dashboardFolder-meta.xml")
		if _, err := os.Stat(expectedFolderPath); err != nil {
			t.Errorf("Expected dashboard folder file not created at %s: %v", expectedFolderPath, err)
		}
	})
}

func TestCopyDashboardFolderWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with dashboard folders
	dashboardsDir := filepath.Join(sourceDir, "dashboards")
	if err := os.MkdirAll(dashboardsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a DashboardFolder metadata file
	folderContent := `<?xml version="1.0" encoding="UTF-8"?>
<DashboardFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>TestFolder</name>
</DashboardFolder>`

	// In metadata format, DashboardFolder files are named: FolderName-meta.xml
	folderPath := filepath.Join(dashboardsDir, "TestFolder-meta.xml")
	if err := os.WriteFile(folderPath, []byte(folderContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create another folder
	anotherFolderPath := filepath.Join(dashboardsDir, "AnotherFolder-meta.xml")
	anotherFolderContent := `<?xml version="1.0" encoding="UTF-8"?>
<DashboardFolder xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>AnotherFolder</name>
</DashboardFolder>`
	if err := os.WriteFile(anotherFolderPath, []byte(anotherFolderContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("dashboardfolder_wildcard", func(t *testing.T) {
		targetDirWildcard := filepath.Join(tempDir, "folder_wildcard")

		// Create package.xml with DashboardFolder wildcard
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "DashboardFolder",
					Members: []pkg.Member{"*"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "folder_package_wildcard.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDirWildcard, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy dashboard folders with package filter: %v", err)
		}

		// Check that both folders were copied
		expectedPath1 := filepath.Join(targetDirWildcard, "dashboards", "TestFolder.dashboardFolder-meta.xml")
		if _, err := os.Stat(expectedPath1); err != nil {
			t.Errorf("Expected dashboard folder file not created at %s: %v", expectedPath1, err)
		}

		expectedPath2 := filepath.Join(targetDirWildcard, "dashboards", "AnotherFolder.dashboardFolder-meta.xml")
		if _, err := os.Stat(expectedPath2); err != nil {
			t.Errorf("Expected dashboard folder file not created at %s: %v", expectedPath2, err)
		}
	})

	t.Run("dashboardfolder_specific", func(t *testing.T) {
		targetDirSpecific := filepath.Join(tempDir, "folder_specific")

		// Create package.xml with specific DashboardFolder
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "DashboardFolder",
					Members: []pkg.Member{"TestFolder"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "folder_package_specific.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDirSpecific, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy specific dashboard folder with package filter: %v", err)
		}

		// Check that only the specific folder was copied
		expectedPath := filepath.Join(targetDirSpecific, "dashboards", "TestFolder.dashboardFolder-meta.xml")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("Expected dashboard folder file not created at %s: %v", expectedPath, err)
		}

		// Check that the other folder was NOT copied
		notExpectedPath := filepath.Join(targetDirSpecific, "dashboards", "AnotherFolder.dashboardFolder-meta.xml")
		if _, err := os.Stat(notExpectedPath); err == nil {
			t.Errorf("Dashboard folder should not have been copied: %s", notExpectedPath)
		}
	})
}

func TestCopySettingsWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with settings
	settingsDir := filepath.Join(sourceDir, "settings")
	if err := os.MkdirAll(settingsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create SecuritySettings file
	securityContent := `<?xml version="1.0" encoding="UTF-8"?>
<SecuritySettings xmlns="http://soap.sforce.com/2006/04/metadata">
    <canUsersGrantLoginAccess>true</canUsersGrantLoginAccess>
    <enableAdminLoginAsAnyUser>false</enableAdminLoginAsAnyUser>
</SecuritySettings>`
	securityPath := filepath.Join(settingsDir, "Security.settings-meta.xml")
	if err := os.WriteFile(securityPath, []byte(securityContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create AccountSettings file
	accountContent := `<?xml version="1.0" encoding="UTF-8"?>
<AccountSettings xmlns="http://soap.sforce.com/2006/04/metadata">
    <enableAccountHistoryTracking>true</enableAccountHistoryTracking>
    <enableAccountTeams>true</enableAccountTeams>
</AccountSettings>`
	accountPath := filepath.Join(settingsDir, "Account.settings-meta.xml")
	if err := os.WriteFile(accountPath, []byte(accountContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create CaseSettings file
	caseContent := `<?xml version="1.0" encoding="UTF-8"?>
<CaseSettings xmlns="http://soap.sforce.com/2006/04/metadata">
    <closeCaseThroughStatusChange>true</closeCaseThroughStatusChange>
</CaseSettings>`
	casePath := filepath.Join(settingsDir, "Case.settings-meta.xml")
	if err := os.WriteFile(casePath, []byte(caseContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("settings_named_member", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "settings_named")

		// Create package.xml with named Settings member
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "Settings",
					Members: []pkg.Member{"Security"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_settings_named.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter to metadata format
		if err := CopyMetadata(sourceDir, targetDir, "metadata", packagePath); err != nil {
			t.Fatalf("Failed to copy Settings with package filter: %v", err)
		}

		// Check that SecuritySettings was copied (in metadata format without -meta.xml)
		expectedPath := filepath.Join(targetDir, "settings", "Security.settings")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("Expected SecuritySettings file not created at %s: %v", expectedPath, err)
		}

		// Check that AccountSettings was NOT copied
		notExpectedPath := filepath.Join(targetDir, "settings", "Account.settings")
		if _, err := os.Stat(notExpectedPath); err == nil {
			t.Errorf("AccountSettings should not have been copied: %s", notExpectedPath)
		}

		// Check that package.xml was copied
		expectedPackagePath := filepath.Join(targetDir, "package.xml")
		if _, err := os.Stat(expectedPackagePath); err != nil {
			t.Errorf("package.xml should have been copied to target: %s", expectedPackagePath)
		}
	})

	t.Run("settings_wildcard", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "settings_wildcard")

		// Create package.xml with Settings wildcard
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "Settings",
					Members: []pkg.Member{"*"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_settings_wildcard.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter to metadata format
		if err := CopyMetadata(sourceDir, targetDir, "metadata", packagePath); err != nil {
			t.Fatalf("Failed to copy Settings with wildcard: %v", err)
		}

		// Check that all Settings files were copied
		expectedSettings := []string{"Security", "Account", "Case"}
		for _, settingName := range expectedSettings {
			expectedPath := filepath.Join(targetDir, "settings", settingName+".settings")
			if _, err := os.Stat(expectedPath); err != nil {
				t.Errorf("Expected %s.settings file not created at %s: %v", settingName, expectedPath, err)
			}
		}
	})

	t.Run("settings_multiple_named", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "settings_multiple")

		// Create package.xml with multiple named Settings members
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name: "Settings",
					Members: []pkg.Member{
						"Security",
						"Case",
					},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_settings_multiple.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter to metadata format
		if err := CopyMetadata(sourceDir, targetDir, "metadata", packagePath); err != nil {
			t.Fatalf("Failed to copy multiple Settings: %v", err)
		}

		// Check that specified Settings were copied
		expectedSettings := []string{"Security", "Case"}
		for _, settingName := range expectedSettings {
			expectedPath := filepath.Join(targetDir, "settings", settingName+".settings")
			if _, err := os.Stat(expectedPath); err != nil {
				t.Errorf("Expected %s.settings file not created at %s: %v", settingName, expectedPath, err)
			}
		}

		// Check that AccountSettings was NOT copied
		notExpectedPath := filepath.Join(targetDir, "settings", "Account.settings")
		if _, err := os.Stat(notExpectedPath); err == nil {
			t.Errorf("AccountSettings should not have been copied: %s", notExpectedPath)
		}
	})
}

func TestCopyPackageXmlAndDestructiveChanges(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")
	objectsDir := filepath.Join(sourceDir, "objects")
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a simple object
	testObject := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
	}
	objectPath := filepath.Join(objectsDir, "TestObject__c.object")
	if err := internal.WriteToFile(testObject, objectPath); err != nil {
		t.Fatal(err)
	}

	// Create package.xml
	packageContent := pkg.Package{
		Xmlns:   "http://soap.sforce.com/2006/04/metadata",
		Version: "58.0",
		Types: []pkg.MetadataItems{
			{
				Name:    "CustomObject",
				Members: []pkg.Member{"TestObject__c"},
			},
		},
	}
	packagePath := filepath.Join(tempDir, "package.xml")
	if err := internal.WriteToFile(packageContent, packagePath); err != nil {
		t.Fatal(err)
	}

	// Create destructiveChanges.xml
	destructiveContent := `<?xml version="1.0" encoding="UTF-8"?>
<Package xmlns="http://soap.sforce.com/2006/04/metadata">
    <types>
        <members>OldObject__c</members>
        <name>CustomObject</name>
    </types>
    <version>58.0</version>
</Package>`
	destructivePath := filepath.Join(tempDir, "destructiveChanges.xml")
	if err := os.WriteFile(destructivePath, []byte(destructiveContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create destructiveChangesPost.xml
	destructivePostContent := `<?xml version="1.0" encoding="UTF-8"?>
<Package xmlns="http://soap.sforce.com/2006/04/metadata">
    <types>
        <members>OldField__c</members>
        <name>CustomField</name>
    </types>
    <version>58.0</version>
</Package>`
	destructivePostPath := filepath.Join(tempDir, "destructiveChangesPost.xml")
	if err := os.WriteFile(destructivePostPath, []byte(destructivePostContent), 0644); err != nil {
		t.Fatal(err)
	}

	targetDir := filepath.Join(tempDir, "target")

	// Copy with package filter
	if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
		t.Fatalf("Failed to copy with package filter: %v", err)
	}

	// Check that package.xml was copied
	expectedPackagePath := filepath.Join(targetDir, "package.xml")
	if _, err := os.Stat(expectedPackagePath); err != nil {
		t.Errorf("package.xml should have been copied to target: %s", expectedPackagePath)
	}

	// Check that destructiveChanges.xml was copied
	expectedDestructivePath := filepath.Join(targetDir, "destructiveChanges.xml")
	if _, err := os.Stat(expectedDestructivePath); err != nil {
		t.Errorf("destructiveChanges.xml should have been copied to target: %s", expectedDestructivePath)
	}

	// Check that destructiveChangesPost.xml was copied
	expectedDestructivePostPath := filepath.Join(targetDir, "destructiveChangesPost.xml")
	if _, err := os.Stat(expectedDestructivePostPath); err != nil {
		t.Errorf("destructiveChangesPost.xml should have been copied to target: %s", expectedDestructivePostPath)
	}

	// Verify the content of copied package.xml
	copiedContent, err := os.ReadFile(expectedPackagePath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(copiedContent), "TestObject__c") {
		t.Error("Copied package.xml does not contain expected content")
	}
}

func TestCopyTargetDirectoryCreation(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")
	objectsDir := filepath.Join(sourceDir, "objects")
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a simple object
	testObject := objects.CustomObject{
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
	}
	objectPath := filepath.Join(objectsDir, "TestObject__c.object")
	if err := internal.WriteToFile(testObject, objectPath); err != nil {
		t.Fatal(err)
	}

	t.Run("creates_non_existent_target", func(t *testing.T) {
		// Use a target directory that doesn't exist
		targetDir := filepath.Join(tempDir, "non", "existent", "nested", "target")

		// Ensure it doesn't exist
		if _, err := os.Stat(targetDir); err == nil {
			t.Fatal("Target directory should not exist before test")
		}

		// Copy without package filter
		if err := CopyMetadata(sourceDir, targetDir, "source", ""); err != nil {
			t.Fatalf("Failed to copy metadata: %v", err)
		}

		// Check that target directory was created
		if _, err := os.Stat(targetDir); err != nil {
			t.Errorf("Target directory should have been created: %v", err)
		}

		// Check that metadata was copied
		expectedPath := filepath.Join(targetDir, "objects", "TestObject__c", "TestObject__c.object-meta.xml")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("Expected object file not created: %s", expectedPath)
		}
	})

	t.Run("creates_target_with_package_filter", func(t *testing.T) {
		// Create package.xml
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "CustomObject",
					Members: []pkg.Member{"TestObject__c"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Use a target directory that doesn't exist
		targetDir := filepath.Join(tempDir, "another", "nested", "target")

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy with package filter: %v", err)
		}

		// Check that package.xml was copied
		expectedPackagePath := filepath.Join(targetDir, "package.xml")
		if _, err := os.Stat(expectedPackagePath); err != nil {
			t.Errorf("package.xml should have been copied to target: %s", expectedPackagePath)
		}
	})
}

func TestCopyDocumentWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with documents
	documentsDir := filepath.Join(sourceDir, "documents", "Images")
	if err := os.MkdirAll(documentsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a Document metadata file
	documentContent := `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="http://soap.sforce.com/2006/04/metadata">
    <internalUseOnly>false</internalUseOnly>
    <name>BromleyBrookPicture</name>
    <public>true</public>
    <contentType>image/bmp</contentType>
</Document>`
	documentMetaPath := filepath.Join(documentsDir, "BromleyBrookPicture.document-meta.xml")
	if err := os.WriteFile(documentMetaPath, []byte(documentContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create the actual document file (BMP image)
	documentFilePath := filepath.Join(documentsDir, "BromleyBrookPicture.bmp")
	if err := os.WriteFile(documentFilePath, []byte("fake bmp content"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create another document with different extension
	pdfDocContent := `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="http://soap.sforce.com/2006/04/metadata">
    <internalUseOnly>false</internalUseOnly>
    <name>Report</name>
    <public>true</public>
    <contentType>application/pdf</contentType>
</Document>`
	pdfMetaPath := filepath.Join(documentsDir, "Report.document-meta.xml")
	if err := os.WriteFile(pdfMetaPath, []byte(pdfDocContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create the actual PDF file
	pdfFilePath := filepath.Join(documentsDir, "Report.pdf")
	if err := os.WriteFile(pdfFilePath, []byte("fake pdf content"), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("document_with_extension_in_package", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "target_doc")

		// Create package.xml with Document reference including actual file extension
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name: "Document",
					Members: []pkg.Member{
						"Images/BromleyBrookPicture.bmp",
						"Images/Report.pdf",
					},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_doc.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter to source format (Document metadata format conversion is not fully implemented)
		if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy documents with package filter: %v", err)
		}

		// Check that the document metadata files were copied
		expectedBmpMetaPath := filepath.Join(targetDir, "documents", "Images", "BromleyBrookPicture.document-meta.xml")
		if _, err := os.Stat(expectedBmpMetaPath); err != nil {
			t.Errorf("Expected BMP document metadata file not created at %s: %v", expectedBmpMetaPath, err)
		}

		expectedPdfMetaPath := filepath.Join(targetDir, "documents", "Images", "Report.document-meta.xml")
		if _, err := os.Stat(expectedPdfMetaPath); err != nil {
			t.Errorf("Expected PDF document metadata file not created at %s: %v", expectedPdfMetaPath, err)
		}
	})

	t.Run("document_wildcard", func(t *testing.T) {
		targetDir := filepath.Join(tempDir, "target_doc_wildcard")

		// Create package.xml with Document wildcard
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "Document",
					Members: []pkg.Member{"*"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_doc_wildcard.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy documents with wildcard: %v", err)
		}

		// Check that both documents were copied to source format
		expectedBmpMetaPath := filepath.Join(targetDir, "documents", "Images", "BromleyBrookPicture.document-meta.xml")
		if _, err := os.Stat(expectedBmpMetaPath); err != nil {
			t.Errorf("Expected BMP document metadata file not created at %s: %v", expectedBmpMetaPath, err)
		}

		expectedPdfMetaPath := filepath.Join(targetDir, "documents", "Images", "Report.document-meta.xml")
		if _, err := os.Stat(expectedPdfMetaPath); err != nil {
			t.Errorf("Expected PDF document metadata file not created at %s: %v", expectedPdfMetaPath, err)
		}
	})
}

func TestCopyDashboardWithPackageFilter(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "force-md-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir := filepath.Join(tempDir, "src")

	// Create source directory structure with dashboard folders
	dashboardsDir := filepath.Join(sourceDir, "dashboards", "TestFolder")
	if err := os.MkdirAll(dashboardsDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a test dashboard file in metadata format
	dashboardContent := `<?xml version="1.0" encoding="UTF-8"?>
<Dashboard xmlns="http://soap.sforce.com/2006/04/metadata">
    <title>Test Dashboard</title>
    <isGridLayout>false</isGridLayout>
    <textColor>#000000</textColor>
    <titleColor>#000000</titleColor>
    <titleSize>12</titleSize>
</Dashboard>`
	dashboardPath := filepath.Join(dashboardsDir, "TestDashboard.dashboard")
	if err := os.WriteFile(dashboardPath, []byte(dashboardContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Also create a dashboard in another folder to test filtering
	anotherDir := filepath.Join(sourceDir, "dashboards", "AnotherFolder")
	if err := os.MkdirAll(anotherDir, 0755); err != nil {
		t.Fatal(err)
	}
	anotherDashboardPath := filepath.Join(anotherDir, "AnotherDashboard.dashboard")
	if err := os.WriteFile(anotherDashboardPath, []byte(dashboardContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("wildcard_selection", func(t *testing.T) {
		targetDirWildcard := filepath.Join(tempDir, "filtered_wildcard")

		// Create package.xml with Dashboard wildcard
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "Dashboard",
					Members: []pkg.Member{"*"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_wildcard.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDirWildcard, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy dashboard with package filter: %v", err)
		}

		// Check that both dashboards were copied with correct folder structure
		expectedPath1 := filepath.Join(targetDirWildcard, "dashboards", "TestFolder", "TestDashboard.dashboard-meta.xml")
		if _, err := os.Stat(expectedPath1); err != nil {
			t.Errorf("Expected dashboard file not created at %s: %v", expectedPath1, err)
		}

		expectedPath2 := filepath.Join(targetDirWildcard, "dashboards", "AnotherFolder", "AnotherDashboard.dashboard-meta.xml")
		if _, err := os.Stat(expectedPath2); err != nil {
			t.Errorf("Expected dashboard file not created at %s: %v", expectedPath2, err)
		}
	})

	t.Run("specific_foldered_item", func(t *testing.T) {
		targetDirSpecific := filepath.Join(tempDir, "filtered_specific")

		// Create package.xml with specific dashboard reference (folder/name format)
		packageContent := pkg.Package{
			Xmlns:   "http://soap.sforce.com/2006/04/metadata",
			Version: "58.0",
			Types: []pkg.MetadataItems{
				{
					Name:    "Dashboard",
					Members: []pkg.Member{"TestFolder/TestDashboard"},
				},
			},
		}
		packagePath := filepath.Join(tempDir, "package_specific.xml")
		if err := internal.WriteToFile(packageContent, packagePath); err != nil {
			t.Fatal(err)
		}

		// Copy with package filter
		if err := CopyMetadata(sourceDir, targetDirSpecific, "source", packagePath); err != nil {
			t.Fatalf("Failed to copy specific dashboard with package filter: %v", err)
		}

		// Check that only the specific dashboard was copied
		expectedPath := filepath.Join(targetDirSpecific, "dashboards", "TestFolder", "TestDashboard.dashboard-meta.xml")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("Expected dashboard file not created at %s: %v", expectedPath, err)
		}

		// Check that the other dashboard was NOT copied
		notExpectedPath := filepath.Join(targetDirSpecific, "dashboards", "AnotherFolder", "AnotherDashboard.dashboard-meta.xml")
		if _, err := os.Stat(notExpectedPath); err == nil {
			t.Errorf("Dashboard should not have been copied: %s", notExpectedPath)
		}
	})
}

func TestCopyToMetadataFormatWritesObjectsThatOnlyHaveFields(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "sfdx")
	targetDir := filepath.Join(tempDir, "src")

	// Standard objects in source format often have custom fields but no
	// object-meta.xml file of their own.
	for _, objectName := range []string{"Account", "Contact"} {
		fieldsDir := filepath.Join(sourceDir, "objects", objectName, "fields")
		if err := os.MkdirAll(fieldsDir, 0755); err != nil {
			t.Fatal(err)
		}
		testField := field.CustomField{
			Xmlns: "http://soap.sforce.com/2006/04/metadata",
			Field: field.Field{
				FullName: objectName + "Field__c",
				Label:    &TextLiteral{Text: "Test Field"},
				Type:     &TextLiteral{Text: "Text"},
				Length:   &IntegerText{Text: "255"},
			},
		}
		fieldPath := filepath.Join(fieldsDir, objectName+"Field__c.field-meta.xml")
		if err := internal.WriteToFile(testField, fieldPath); err != nil {
			t.Fatal(err)
		}
	}

	if err := CopyMetadata(sourceDir, targetDir, "metadata", ""); err != nil {
		t.Fatal(err)
	}

	for _, objectName := range []string{"Account", "Contact"} {
		objectPath := filepath.Join(targetDir, "objects", objectName+".object")
		data, err := os.ReadFile(objectPath)
		if err != nil {
			t.Fatalf("expected %s to be written: %v", objectPath, err)
		}
		if !strings.Contains(string(data), objectName+"Field__c") {
			t.Errorf("%s does not contain %sField__c:\n%s", objectPath, objectName, data)
		}
	}
	if _, err := os.Stat(filepath.Join(targetDir, "objects", ".object")); err == nil {
		t.Error("an object with an empty name was written")
	}
}

func TestCopyCallCenterToSourceFormat(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "sfdx")

	if err := os.MkdirAll(filepath.Join(sourceDir, "callCenters"), 0755); err != nil {
		t.Fatal(err)
	}
	callCenter := `<?xml version="1.0" encoding="UTF-8"?>
<CallCenter xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <adapterUrl>https://example.com/adapter</adapterUrl>
    <displayName>Example Phone</displayName>
    <displayNameLabel>Display Name</displayNameLabel>
    <internalNameLabel>Internal Name</internalNameLabel>
    <sections>
        <items>
            <label>Description</label>
            <name>reqDescription</name>
            <value xsi:nil="true"/>
        </items>
        <items>
            <label>CTI Adapter URL</label>
            <name>reqAdapterUrl</name>
            <value>https://example.com/adapter</value>
        </items>
        <label>General Info</label>
        <name>reqGeneralInfo</name>
    </sections>
</CallCenter>
`
	if err := os.WriteFile(filepath.Join(sourceDir, "callCenters", "ExamplePhone.callCenter"), []byte(callCenter), 0644); err != nil {
		t.Fatal(err)
	}

	if err := CopyMetadata(sourceDir, targetDir, "source", ""); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(filepath.Join(targetDir, "callCenters", "ExamplePhone.callCenter-meta.xml"))
	if err != nil {
		t.Fatalf("Expected call center file not created: %v", err)
	}
	if string(got) != callCenter {
		t.Errorf("copied call center differs from source:\n%s", got)
	}
}

// territoryFixture maps each metadata-format path of a territory model to its
// contents. Both models define a West territory and a West rule, and the US
// model and its West territory carry custom field values. The US West
// territory also sets its access to leads, and the US model sets its record
// visibility and user record access. The Canada West territory sets no
// access levels, and the Canada West rule has an item with no value.
var territoryFixture = map[string]string{
	"territory2Types/Region.territory2Type": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Type xmlns="http://soap.sforce.com/2006/04/metadata">
    <description>Sales regions</description>
    <name>Region</name>
    <priority>1</priority>
</Territory2Type>
`,
	"territory2Models/US/US.territory2Model": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Model xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <customFields>
        <name>Activation_Date__c</name>
        <value xsi:type="xsd:date">2026-01-01</value>
    </customFields>
    <description>United States</description>
    <name>US</name>
    <recordVisibility>Private</recordVisibility>
    <userRecordAccess>Read</userRecordAccess>
</Territory2Model>
`,
	"territory2Models/US/territories/All_US.territory2": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata">
    <accountAccessLevel>Read</accountAccessLevel>
    <caseAccessLevel>None</caseAccessLevel>
    <contactAccessLevel>Edit</contactAccessLevel>
    <name>All US</name>
    <opportunityAccessLevel>None</opportunityAccessLevel>
    <territory2Type>Region</territory2Type>
</Territory2>
`,
	"territory2Models/US/territories/West.territory2": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <accountAccessLevel>Edit</accountAccessLevel>
    <caseAccessLevel>Read</caseAccessLevel>
    <contactAccessLevel>Edit</contactAccessLevel>
    <customFields>
        <name>Region_Code__c</name>
        <value xsi:type="xsd:string">W &amp; NW</value>
    </customFields>
    <customFields>
        <name>Manager_Phone__c</name>
        <value xsi:nil="true"/>
    </customFields>
    <description>Western states</description>
    <name>West</name>
    <objectAccessLevels>
        <accessLevel>All</accessLevel>
        <objectType>Lead</objectType>
    </objectAccessLevels>
    <opportunityAccessLevel>Read</opportunityAccessLevel>
    <parentTerritory>All_US</parentTerritory>
    <ruleAssociations>
        <inherited>false</inherited>
        <ruleName>West</ruleName>
    </ruleAssociations>
    <territory2Type>Region</territory2Type>
</Territory2>
`,
	"territory2Models/US/rules/West.territory2Rule": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Rule xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <booleanFilter>1 OR 2</booleanFilter>
    <name>West</name>
    <objectType>Account</objectType>
    <ruleItems>
        <field>Account.BillingState</field>
        <operation>equals</operation>
        <value>CA</value>
    </ruleItems>
    <ruleItems>
        <field>Account.BillingState</field>
        <operation>equals</operation>
        <value>OR</value>
    </ruleItems>
</Territory2Rule>
`,
	"territory2Models/Canada/Canada.territory2Model": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Model xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>Canada</name>
</Territory2Model>
`,
	"territory2Models/Canada/territories/West.territory2": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2 xmlns="http://soap.sforce.com/2006/04/metadata">
    <name>West</name>
    <territory2Type>Region</territory2Type>
</Territory2>
`,
	"territory2Models/Canada/rules/West.territory2Rule": `<?xml version="1.0" encoding="UTF-8"?>
<Territory2Rule xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>false</active>
    <name>West</name>
    <objectType>Account</objectType>
    <ruleItems>
        <field>Account.BillingState</field>
        <operation>equals</operation>
        <value>BC</value>
    </ruleItems>
    <ruleItems>
        <field>Account.BillingCity</field>
        <operation>equals</operation>
    </ruleItems>
</Territory2Rule>
`,
}

func writeTerritoryFixture(t *testing.T, dir string) {
	t.Helper()
	for relativePath, content := range territoryFixture {
		path := filepath.Join(dir, relativePath)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
}

// assertTerritoryFiles checks that dir holds exactly the fixture's files, each
// unchanged, with suffix appended to every path.
func assertTerritoryFiles(t *testing.T, dir, suffix string) {
	t.Helper()
	found := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		found++
		relativePath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		want, ok := territoryFixture[strings.TrimSuffix(filepath.ToSlash(relativePath), suffix)]
		if !ok {
			t.Errorf("unexpected file %s", relativePath)
			return nil
		}
		got, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if string(got) != want {
			t.Errorf("%s differs from source:\n%s", relativePath, got)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if found != len(territoryFixture) {
		t.Errorf("wrote %d files, want %d", found, len(territoryFixture))
	}
}

func TestCopyTerritoryModelsBetweenFormats(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	sfdxDir := filepath.Join(tempDir, "sfdx")
	mdapiDir := filepath.Join(tempDir, "mdapi")
	writeTerritoryFixture(t, sourceDir)

	if err := CopyMetadata(sourceDir, sfdxDir, "source", ""); err != nil {
		t.Fatal(err)
	}
	assertTerritoryFiles(t, sfdxDir, "-meta.xml")

	if err := CopyMetadata(sfdxDir, mdapiDir, "metadata", ""); err != nil {
		t.Fatal(err)
	}
	assertTerritoryFiles(t, mdapiDir, "")
}

func TestCopyTerritoriesByModelQualifiedPackageMembers(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "sfdx")
	writeTerritoryFixture(t, sourceDir)

	packageXML := `<?xml version="1.0" encoding="UTF-8"?>
<Package xmlns="http://soap.sforce.com/2006/04/metadata">
    <types>
        <members>US.West</members>
        <name>Territory2</name>
    </types>
    <types>
        <members>Canada.West</members>
        <name>Territory2Rule</name>
    </types>
    <version>67.0</version>
</Package>
`
	packagePath := filepath.Join(tempDir, "package.xml")
	if err := os.WriteFile(packagePath, []byte(packageXML), 0644); err != nil {
		t.Fatal(err)
	}

	if err := CopyMetadata(sourceDir, targetDir, "source", packagePath); err != nil {
		t.Fatal(err)
	}

	var got []string
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		relativePath, err := filepath.Rel(targetDir, path)
		got = append(got, filepath.ToSlash(relativePath))
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"package.xml",
		"territory2Models/Canada/rules/West.territory2Rule-meta.xml",
		"territory2Models/US/territories/West.territory2-meta.xml",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("copied files:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func TestCopyServiceChannelToSourceFormat(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "sfdx")

	if err := os.MkdirAll(filepath.Join(sourceDir, "serviceChannels"), 0755); err != nil {
		t.Fatal(err)
	}
	serviceChannel := `<?xml version="1.0" encoding="UTF-8"?>
<ServiceChannel xmlns="http://soap.sforce.com/2006/04/metadata">
    <afterConvoWorkMaxTime>90</afterConvoWorkMaxTime>
    <capacityModel>TAB_BASED</capacityModel>
    <doesMinimizeWidgetOnAccept>true</doesMinimizeWidgetOnAccept>
    <hasAcwExtensionEnabled>false</hasAcwExtensionEnabled>
    <hasAfterConvoWorkTimer>true</hasAfterConvoWorkTimer>
    <hasAutoAcceptEnabled>false</hasAutoAcceptEnabled>
    <isInterruptible>false</isInterruptible>
    <label>Phone</label>
    <relatedEntityType>VoiceCall</relatedEntityType>
    <serviceChannelStatusFieldMappings>
        <type>COMPLETED</type>
        <value>Closed</value>
    </serviceChannelStatusFieldMappings>
    <statusField>Status</statusField>
</ServiceChannel>
`
	if err := os.WriteFile(filepath.Join(sourceDir, "serviceChannels", "sfdc_phone.serviceChannel"), []byte(serviceChannel), 0644); err != nil {
		t.Fatal(err)
	}

	if err := CopyMetadata(sourceDir, targetDir, "source", ""); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(filepath.Join(targetDir, "serviceChannels", "sfdc_phone.serviceChannel-meta.xml"))
	if err != nil {
		t.Fatalf("Expected service channel file not created: %v", err)
	}
	if string(got) != serviceChannel {
		t.Errorf("copied service channel differs from source:\n%s", got)
	}
}

func TestCopySkillToSourceFormat(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "sfdx")

	if err := os.MkdirAll(filepath.Join(sourceDir, "skills"), 0755); err != nil {
		t.Fatal(err)
	}
	skill := `<?xml version="1.0" encoding="UTF-8"?>
<Skill xmlns="http://soap.sforce.com/2006/04/metadata">
    <assignments>
        <profiles>
            <profile>admin</profile>
        </profiles>
        <users>
            <user>first@example.com</user>
        </users>
    </assignments>
    <description>Routes support cases</description>
    <label>Support Team</label>
</Skill>
`
	if err := os.WriteFile(filepath.Join(sourceDir, "skills", "Support_Team.skill"), []byte(skill), 0644); err != nil {
		t.Fatal(err)
	}

	if err := CopyMetadata(sourceDir, targetDir, "source", ""); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(filepath.Join(targetDir, "skills", "Support_Team.skill-meta.xml"))
	if err != nil {
		t.Fatalf("Expected skill file not created: %v", err)
	}
	if string(got) != skill {
		t.Errorf("copied skill differs from source:\n%s", got)
	}
}

func TestCopyServicePresenceStatusToSourceFormat(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "src")
	targetDir := filepath.Join(tempDir, "sfdx")

	if err := os.MkdirAll(filepath.Join(sourceDir, "servicePresenceStatuses"), 0755); err != nil {
		t.Fatal(err)
	}
	statuses := map[string]string{
		"Available": `<?xml version="1.0" encoding="UTF-8"?>
<ServicePresenceStatus xmlns="http://soap.sforce.com/2006/04/metadata">
    <channels>
        <channel>Cases</channel>
        <channel>sfdc_phone</channel>
    </channels>
    <label>Available</label>
</ServicePresenceStatus>
`,
		"Busy": `<?xml version="1.0" encoding="UTF-8"?>
<ServicePresenceStatus xmlns="http://soap.sforce.com/2006/04/metadata">
    <label>Busy</label>
</ServicePresenceStatus>
`,
	}
	for name, content := range statuses {
		if err := os.WriteFile(filepath.Join(sourceDir, "servicePresenceStatuses", name+".servicePresenceStatus"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	if err := CopyMetadata(sourceDir, targetDir, "source", ""); err != nil {
		t.Fatal(err)
	}

	for name, content := range statuses {
		got, err := os.ReadFile(filepath.Join(targetDir, "servicePresenceStatuses", name+".servicePresenceStatus-meta.xml"))
		if err != nil {
			t.Fatalf("Expected service presence status file not created: %v", err)
		}
		if string(got) != content {
			t.Errorf("copied service presence status %s differs from source:\n%s", name, got)
		}
	}
}
