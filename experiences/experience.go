package experience

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "ExperienceBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ExperienceBundle struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"ExperienceBundle"`
	Xmlns   string   `xml:"xmlns,attr"`
	Label   struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	ExperienceType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	UrlPathPrefix struct {
		Text string `xml:",chardata"`
	} `xml:"urlPathPrefix"`
}

func (c *ExperienceBundle) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ExperienceBundle) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*ExperienceBundle, error) {
	p := &ExperienceBundle{}
	return p, internal.ParseMetadataXml(p, path)
}
