package globalValueSetTranslation

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "GlobalValueSetTranslation"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type GlobalValueSetTranslation struct {
	internal.MetadataInfo
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

func (c *GlobalValueSetTranslation) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *GlobalValueSetTranslation) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*GlobalValueSetTranslation, error) {
	p := &GlobalValueSetTranslation{}
	return p, internal.ParseMetadataXml(p, path)
}
