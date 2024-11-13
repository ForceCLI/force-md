package internal

var TypeRegistry metadataTypeRegistry

func init() {
	TypeRegistry = make(metadataTypeRegistry)
}

type OpenMetadataFunc func(path string) (RegisterableMetadata, error)

type metadataTypeRegistry map[string]OpenMetadataFunc

func (r *metadataTypeRegistry) Register(metadataType string, openFunc OpenMetadataFunc) {
	if _, ok := (*r)[metadataType]; ok {
		panic("Duplicate metadata registration for " + metadataType)
	}
	(*r)[metadataType] = openFunc
}
