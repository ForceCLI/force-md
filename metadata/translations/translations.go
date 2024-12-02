package translations

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Translations"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Translations struct {
	metadata.MetadataInfo
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

func (c *Translations) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Translations) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Translations, error) {
	p := &Translations{}
	return p, metadata.ParseMetadataXml(p, path)
}
