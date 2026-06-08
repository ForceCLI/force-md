package datacategorygroup

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DataCategoryGroup"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// DataCategory is a node in a data category group's category tree. The root
// node's children are the top-level categories; each child may have further
// children, forming the tree.
type DataCategory struct {
	DataCategory []DataCategory `xml:"dataCategory"`
	Label        string         `xml:"label"`
	Name         string         `xml:"name"`
}

// ObjectUsage lists the SObjects a data category group applies to (e.g.,
// KnowledgeArticleVersion, Question).
type ObjectUsage struct {
	Object []string `xml:"object"`
}

type DataCategoryGroup struct {
	metadata.MetadataInfo
	XMLName      xml.Name     `xml:"DataCategoryGroup"`
	Xmlns        string       `xml:"xmlns,attr"`
	Active       BooleanText  `xml:"active"`
	Description  *string      `xml:"description"`
	Label        string       `xml:"label"`
	DataCategory DataCategory `xml:"dataCategory"`
	ObjectUsage  *ObjectUsage `xml:"objectUsage"`
}

func (c *DataCategoryGroup) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DataCategoryGroup) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DataCategoryGroup, error) {
	p := &DataCategoryGroup{}
	return p, metadata.ParseMetadataXml(p, path)
}
