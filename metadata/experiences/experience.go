package experience

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ExperienceBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ExperienceBundle struct {
	metadata.MetadataInfo
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

func (c *ExperienceBundle) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ExperienceBundle) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ExperienceBundle, error) {
	p := &ExperienceBundle{}
	return p, metadata.ParseMetadataXml(p, path)
}
