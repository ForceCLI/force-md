package repo

import (
	"github.com/ForceCLI/force-md/registry"
)

// Re-export registry functions for backward compatibility
// These can be removed once all callers are updated to use registry directly

// GetMetadataDirectory returns the canonical directory name for a given metadata type.
func GetMetadataDirectory(metadataType string) string {
	return registry.GetMetadataDirectory(metadataType)
}

// GetParentType returns the parent metadata type for child types like CustomField
func GetParentType(childType string) string {
	return registry.GetParentType(childType)
}

// IsChildType returns true if the metadata type is a child of another type
func IsChildType(metadataType string) bool {
	return registry.IsChildType(metadataType)
}

// GetCanonicalDirectoryName returns the proper directory name for a metadata type
// considering both the type itself and any parent-child relationships
func GetCanonicalDirectoryName(metadataType string) string {
	return registry.GetCanonicalDirectoryName(metadataType)
}

// GetMetadataSuffix returns the file suffix for a given metadata type
func GetMetadataSuffix(metadataType string) string {
	return registry.GetMetadataSuffix(metadataType)
}
