package internal

import "github.com/ForceCLI/force-md/metadata"

type metadataTypeRegistry map[string]OpenMetadataFunc

var TypeRegistry metadataTypeRegistry

func init() {
	TypeRegistry = make(metadataTypeRegistry)
}

type OpenMetadataFunc func(path string) (metadata.RegisterableMetadata, error)

func (r *metadataTypeRegistry) Register(metadataType string, openFunc OpenMetadataFunc) {
	if _, ok := (*r)[metadataType]; ok {
		panic("Duplicate metadata registration for " + metadataType)
	}
	(*r)[metadataType] = openFunc
}
