package index

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "Index"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type IndexFilter func(BigObjectIndex) bool

type Index struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Index"`
	Xmlns   string   `xml:"xmlns,attr"`
	BigObjectIndex
}

type BigObjectIndex struct {
	FullName string `xml:"fullName"`
	Fields   []struct {
		Name          string `xml:"name"`
		SortDirection string `xml:"sortDirection"`
	} `xml:"fields"`
	Label string `xml:"label"`
}

func (c *Index) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Index) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *Index) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Index, error) {
	p := &Index{}
	return p, metadata.ParseMetadataXml(p, path)
}
