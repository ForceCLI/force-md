// Package registry provides metadata type registry information without circular dependencies
package registry

import (
	_ "embed"
	"encoding/json"
	"strings"
)

// From https://github.com/forcedotcom/source-deploy-retrieve/blob/main/src/registry/metadataRegistry.json
//
//go:embed metadataRegistry.json
var metadataRegistryJSON []byte

type MetadataTypeInfo struct {
	DirectoryName       string `json:"directoryName"`
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Suffix              string `json:"suffix,omitempty"`
	StrictDirectoryName bool   `json:"strictDirectoryName"`
}

type MetadataRegistry struct {
	ChildTypes           map[string]string            `json:"childTypes"`
	StrictDirectoryNames map[string]string            `json:"strictDirectoryNames"`
	Suffixes             map[string]string            `json:"suffixes"`
	Types                map[string]*MetadataTypeInfo `json:"types"`
}

var registry *MetadataRegistry

func init() {
	registry = &MetadataRegistry{}
	if err := json.Unmarshal(metadataRegistryJSON, registry); err != nil {
		panic("Failed to unmarshal metadata registry: " + err.Error())
	}
}

// GetMetadataDirectory returns the canonical directory name for a given metadata type.
func GetMetadataDirectory(metadataType string) string {
	// First check the types registry which has the directoryName
	metadataTypeLower := strings.ToLower(metadataType)
	if typeInfo, ok := registry.Types[metadataTypeLower]; ok && typeInfo.DirectoryName != "" {
		return typeInfo.DirectoryName
	}

	// Then check strictDirectoryNames (maps directory -> type, so we reverse it)
	for dir, typeValue := range registry.StrictDirectoryNames {
		if strings.ToLower(typeValue) == metadataTypeLower {
			return dir
		}
	}

	// Default to lowercase with 's' for plural
	return strings.ToLower(metadataType) + "s"
}

// GetParentType returns the parent metadata type for child types like CustomField
func GetParentType(childType string) string {
	childTypeLower := strings.ToLower(childType)
	if parent, ok := registry.ChildTypes[childTypeLower]; ok {
		return parent
	}
	return ""
}

// IsChildType returns true if the metadata type is a child of another type
func IsChildType(metadataType string) bool {
	return GetParentType(metadataType) != ""
}

// GetCanonicalDirectoryName returns the proper directory name for a metadata type
// considering both the type itself and any parent-child relationships
func GetCanonicalDirectoryName(metadataType string) string {
	// First check if it's a child type
	if parent := GetParentType(metadataType); parent != "" {
		// For child types, return the parent's directory
		return GetMetadataDirectory(parent)
	}

	// Otherwise return the directory for this type
	return GetMetadataDirectory(metadataType)
}

// GetMetadataSuffix returns the file suffix for a given metadata type
func GetMetadataSuffix(metadataType string) string {
	// Check the types registry for the suffix
	metadataTypeLower := strings.ToLower(metadataType)
	if typeInfo, ok := registry.Types[metadataTypeLower]; ok && typeInfo.Suffix != "" {
		return typeInfo.Suffix
	}

	// Check the suffixes map
	if suffix, ok := registry.Suffixes[metadataTypeLower]; ok {
		return suffix
	}

	// No suffix for this type
	return ""
}
