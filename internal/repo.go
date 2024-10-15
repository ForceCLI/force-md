package internal

import "fmt"

var Metadata *repo

func init() {
	Metadata = &repo{
		openMetadata: make(map[MetadataType]*MetadataByName),
	}
}

type MetadataByName map[string]RegisterableMetadata

type repo struct {
	openMetadata map[MetadataType]*MetadataByName
}

func (o *repo) Types() []MetadataType {
	var types []string
	for i := range o.openMetadata {
		types = append(types, i)
	}
	return types
}

func (o *repo) Items(t MetadataType) MetadataByName {
	if _, ok := o.openMetadata[t]; !ok {
		return make(MetadataByName)
	}
	return *(*o).openMetadata[t]
}

func (o *repo) Open(file string) (MetadataPointer, error) {
	m, err := MetadataFromPath(file)
	if err != nil {
		return m, fmt.Errorf("invalid file %s: %w", file, err)
	}
	metadataType := m.Type()
	name := m.GetMetadataInfo().Name()
	if _, ok := o.openMetadata[metadataType]; !ok {
		items := make(MetadataByName)
		o.openMetadata[metadataType] = &items
	}
	items := o.openMetadata[metadataType]
	(*items)[name] = m
	return m, nil
}
