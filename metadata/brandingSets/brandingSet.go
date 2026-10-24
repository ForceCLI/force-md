package brandingSets

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "BrandingSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// BrandingSet holds the colors and images of a Lightning Experience theme
type BrandingSet struct {
	metadata.MetadataInfo
	XMLName             xml.Name `xml:"BrandingSet"`
	Xmlns               string   `xml:"xmlns,attr"`
	BrandingSetProperty []struct {
		PropertyName  TextLiteral `xml:"propertyName"`
		PropertyValue TextLiteral `xml:"propertyValue"`
	} `xml:"brandingSetProperty"`
	Description     *TextLiteral `xml:"description"`
	MasterLabel     TextLiteral  `xml:"masterLabel"`
	BrandingSetType *TextLiteral `xml:"type"`
}

func (c *BrandingSet) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

// KeepEmptyElements marks branding sets as written with empty elements open,
// as Salesforce writes <propertyValue></propertyValue> for a property
// without a value.
func (c *BrandingSet) KeepEmptyElements() {}

func (c *BrandingSet) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*BrandingSet, error) {
	p := &BrandingSet{}
	return p, metadata.ParseMetadataXml(p, path)
}
