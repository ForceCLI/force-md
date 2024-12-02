package globalValueSetTranslation

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "GlobalValueSetTranslation"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type GlobalValueSetTranslation struct {
	metadata.MetadataInfo
	XMLName          xml.Name `xml:"GlobalValueSetTranslation"`
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

func (c *GlobalValueSetTranslation) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *GlobalValueSetTranslation) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*GlobalValueSetTranslation, error) {
	p := &GlobalValueSetTranslation{}
	return p, metadata.ParseMetadataXml(p, path)
}
