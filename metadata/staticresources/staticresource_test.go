package staticresource

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

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

func TestZipDirectory(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := ioutil.TempDir("", "zipdir-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test files and directories
	testFiles := map[string]string{
		"index.html":           "<html><body>Test</body></html>",
		"css/styles.css":       "body { margin: 0; }",
		"js/app.js":            "console.log('test');",
		"images/logo.png":      "fake png content",
		"nested/deep/file.txt": "nested content",
	}

	for path, content := range testFiles {
		fullPath := filepath.Join(tmpDir, path)
		dir := filepath.Dir(fullPath)
		require.NoError(t, os.MkdirAll(dir, 0755))
		require.NoError(t, ioutil.WriteFile(fullPath, []byte(content), 0644))
	}

	// Also create an empty directory
	emptyDir := filepath.Join(tmpDir, "empty")
	require.NoError(t, os.MkdirAll(emptyDir, 0755))

	// Zip the directory
	zipContent, err := zipDirectory(tmpDir)
	require.NoError(t, err)
	require.NotNil(t, zipContent)
	require.True(t, len(zipContent) > 0)

	// Verify the zip content
	reader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	require.NoError(t, err)

	// Check that all files are present in the zip
	foundFiles := make(map[string]bool)
	for _, file := range reader.File {
		foundFiles[file.Name] = true

		// Verify compression method for files
		if !file.FileInfo().IsDir() {
			assert.Equal(t, zip.Deflate, file.Method, "File %s should use Deflate compression", file.Name)
		}

		// Verify directory names end with /
		if file.FileInfo().IsDir() {
			assert.True(t, filepath.ToSlash(file.Name)[len(file.Name)-1] == '/', "Directory %s should end with /", file.Name)
		}

		// Verify content for files
		if !file.FileInfo().IsDir() {
			rc, err := file.Open()
			require.NoError(t, err)
			content, err := ioutil.ReadAll(rc)
			rc.Close()
			require.NoError(t, err)

			// Check if this file's content matches what we expect
			for originalPath, expectedContent := range testFiles {
				if filepath.ToSlash(originalPath) == file.Name {
					assert.Equal(t, expectedContent, string(content), "Content mismatch for %s", file.Name)
					break
				}
			}
		}
	}

	// Check that all expected files are in the zip
	for path := range testFiles {
		slashPath := filepath.ToSlash(path)
		assert.True(t, foundFiles[slashPath], "Expected file %s not found in zip", slashPath)
	}

	// Check that expected directories are in the zip
	expectedDirs := []string{"css/", "js/", "images/", "nested/", "nested/deep/", "empty/"}
	for _, dir := range expectedDirs {
		assert.True(t, foundFiles[dir], "Expected directory %s not found in zip", dir)
	}
}

func TestZipDirectoryWithSingleFile(t *testing.T) {
	// Test zipping a directory with just one file
	tmpDir, err := ioutil.TempDir("", "zipdir-single-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	testFile := filepath.Join(tmpDir, "single.txt")
	testContent := "single file content"
	require.NoError(t, ioutil.WriteFile(testFile, []byte(testContent), 0644))

	zipContent, err := zipDirectory(tmpDir)
	require.NoError(t, err)

	// Verify the zip content
	reader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	require.NoError(t, err)

	assert.Equal(t, 1, len(reader.File), "Expected 1 file in zip")

	file := reader.File[0]
	assert.Equal(t, "single.txt", file.Name)
	assert.Equal(t, zip.Deflate, file.Method, "File should use Deflate compression")

	// Verify content
	rc, err := file.Open()
	require.NoError(t, err)
	defer rc.Close()

	content, err := ioutil.ReadAll(rc)
	require.NoError(t, err)
	assert.Equal(t, testContent, string(content))
}

func TestZipDirectoryEmptyDir(t *testing.T) {
	// Test zipping an empty directory
	tmpDir, err := ioutil.TempDir("", "zipdir-empty-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create just an empty subdirectory
	emptyDir := filepath.Join(tmpDir, "empty")
	require.NoError(t, os.MkdirAll(emptyDir, 0755))

	zipContent, err := zipDirectory(tmpDir)
	require.NoError(t, err)

	// Verify the zip content
	reader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	require.NoError(t, err)

	// Should contain just the empty directory
	assert.Equal(t, 1, len(reader.File), "Expected 1 entry in zip")

	file := reader.File[0]
	assert.Equal(t, "empty/", file.Name, "Expected directory name 'empty/'")
	assert.True(t, file.FileInfo().IsDir(), "Expected entry to be a directory")
}

func TestZipDirectoryTimestamps(t *testing.T) {
	// Test that all files have consistent timestamps
	tmpDir, err := ioutil.TempDir("", "zipdir-timestamp-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test files with different actual timestamps
	file1 := filepath.Join(tmpDir, "file1.txt")
	require.NoError(t, ioutil.WriteFile(file1, []byte("content1"), 0644))

	// Sleep to ensure different filesystem timestamp
	time.Sleep(10 * time.Millisecond)

	file2 := filepath.Join(tmpDir, "file2.txt")
	require.NoError(t, ioutil.WriteFile(file2, []byte("content2"), 0644))

	// Zip the directory
	beforeZip := time.Now().Add(-1 * time.Second) // Allow 1 second buffer
	zipContent, err := zipDirectory(tmpDir)
	require.NoError(t, err)
	afterZip := time.Now().Add(1 * time.Second) // Allow 1 second buffer

	// Verify the zip content
	reader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	require.NoError(t, err)

	assert.Equal(t, 2, len(reader.File), "Expected 2 files in zip")

	// Check that both files have the same timestamp
	var firstTimestamp time.Time
	for i, file := range reader.File {
		if i == 0 {
			firstTimestamp = file.Modified
			// Verify timestamp is reasonable (within the test execution window)
			assert.True(t, !firstTimestamp.Before(beforeZip) && !firstTimestamp.After(afterZip),
				"File timestamp %v should be between %v and %v", firstTimestamp, beforeZip, afterZip)
		} else {
			// All files should have the same timestamp
			assert.True(t, file.Modified.Equal(firstTimestamp),
				"File %s has different timestamp %v, expected %v", file.Name, file.Modified, firstTimestamp)
		}
	}
}

func TestZipDirectoryCompression(t *testing.T) {
	// Test that files are actually compressed
	tmpDir, err := ioutil.TempDir("", "zipdir-compression-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a file with repetitive content that should compress well
	testFile := filepath.Join(tmpDir, "compressible.txt")
	// Create 10KB of repeated text
	repeatText := "This is a test sentence that will be repeated many times. "
	var content bytes.Buffer
	for i := 0; i < 170; i++ { // ~10KB
		content.WriteString(repeatText)
	}

	require.NoError(t, ioutil.WriteFile(testFile, content.Bytes(), 0644))

	originalSize := content.Len()

	// Zip the directory
	zipContent, err := zipDirectory(tmpDir)
	require.NoError(t, err)

	// The zip should be significantly smaller than the original due to compression
	zipSize := len(zipContent)
	compressionRatio := float64(zipSize) / float64(originalSize)

	// With repetitive text, we should get at least 50% compression
	assert.True(t, compressionRatio < 0.5,
		"Compression ratio too low: %.2f (zip size: %d, original: %d)", compressionRatio, zipSize, originalSize)

	// Verify the content is correct when decompressed
	reader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	require.NoError(t, err)

	assert.Equal(t, 1, len(reader.File), "Expected 1 file in zip")

	file := reader.File[0]
	rc, err := file.Open()
	require.NoError(t, err)
	defer rc.Close()

	decompressed, err := ioutil.ReadAll(rc)
	require.NoError(t, err)

	assert.Equal(t, content.Bytes(), decompressed, "Decompressed content should match original")
}

func TestZipDirectorySalesforceCompliance(t *testing.T) {
	// Test that the zip file format is compatible with Salesforce requirements
	tmpDir, err := ioutil.TempDir("", "zipdir-salesforce-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a typical static resource structure
	require.NoError(t, os.MkdirAll(filepath.Join(tmpDir, "css"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(tmpDir, "js"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(tmpDir, "img"), 0755))

	files := map[string]string{
		"css/main.css":     "body { font-family: Arial; }",
		"js/app.js":        "function init() { console.log('ready'); }",
		"img/logo.svg":     "<svg></svg>",
		"index.html":       "<html><head></head><body></body></html>",
	}

	for path, content := range files {
		fullPath := filepath.Join(tmpDir, path)
		require.NoError(t, ioutil.WriteFile(fullPath, []byte(content), 0644))
	}

	// Zip the directory
	zipContent, err := zipDirectory(tmpDir)
	require.NoError(t, err)

	// Verify the zip has proper structure for Salesforce
	reader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	require.NoError(t, err)

	// Check all expected files and directories are present
	expectedEntries := map[string]bool{
		"css/":         true,  // directory
		"css/main.css": false, // file
		"js/":          true,  // directory
		"js/app.js":    false, // file
		"img/":         true,  // directory
		"img/logo.svg": false, // file
		"index.html":   false, // file
	}

	foundEntries := make(map[string]bool)
	for _, file := range reader.File {
		foundEntries[file.Name] = true

		isDir, expectedIsDir := expectedEntries[file.Name]
		if expectedIsDir {
			assert.Equal(t, isDir, file.FileInfo().IsDir(),
				"Entry %s directory status mismatch", file.Name)
		}

		// Verify all files use Deflate compression
		if !file.FileInfo().IsDir() {
			assert.Equal(t, zip.Deflate, file.Method,
				"File %s should use Deflate compression for Salesforce compatibility", file.Name)
		}

		// Verify no file has zero timestamp (Salesforce requirement)
		assert.False(t, file.Modified.IsZero(),
			"File %s should have a valid timestamp for Salesforce", file.Name)
	}

	// Verify all expected entries were found
	for entry := range expectedEntries {
		assert.True(t, foundEntries[entry], "Expected entry %s not found in zip", entry)
	}
}
