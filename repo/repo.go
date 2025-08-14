package repo

import (
	"github.com/ForceCLI/force-md/metadata"
)

type Repo struct {
	openMetadata map[metadata.MetadataType]*MetadataByPath
}

type MetadataByPath map[metadata.MetadataFilePath]metadata.RegisterableMetadata

func NewRepo() *Repo {
	return &Repo{
		openMetadata: make(map[metadata.MetadataType]*MetadataByPath),
	}
}

func (o *Repo) Types() []metadata.MetadataType {
	var types []string
	for i := range o.openMetadata {
		types = append(types, i)
	}
	return types
}

// Items returns all metadata items for a type as a slice
func (o *Repo) Items(t metadata.MetadataType) []metadata.RegisterableMetadata {
	var items []metadata.RegisterableMetadata
	if pathMap, ok := o.openMetadata[t]; ok {
		for _, item := range *pathMap {
			items = append(items, item)
		}
	}
	return items
}

func (o *Repo) Open(file string) (metadata.MetadataPointer, error) {
	m, err := MetadataFromPath(file)
	if err != nil {
		return m, err
	}
	metadataType := m.Type()
	path := m.GetMetadataInfo().Path()
	if _, ok := o.openMetadata[metadataType]; !ok {
		items := make(MetadataByPath)
		o.openMetadata[metadataType] = &items
	}
	items := o.openMetadata[metadataType]
	(*items)[path] = m
	return m, nil
}
