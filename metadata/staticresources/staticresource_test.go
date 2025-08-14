package staticresource

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ForceCLI/force-md/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStaticResourceFiles(t *testing.T) {
	tests := []struct {
		name          string
		resourceName  string
		contentFile   string
		contentData   string
		expectedFiles int
		format        metadata.Format
	}{
		{
			name:          "CSS file with .resource extension",
			resourceName:  "TestStyle",
			contentFile:   "TestStyle.resource",
			contentData:   ".test { color: red; }",
			expectedFiles: 2,
			format:        metadata.MetadataFormat,
		},
		{
			name:          "CSS file with actual .css extension",
			resourceName:  "TestStyle",
			contentFile:   "TestStyle.css",
			contentData:   ".test { color: blue; }",
			expectedFiles: 2,
			format:        metadata.SourceFormat,
		},
		{
			name:          "JavaScript file with .js extension",
			resourceName:  "TestScript",
			contentFile:   "TestScript.js",
			contentData:   "console.log('test');",
			expectedFiles: 2,
			format:        metadata.SourceFormat,
		},
		{
			name:          "ZIP file with .zip extension",
			resourceName:  "TestArchive",
			contentFile:   "TestArchive.zip",
			contentData:   "PK\x03\x04", // ZIP file header
			expectedFiles: 2,
			format:        metadata.MetadataFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp directory
			tmpDir, err := ioutil.TempDir("", "staticresource-test")
			require.NoError(t, err)
			defer os.RemoveAll(tmpDir)

			// Create metadata file
			metadataPath := filepath.Join(tmpDir, tt.resourceName+".resource-meta.xml")
			metadataContent := `<?xml version="1.0" encoding="UTF-8"?>
<StaticResource xmlns="http://soap.sforce.com/2006/04/metadata">
    <cacheControl>Public</cacheControl>
    <contentType>text/css</contentType>
</StaticResource>`
			err = ioutil.WriteFile(metadataPath, []byte(metadataContent), 0644)
			require.NoError(t, err)

			// Create content file
			contentPath := filepath.Join(tmpDir, tt.contentFile)
			err = ioutil.WriteFile(contentPath, []byte(tt.contentData), 0644)
			require.NoError(t, err)

			// Open and parse the static resource
			sr, err := Open(metadataPath)
			require.NoError(t, err)

			// Get files
			files, err := sr.Files(tt.format)
			require.NoError(t, err)

			// Verify we got the expected number of files
			assert.Equal(t, tt.expectedFiles, len(files))

			// Verify metadata file exists
			foundMetadata := false
			foundContent := false
			for path, content := range files {
				if filepath.Ext(path) == ".xml" {
					foundMetadata = true
					assert.Contains(t, string(content), "StaticResource")
				} else if filepath.Ext(path) == ".resource" {
					foundContent = true
					assert.Equal(t, tt.contentData, string(content))
				}
			}

			assert.True(t, foundMetadata, "Should have metadata file")
			assert.True(t, foundContent, "Should have content file")
		})
	}
}

func TestStaticResourceNoContent(t *testing.T) {
	// Create temp directory
	tmpDir, err := ioutil.TempDir("", "staticresource-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create only metadata file, no content
	metadataPath := filepath.Join(tmpDir, "NoContent.resource-meta.xml")
	metadataContent := `<?xml version="1.0" encoding="UTF-8"?>
<StaticResource xmlns="http://soap.sforce.com/2006/04/metadata">
    <cacheControl>Private</cacheControl>
    <contentType>application/octet-stream</contentType>
</StaticResource>`
	err = ioutil.WriteFile(metadataPath, []byte(metadataContent), 0644)
	require.NoError(t, err)

	// Open and parse the static resource
	sr, err := Open(metadataPath)
	require.NoError(t, err)

	// Get files for metadata format
	files, err := sr.Files(metadata.MetadataFormat)
	require.NoError(t, err)

	// Should only have metadata file when no content exists
	assert.Equal(t, 1, len(files))

	// Verify it's the metadata file
	for path, content := range files {
		assert.True(t, filepath.Ext(path) == ".xml")
		assert.Contains(t, string(content), "StaticResource")
	}
}

func TestStaticResourceDirectory(t *testing.T) {
	// Create temp directory
	tmpDir, err := ioutil.TempDir("", "staticresource-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create metadata file
	metadataPath := filepath.Join(tmpDir, "FontAwesome.resource-meta.xml")
	metadataContent := `<?xml version="1.0" encoding="UTF-8"?>
<StaticResource xmlns="http://soap.sforce.com/2006/04/metadata">
    <cacheControl>Public</cacheControl>
    <contentType>application/zip</contentType>
</StaticResource>`
	err = ioutil.WriteFile(metadataPath, []byte(metadataContent), 0644)
	require.NoError(t, err)

	// Create a directory structure that represents an unzipped static resource
	resourceDir := filepath.Join(tmpDir, "FontAwesome")
	require.NoError(t, os.MkdirAll(filepath.Join(resourceDir, "css"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(resourceDir, "fonts"), 0755))

	// Add some files to the directory
	cssContent := "body { font-family: FontAwesome; }"
	err = ioutil.WriteFile(filepath.Join(resourceDir, "css", "font-awesome.css"), []byte(cssContent), 0644)
	require.NoError(t, err)

	fontContent := "binary font data"
	err = ioutil.WriteFile(filepath.Join(resourceDir, "fonts", "fontawesome.woff"), []byte(fontContent), 0644)
	require.NoError(t, err)

	// Open and parse the static resource
	sr, err := Open(metadataPath)
	require.NoError(t, err)

	// Get files for metadata format (should zip the directory)
	files, err := sr.Files(metadata.MetadataFormat)
	require.NoError(t, err)

	// Should have metadata and zipped content
	assert.Equal(t, 2, len(files))

	// Verify we have both files
	foundMetadata := false
	foundZip := false
	for path, content := range files {
		if filepath.Ext(path) == ".xml" {
			foundMetadata = true
			assert.Contains(t, string(content), "StaticResource")
		} else if filepath.Ext(path) == ".resource" {
			foundZip = true
			// Verify it's a valid zip file by checking for ZIP header
			assert.True(t, len(content) > 4)
			assert.Equal(t, []byte("PK"), content[0:2], "Should be a valid ZIP file")
		}
	}

	assert.True(t, foundMetadata, "Should have metadata file")
	assert.True(t, foundZip, "Should have zipped resource file")
}
