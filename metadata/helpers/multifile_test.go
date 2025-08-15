package helpers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadCompanionFile(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	tests := []struct {
		name          string
		setupFiles    map[string]string
		sourcePath    string
		metaSuffix    string
		fileExt       string
		expectNil     bool
		expectContent string
	}{
		{
			name: "Source format - load companion file",
			setupFiles: map[string]string{
				"MyClass.cls":          "public class MyClass {}",
				"MyClass.cls-meta.xml": "<ApexClass/>",
			},
			sourcePath:    filepath.Join(tmpDir, "MyClass.cls-meta.xml"),
			metaSuffix:    "-meta.xml",
			fileExt:       ".cls",
			expectContent: "public class MyClass {}",
		},
		{
			name: "Metadata format - .cls.xml to .cls",
			setupFiles: map[string]string{
				"MyClass.cls":     "public class MyClass {}",
				"MyClass.cls.xml": "<ApexClass/>",
			},
			sourcePath:    filepath.Join(tmpDir, "MyClass.cls.xml"),
			metaSuffix:    "-meta.xml",
			fileExt:       ".cls",
			expectContent: "public class MyClass {}",
		},
		{
			name: "Metadata format - plain .xml to .cls",
			setupFiles: map[string]string{
				"MyClass.cls": "public class MyClass {}",
				"MyClass.xml": "<ApexClass/>",
			},
			sourcePath:    filepath.Join(tmpDir, "MyClass.xml"),
			metaSuffix:    "-meta.xml",
			fileExt:       ".cls",
			expectContent: "public class MyClass {}",
		},
		{
			name: "Companion file doesn't exist",
			setupFiles: map[string]string{
				"MyClass.cls-meta.xml": "<ApexClass/>",
			},
			sourcePath: filepath.Join(tmpDir, "MyClass.cls-meta.xml"),
			metaSuffix: "-meta.xml",
			fileExt:    ".cls",
			expectNil:  true,
		},
		{
			name:       "Empty source path",
			sourcePath: "",
			metaSuffix: "-meta.xml",
			fileExt:    ".cls",
			expectNil:  true,
		},
		{
			name: "Trigger file",
			setupFiles: map[string]string{
				"MyTrigger.trigger":          "trigger MyTrigger on Account {}",
				"MyTrigger.trigger-meta.xml": "<ApexTrigger/>",
			},
			sourcePath:    filepath.Join(tmpDir, "MyTrigger.trigger-meta.xml"),
			metaSuffix:    "-meta.xml",
			fileExt:       ".trigger",
			expectContent: "trigger MyTrigger on Account {}",
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

			// Test LoadCompanionFile
			result := LoadCompanionFile(tt.sourcePath, tt.metaSuffix, tt.fileExt)

			if tt.expectNil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tt.expectContent, string(result))
			}

			// Cleanup test files
			for filename := range tt.setupFiles {
				os.Remove(filepath.Join(tmpDir, filename))
			}
		})
	}
}

func TestLoadBundleFiles(t *testing.T) {
	// Create a temporary directory structure for testing
	tmpDir := t.TempDir()
	bundleDir := filepath.Join(tmpDir, "myComponent")
	require.NoError(t, os.MkdirAll(bundleDir, 0755))

	// Create __tests__ directory
	testsDir := filepath.Join(bundleDir, "__tests__")
	require.NoError(t, os.MkdirAll(testsDir, 0755))

	// Create test files
	files := map[string]string{
		"myComponent.js":          "export default class MyComponent extends LightningElement {}",
		"myComponent.html":        "<template><div>Hello</div></template>",
		"myComponent.css":         ".container { color: red; }",
		"myComponent.js-meta.xml": "<LightningComponentBundle/>",
		"__tests__/test.js":       "// test file",
	}

	for filename, content := range files {
		filePath := filepath.Join(bundleDir, filename)
		err := os.WriteFile(filePath, []byte(content), 0644)
		require.NoError(t, err)
	}

	tests := []struct {
		name         string
		sourcePath   string
		skipPatterns []string
		expectFiles  map[string]string
		expectError  bool
	}{
		{
			name:         "Load LWC bundle files",
			sourcePath:   filepath.Join(bundleDir, "myComponent.js-meta.xml"),
			skipPatterns: []string{"__tests__", "-meta.xml"},
			expectFiles: map[string]string{
				"myComponent.js":   "export default class MyComponent extends LightningElement {}",
				"myComponent.html": "<template><div>Hello</div></template>",
				"myComponent.css":  ".container { color: red; }",
			},
		},
		{
			name:         "Load all files except meta",
			sourcePath:   filepath.Join(bundleDir, "myComponent.js-meta.xml"),
			skipPatterns: []string{"-meta.xml"},
			expectFiles: map[string]string{
				"myComponent.js":    "export default class MyComponent extends LightningElement {}",
				"myComponent.html":  "<template><div>Hello</div></template>",
				"myComponent.css":   ".container { color: red; }",
				"__tests__/test.js": "// test file",
			},
		},
		{
			name:         "Empty source path",
			sourcePath:   "",
			skipPatterns: []string{},
			expectFiles:  nil,
		},
		{
			name:         "Non-meta.xml file",
			sourcePath:   filepath.Join(bundleDir, "myComponent.js"),
			skipPatterns: []string{},
			expectFiles:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := LoadBundleFiles(tt.sourcePath, tt.skipPatterns...)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				if tt.expectFiles == nil {
					assert.Nil(t, result)
				} else {
					assert.Equal(t, len(tt.expectFiles), len(result))
					for path, expectedContent := range tt.expectFiles {
						assert.Equal(t, expectedContent, string(result[path]))
					}
				}
			}
		})
	}
}

func TestGenerateMetadataFilePaths(t *testing.T) {
	xmlContent := []byte("<ApexClass/>")
	codeContent := []byte("public class MyClass {}")

	tests := []struct {
		name        string
		dirName     string
		baseName    string
		metaExt     string
		codeExt     string
		xmlContent  []byte
		codeContent []byte
		expectPaths map[string]string
	}{
		{
			name:        "ApexClass with code",
			dirName:     "classes",
			baseName:    "MyClass",
			metaExt:     ".cls-meta.xml",
			codeExt:     ".cls",
			xmlContent:  xmlContent,
			codeContent: codeContent,
			expectPaths: map[string]string{
				"classes/MyClass.cls-meta.xml": string(xmlContent),
				"classes/MyClass.cls":          string(codeContent),
			},
		},
		{
			name:        "ApexClass without code",
			dirName:     "classes",
			baseName:    "MyClass",
			metaExt:     ".cls-meta.xml",
			codeExt:     ".cls",
			xmlContent:  xmlContent,
			codeContent: nil,
			expectPaths: map[string]string{
				"classes/MyClass.cls-meta.xml": string(xmlContent),
			},
		},
		{
			name:        "ApexTrigger with code",
			dirName:     "triggers",
			baseName:    "MyTrigger",
			metaExt:     ".trigger-meta.xml",
			codeExt:     ".trigger",
			xmlContent:  xmlContent,
			codeContent: []byte("trigger MyTrigger on Account {}"),
			expectPaths: map[string]string{
				"triggers/MyTrigger.trigger-meta.xml": string(xmlContent),
				"triggers/MyTrigger.trigger":          "trigger MyTrigger on Account {}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateMetadataFilePaths(
				tt.dirName,
				tt.baseName,
				tt.metaExt,
				tt.codeExt,
				tt.xmlContent,
				tt.codeContent,
			)

			assert.Equal(t, len(tt.expectPaths), len(result))
			for path, content := range tt.expectPaths {
				assert.Equal(t, content, string(result[path]))
			}
		})
	}
}
