package registry

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
)

// FilesHelper provides a helper for implementing the Files() method for metadata types
// that consist of a single XML file without associated code files
type FilesHelper struct {
	MetadataType string
	Name         string
	Content      interface{} // The metadata struct to marshal
}

// GenerateFiles generates the files map for a metadata component
func (h *FilesHelper) GenerateFiles(format string) (map[string][]byte, error) {
	if h.Name == "" {
		return nil, fmt.Errorf("metadata name is empty")
	}

	// Get the directory name
	dirName := GetCanonicalDirectoryName(h.MetadataType)

	// Get the file suffix
	suffix := GetMetadataSuffix(h.MetadataType)
	if suffix == "" {
		// Default suffix based on type name
		suffix = h.MetadataType
	}

	// Marshal the XML content
	xmlContent, err := xml.MarshalIndent(h.Content, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %s metadata: %w", h.MetadataType, err)
	}
	// Add XML declaration
	xmlContent = append([]byte(xml.Header), xmlContent...)

	files := make(map[string][]byte)

	// For most metadata types, the file naming is the same for both formats
	// They all use -meta.xml suffix
	fileName := h.Name + "." + suffix + "-meta.xml"
	files[filepath.Join(dirName, fileName)] = xmlContent

	return files, nil
}
