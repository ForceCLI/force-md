package repo

import (
	"github.com/ForceCLI/force-md/metadata"
	log "github.com/sirupsen/logrus"
)

type Repo struct {
	openMetadata map[metadata.MetadataType]*MetadataByName
}

type MetadataByName map[metadata.MetadataObjectName]metadata.RegisterableMetadata

func NewRepo() *Repo {
	return &Repo{
		openMetadata: make(map[metadata.MetadataType]*MetadataByName),
	}
}

func (o *Repo) Types() []metadata.MetadataType {
	var types []string
	for i := range o.openMetadata {
		types = append(types, i)
	}
	return types
}

func (o *Repo) Items(t metadata.MetadataType) MetadataByName {
	if _, ok := o.openMetadata[t]; !ok {
		return make(MetadataByName)
	}
	return *(*o).openMetadata[t]
}

func (o *Repo) Open(file string) (metadata.MetadataPointer, error) {
	m, err := MetadataFromPath(file)
	if err != nil {
		return m, err
	}
	metadataType := m.Type()
	name := m.GetMetadataInfo().Name()
	if _, ok := o.openMetadata[metadataType]; !ok {
		items := make(MetadataByName)
		o.openMetadata[metadataType] = &items
	}
	items := o.openMetadata[metadataType]
	if _, exists := (*items)[name]; exists {
		log.Warnf("file %s of type %s already registered", name, metadataType)
	}
	(*items)[name] = m
	return m, nil
}
