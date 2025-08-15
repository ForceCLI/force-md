package helpers

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// LoadCompanionFile loads a companion file for metadata that has separate code files.
// For example, ApexClass has .cls-meta.xml (metadata) and .cls (code).
//
// Parameters:
//   - sourcePath: The path to the metadata file (e.g., "MyClass.cls-meta.xml")
//   - metaSuffix: The metadata suffix to remove (e.g., "-meta.xml")
//   - fileExt: The companion file extension (e.g., ".cls")
//
// Returns the file content or nil if the file doesn't exist.
func LoadCompanionFile(sourcePath, metaSuffix, fileExt string) []byte {
	if sourcePath == "" {
		return nil
	}

	var companionPath string
	if strings.HasSuffix(sourcePath, metaSuffix) {
		// Source format: remove the meta suffix to get companion file
		companionPath = strings.TrimSuffix(sourcePath, metaSuffix)
	} else if strings.HasSuffix(sourcePath, ".xml") {
		// Metadata format: could be .cls.xml or just .xml
		base := strings.TrimSuffix(sourcePath, ".xml")
		if strings.HasSuffix(base, fileExt) {
			// It's already the right format (e.g., .cls.xml -> .cls)
			companionPath = base
		} else {
			// Add the extension (e.g., MyClass.xml -> MyClass.cls)
			companionPath = base + fileExt
		}
	}

	if companionPath != "" {
		if content, err := os.ReadFile(companionPath); err == nil {
			return content
		}
		// File doesn't exist, that's okay
	}
	return nil
}

// LoadBundleFiles loads all files from a bundle directory (like LWC or Aura).
//
// Parameters:
//   - sourcePath: The path to the metadata file (e.g., "myComponent.js-meta.xml")
//   - skipPatterns: Patterns to skip (e.g., "__tests__", "-meta.xml")
//
// Returns a map of relative path to file content.
func LoadBundleFiles(sourcePath string, skipPatterns ...string) (map[string][]byte, error) {
	if sourcePath == "" || !strings.HasSuffix(sourcePath, "-meta.xml") {
		return nil, nil
	}

	bundleDir := filepath.Dir(sourcePath)
	files := make(map[string][]byte)

	err := filepath.WalkDir(bundleDir, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from bundle directory
		relPath, err := filepath.Rel(bundleDir, filePath)
		if err != nil {
			return err
		}

		// Check skip patterns
		for _, pattern := range skipPatterns {
			if strings.Contains(relPath, pattern) {
				if d.IsDir() && pattern == "__tests__" {
					return filepath.SkipDir
				}
				return nil
			}
			if strings.HasSuffix(filePath, pattern) {
				return nil
			}
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Read the file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		// Store with relative path from bundle directory
		files[relPath] = content
		return nil
	})

	return files, err
}

// GenerateMetadataFilePaths generates the file paths for metadata and companion files.
//
// Parameters:
//   - dirName: The canonical directory name (e.g., "classes")
//   - baseName: The base file name without extension (e.g., "MyClass")
//   - metaExt: The metadata file extension (e.g., ".cls-meta.xml")
//   - codeExt: The code file extension (e.g., ".cls")
//   - xmlContent: The marshalled XML content
//   - codeContent: The code content (can be nil)
//
// Returns a map of file paths to content.
func GenerateMetadataFilePaths(dirName, baseName, metaExt, codeExt string, xmlContent, codeContent []byte) map[string][]byte {
	files := make(map[string][]byte)

	// Add metadata file
	files[filepath.Join(dirName, baseName+metaExt)] = xmlContent

	// Add code file if present
	if codeContent != nil {
		files[filepath.Join(dirName, baseName+codeExt)] = codeContent
	}

	return files
}
