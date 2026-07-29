package productAttributeSets

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ProductAttributeSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ProductAttributeSet struct {
	metadata.MetadataInfo
	XMLName       xml.Name     `xml:"ProductAttributeSet"`
	Xmlns         string       `xml:"xmlns,attr"`
	Description   *TextLiteral `xml:"description"`
	DeveloperName *TextLiteral `xml:"developerName"`
	MasterLabel   *TextLiteral `xml:"masterLabel"`
}

func (p *ProductAttributeSet) SetMetadata(m metadata.MetadataInfo) {
	p.MetadataInfo = m
}

func (p *ProductAttributeSet) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ProductAttributeSet, error) {
	p := &ProductAttributeSet{}
	return p, metadata.ParseMetadataXml(p, path)
}
