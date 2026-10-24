package lightningExperienceThemes

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "LightningExperienceTheme"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// LightningExperienceTheme names the branding set a Lightning Experience
// theme uses
type LightningExperienceTheme struct {
	metadata.MetadataInfo
	XMLName                    xml.Name     `xml:"LightningExperienceTheme"`
	Xmlns                      string       `xml:"xmlns,attr"`
	DefaultBrandingSet         TextLiteral  `xml:"defaultBrandingSet"`
	Description                *TextLiteral `xml:"description"`
	DesignSystemVersion        *TextLiteral `xml:"designSystemVersion"`
	IsDarkModeEnabled          *TextLiteral `xml:"isDarkModeEnabled"`
	MasterLabel                TextLiteral  `xml:"masterLabel"`
	ShouldOverrideLoadingImage *TextLiteral `xml:"shouldOverrideLoadingImage"`
}

func (c *LightningExperienceTheme) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *LightningExperienceTheme) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*LightningExperienceTheme, error) {
	p := &LightningExperienceTheme{}
	return p, metadata.ParseMetadataXml(p, path)
}
