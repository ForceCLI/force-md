package standardValueSetTranslation

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "StandardValueSetTranslation"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type StandardValueSetTranslation struct {
	metadata.MetadataInfo
	XMLName          xml.Name `xml:"StandardValueSetTranslation"`
	Xmlns            string   `xml:"xmlns,attr"`
	ValueTranslation []struct {
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		Translation struct {
			Text string `xml:",chardata"`
		} `xml:"translation"`
	} `xml:"valueTranslation"`
}

func (c *StandardValueSetTranslation) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *StandardValueSetTranslation) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*StandardValueSetTranslation, error) {
	p := &StandardValueSetTranslation{}
	return p, metadata.ParseMetadataXml(p, path)
}
