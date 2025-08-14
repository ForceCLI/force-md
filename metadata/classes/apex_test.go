package apexClass

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApexClassOpen(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Create a test metadata file
	metaPath := filepath.Join(tmpDir, "MyClass.cls-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>58.0</apiVersion>
    <status>Active</status>
</ApexClass>`

	err := os.WriteFile(metaPath, []byte(metaContent), 0644)
	require.NoError(t, err)

	// Test Open function
	apexClass, err := Open(metaPath)
	require.NoError(t, err)
	assert.NotNil(t, apexClass)
	assert.Equal(t, metaPath, apexClass.SourcePath)
	assert.Equal(t, "58.0", apexClass.ApiVersion.Text)
	assert.Equal(t, "Active", apexClass.Status.Text)
	assert.Nil(t, apexClass.SourceCode) // Should not load code in Open
}

func TestApexClassFiles(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		setupFiles  map[string]string
		sourcePath  string
		format      metadata.Format
		expectFiles map[string]bool // map of file path patterns to whether they should exist
		expectCode  bool
	}{
		{
			name: "Source format with code",
			setupFiles: map[string]string{
				"MyClass.cls":          "public class MyClass {}",
				"MyClass.cls-meta.xml": `<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata"><apiVersion>58.0</apiVersion></ApexClass>`,
			},
			sourcePath: filepath.Join(tmpDir, "MyClass.cls-meta.xml"),
			format:     metadata.SourceFormat,
			expectFiles: map[string]bool{
				"classes/MyClass.cls-meta.xml": true,
				"classes/MyClass.cls":          true,
			},
			expectCode: true,
		},
		{
			name: "Source format without code",
			setupFiles: map[string]string{
				"MyClass.cls-meta.xml": `<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata"><apiVersion>58.0</apiVersion></ApexClass>`,
			},
			sourcePath: filepath.Join(tmpDir, "MyClass.cls-meta.xml"),
			format:     metadata.SourceFormat,
			expectFiles: map[string]bool{
				"classes/MyClass.cls-meta.xml": true,
				"classes/MyClass.cls":          false,
			},
			expectCode: false,
		},
		{
			name: "Metadata format with code",
			setupFiles: map[string]string{
				"MyClass.cls":          "public class MyClass {}",
				"MyClass.cls-meta.xml": `<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata"><apiVersion>58.0</apiVersion></ApexClass>`,
			},
			sourcePath: filepath.Join(tmpDir, "MyClass.cls-meta.xml"),
			format:     metadata.MetadataFormat,
			expectFiles: map[string]bool{
				"classes/MyClass.cls-meta.xml": true,
				"classes/MyClass.cls":          true,
			},
			expectCode: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test files
			for filename, content := range tt.setupFiles {
				filePath := filepath.Join(tmpDir, filename)
				err := os.WriteFile(filePath, []byte(content), 0644)
				require.NoError(t, err)
			}

			// Create ApexClass instance using Open to properly set metadata
			apexClass, err := Open(tt.sourcePath)
			require.NoError(t, err)

			// Test Files method
			files, err := apexClass.Files(tt.format)
			require.NoError(t, err)

			// Check expected files
			for path, shouldExist := range tt.expectFiles {
				if shouldExist {
					assert.Contains(t, files, path)
				} else {
					assert.NotContains(t, files, path)
				}
			}

			// Check if code was loaded
			if tt.expectCode {
				assert.NotNil(t, apexClass.SourceCode)
				assert.Equal(t, "public class MyClass {}", string(apexClass.SourceCode))
			}

			// Cleanup test files
			for filename := range tt.setupFiles {
				os.Remove(filepath.Join(tmpDir, filename))
			}
		})
	}
}

func TestApexClassFilesUnsupportedFormat(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	metaPath := filepath.Join(tmpDir, "MyClass.cls-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<ApexClass xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>58.0</apiVersion>
</ApexClass>`
	err := os.WriteFile(metaPath, []byte(metaContent), 0644)
	require.NoError(t, err)

	apexClass, err := Open(metaPath)
	require.NoError(t, err)

	// Test with an invalid format
	_, err = apexClass.Files(metadata.Format("invalid"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported format")
}
