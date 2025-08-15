package lwc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLightningComponentBundleTidy(t *testing.T) {
	tests := []struct {
		name     string
		lwc      *LightningComponentBundle
		expected *LightningComponentBundle
	}{
		{
			name: "Sort targets",
			lwc: &LightningComponentBundle{
				Targets: &struct {
					Target []struct {
						Text string `xml:",chardata"`
					} `xml:"target"`
				}{
					Target: []struct {
						Text string `xml:",chardata"`
					}{
						{Text: "lightning__RecordPage"},
						{Text: "lightning__AppPage"},
						{Text: "lightning__HomePage"},
					},
				},
			},
			expected: &LightningComponentBundle{
				Targets: &struct {
					Target []struct {
						Text string `xml:",chardata"`
					} `xml:"target"`
				}{
					Target: []struct {
						Text string `xml:",chardata"`
					}{
						{Text: "lightning__AppPage"},
						{Text: "lightning__HomePage"},
						{Text: "lightning__RecordPage"},
					},
				},
			},
		},
		{
			name: "Sort targetConfigs and properties",
			lwc: &LightningComponentBundle{
				TargetConfigs: &struct {
					TargetConfig []struct {
						Xmlns    *string `xml:"xmlns,attr"`
						Targets  string  `xml:"targets,attr"`
						Property []struct {
							Name        string  `xml:"name,attr"`
							Label       *string `xml:"label,attr"`
							Type        string  `xml:"type,attr"`
							Role        *string `xml:"role,attr"`
							Description *string `xml:"description,attr"`
							Datasource  *string `xml:"datasource,attr"`
							Default     *string `xml:"default,attr"`
							Required    *string `xml:"required,attr"`
						} `xml:"property"`
						SupportedFormFactors *struct {
							SupportedFormFactor []struct {
								Type string `xml:"type,attr"`
							} `xml:"supportedFormFactor"`
						} `xml:"supportedFormFactors"`
						ActionType *struct {
							Text string `xml:",chardata"`
						} `xml:"actionType"`
					} `xml:"targetConfig"`
				}{
					TargetConfig: []struct {
						Xmlns    *string `xml:"xmlns,attr"`
						Targets  string  `xml:"targets,attr"`
						Property []struct {
							Name        string  `xml:"name,attr"`
							Label       *string `xml:"label,attr"`
							Type        string  `xml:"type,attr"`
							Role        *string `xml:"role,attr"`
							Description *string `xml:"description,attr"`
							Datasource  *string `xml:"datasource,attr"`
							Default     *string `xml:"default,attr"`
							Required    *string `xml:"required,attr"`
						} `xml:"property"`
						SupportedFormFactors *struct {
							SupportedFormFactor []struct {
								Type string `xml:"type,attr"`
							} `xml:"supportedFormFactor"`
						} `xml:"supportedFormFactors"`
						ActionType *struct {
							Text string `xml:",chardata"`
						} `xml:"actionType"`
					}{
						{
							Targets: "lightning__RecordPage",
							Property: []struct {
								Name        string  `xml:"name,attr"`
								Label       *string `xml:"label,attr"`
								Type        string  `xml:"type,attr"`
								Role        *string `xml:"role,attr"`
								Description *string `xml:"description,attr"`
								Datasource  *string `xml:"datasource,attr"`
								Default     *string `xml:"default,attr"`
								Required    *string `xml:"required,attr"`
							}{
								{Name: "recordId", Type: "String"},
								{Name: "objectApiName", Type: "String"},
							},
						},
						{
							Targets: "lightning__AppPage",
							Property: []struct {
								Name        string  `xml:"name,attr"`
								Label       *string `xml:"label,attr"`
								Type        string  `xml:"type,attr"`
								Role        *string `xml:"role,attr"`
								Description *string `xml:"description,attr"`
								Datasource  *string `xml:"datasource,attr"`
								Default     *string `xml:"default,attr"`
								Required    *string `xml:"required,attr"`
							}{
								{Name: "title", Type: "String"},
							},
						},
					},
				},
			},
			expected: &LightningComponentBundle{
				TargetConfigs: &struct {
					TargetConfig []struct {
						Xmlns    *string `xml:"xmlns,attr"`
						Targets  string  `xml:"targets,attr"`
						Property []struct {
							Name        string  `xml:"name,attr"`
							Label       *string `xml:"label,attr"`
							Type        string  `xml:"type,attr"`
							Role        *string `xml:"role,attr"`
							Description *string `xml:"description,attr"`
							Datasource  *string `xml:"datasource,attr"`
							Default     *string `xml:"default,attr"`
							Required    *string `xml:"required,attr"`
						} `xml:"property"`
						SupportedFormFactors *struct {
							SupportedFormFactor []struct {
								Type string `xml:"type,attr"`
							} `xml:"supportedFormFactor"`
						} `xml:"supportedFormFactors"`
						ActionType *struct {
							Text string `xml:",chardata"`
						} `xml:"actionType"`
					} `xml:"targetConfig"`
				}{
					TargetConfig: []struct {
						Xmlns    *string `xml:"xmlns,attr"`
						Targets  string  `xml:"targets,attr"`
						Property []struct {
							Name        string  `xml:"name,attr"`
							Label       *string `xml:"label,attr"`
							Type        string  `xml:"type,attr"`
							Role        *string `xml:"role,attr"`
							Description *string `xml:"description,attr"`
							Datasource  *string `xml:"datasource,attr"`
							Default     *string `xml:"default,attr"`
							Required    *string `xml:"required,attr"`
						} `xml:"property"`
						SupportedFormFactors *struct {
							SupportedFormFactor []struct {
								Type string `xml:"type,attr"`
							} `xml:"supportedFormFactor"`
						} `xml:"supportedFormFactors"`
						ActionType *struct {
							Text string `xml:",chardata"`
						} `xml:"actionType"`
					}{
						{
							Targets: "lightning__AppPage",
							Property: []struct {
								Name        string  `xml:"name,attr"`
								Label       *string `xml:"label,attr"`
								Type        string  `xml:"type,attr"`
								Role        *string `xml:"role,attr"`
								Description *string `xml:"description,attr"`
								Datasource  *string `xml:"datasource,attr"`
								Default     *string `xml:"default,attr"`
								Required    *string `xml:"required,attr"`
							}{
								{Name: "title", Type: "String"},
							},
						},
						{
							Targets: "lightning__RecordPage",
							Property: []struct {
								Name        string  `xml:"name,attr"`
								Label       *string `xml:"label,attr"`
								Type        string  `xml:"type,attr"`
								Role        *string `xml:"role,attr"`
								Description *string `xml:"description,attr"`
								Datasource  *string `xml:"datasource,attr"`
								Default     *string `xml:"default,attr"`
								Required    *string `xml:"required,attr"`
							}{
								{Name: "objectApiName", Type: "String"},
								{Name: "recordId", Type: "String"},
							},
						},
					},
				},
			},
		},
		{
			name: "Handle nil fields gracefully",
			lwc: &LightningComponentBundle{
				Targets:       nil,
				TargetConfigs: nil,
			},
			expected: &LightningComponentBundle{
				Targets:       nil,
				TargetConfigs: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lwc.Tidy()
			assert.Equal(t, tt.expected, tt.lwc)
		})
	}
}

func TestLightningComponentBundleOpen(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	bundleDir := filepath.Join(tmpDir, "myComponent")
	require.NoError(t, os.MkdirAll(bundleDir, 0755))

	// Create test metadata file
	metaPath := filepath.Join(bundleDir, "myComponent.js-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<LightningComponentBundle xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>58.0</apiVersion>
    <isExposed>true</isExposed>
</LightningComponentBundle>`

	err := os.WriteFile(metaPath, []byte(metaContent), 0644)
	require.NoError(t, err)

	// Test Open function
	lwc, err := Open(metaPath)
	require.NoError(t, err)
	assert.NotNil(t, lwc)
	assert.Equal(t, metaPath, lwc.SourcePath)
	assert.Equal(t, "58.0", lwc.ApiVersion.Text)
	assert.Equal(t, "true", lwc.IsExposed.Text)
	assert.Empty(t, lwc.BundleFiles) // Should not load files in Open
}

func TestLightningComponentBundleFiles(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	bundleDir := filepath.Join(tmpDir, "myComponent")
	require.NoError(t, os.MkdirAll(bundleDir, 0755))

	// Create __tests__ directory
	testsDir := filepath.Join(bundleDir, "__tests__")
	require.NoError(t, os.MkdirAll(testsDir, 0755))

	tests := []struct {
		name         string
		setupFiles   map[string]string
		format       metadata.Format
		expectFiles  map[string]bool
		expectBundle map[string]string // Expected bundle files after loading
	}{
		{
			name: "Source format with bundle files",
			setupFiles: map[string]string{
				"myComponent.js":          "export default class MyComponent extends LightningElement {}",
				"myComponent.html":        "<template><div>Hello</div></template>",
				"myComponent.css":         ".container { color: red; }",
				"myComponent.js-meta.xml": `<LightningComponentBundle xmlns="http://soap.sforce.com/2006/04/metadata"><apiVersion>58.0</apiVersion></LightningComponentBundle>`,
				"__tests__/test.js":       "// test file",
			},
			format: metadata.SourceFormat,
			expectFiles: map[string]bool{
				"lwc/myComponent/myComponent.js-meta.xml": true,
				"lwc/myComponent/myComponent.js":          true,
				"lwc/myComponent/myComponent.html":        true,
				"lwc/myComponent/myComponent.css":         true,
			},
			expectBundle: map[string]string{
				"myComponent.js":   "export default class MyComponent extends LightningElement {}",
				"myComponent.html": "<template><div>Hello</div></template>",
				"myComponent.css":  ".container { color: red; }",
			},
		},
		{
			name: "Metadata format with bundle files",
			setupFiles: map[string]string{
				"myComponent.js":          "export default class MyComponent extends LightningElement {}",
				"myComponent.js-meta.xml": `<LightningComponentBundle xmlns="http://soap.sforce.com/2006/04/metadata"><apiVersion>58.0</apiVersion></LightningComponentBundle>`,
			},
			format: metadata.MetadataFormat,
			expectFiles: map[string]bool{
				"lwc/myComponent/myComponent.js-meta.xml": true,
				"lwc/myComponent/myComponent.js":          true,
			},
			expectBundle: map[string]string{
				"myComponent.js": "export default class MyComponent extends LightningElement {}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test files
			for filename, content := range tt.setupFiles {
				filePath := filepath.Join(bundleDir, filename)
				// Create subdirectory if needed
				dir := filepath.Dir(filePath)
				if dir != bundleDir {
					require.NoError(t, os.MkdirAll(dir, 0755))
				}
				err := os.WriteFile(filePath, []byte(content), 0644)
				require.NoError(t, err)
			}

			// Create LightningComponentBundle instance using Open to properly set metadata
			lwc, err := Open(filepath.Join(bundleDir, "myComponent.js-meta.xml"))
			require.NoError(t, err)

			// Test Files method
			files, err := lwc.Files(tt.format)
			require.NoError(t, err)

			// Check expected files
			for path, shouldExist := range tt.expectFiles {
				if shouldExist {
					assert.Contains(t, files, path, "Expected file %s to exist", path)
				} else {
					assert.NotContains(t, files, path, "Expected file %s to not exist", path)
				}
			}

			// Check bundle files were loaded correctly (excluding __tests__)
			assert.Equal(t, len(tt.expectBundle), len(lwc.BundleFiles))
			for filename, content := range tt.expectBundle {
				assert.Equal(t, content, string(lwc.BundleFiles[filename]))
			}

			// Cleanup test files
			os.RemoveAll(bundleDir)
			require.NoError(t, os.MkdirAll(bundleDir, 0755))
		})
	}
}

func TestLightningComponentBundleFilesUnsupportedFormat(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	bundleDir := filepath.Join(tmpDir, "myComponent")
	require.NoError(t, os.MkdirAll(bundleDir, 0755))

	metaPath := filepath.Join(bundleDir, "myComponent.js-meta.xml")
	metaContent := `<?xml version="1.0" encoding="UTF-8"?>
<LightningComponentBundle xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>58.0</apiVersion>
</LightningComponentBundle>`
	err := os.WriteFile(metaPath, []byte(metaContent), 0644)
	require.NoError(t, err)

	lwc, err := Open(metaPath)
	require.NoError(t, err)

	// Test with an invalid format
	_, err = lwc.Files(metadata.Format("invalid"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported format")
}
