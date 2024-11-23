package translations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "Translations"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type Translations struct {
	internal.MetadataInfo
	XMLName      xml.Name `xml:"Translations"`
	Xmlns        string   `xml:"xmlns,attr"`
	CustomLabels []struct {
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"customLabels"`
}

func (c *Translations) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Translations) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Translations, error) {
	p := &Translations{}
	return p, internal.ParseMetadataXml(p, path)
}
