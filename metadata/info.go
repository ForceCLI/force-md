package metadata

import (
	"path/filepath"
	"strings"
)

type MetadataObjectName string
type MetadataFilePath string

type MetadataInfo struct {
	path     MetadataFilePath
	name     MetadataObjectName
	contents []byte
}

func (m MetadataInfo) NameFromPath(path string) MetadataObjectName {
	return NameFromPath(path)
}

func (m MetadataInfo) Name() MetadataObjectName {
	return m.name
}

func (m MetadataInfo) Path() MetadataFilePath {
	return m.path
}

func (m MetadataInfo) Contents() []byte {
	return m.contents
}

func (m MetadataInfo) GetMetadataInfo() MetadataInfo {
	return m
}

func NameFromPath(path string) MetadataObjectName {
	name := strings.TrimSuffix(filepath.Base(path), "-meta.xml")
	ext := filepath.Ext(name)
	return MetadataObjectName(strings.TrimSuffix(name, ext))
}
