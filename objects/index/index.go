package index

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects/split"
)

const NAME = "Index"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type IndexFilter func(BigObjectIndex) bool

type Index struct {
	internal.MetadataInfo
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

func (c *Index) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Index) NameFromPath(path string) string {
	return split.NameFromPath(path)
}

func (c *Index) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Index, error) {
	p := &Index{}
	return p, internal.ParseMetadataXml(p, path)
}
