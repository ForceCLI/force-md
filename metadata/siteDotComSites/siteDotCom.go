package siteDotComSite

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "SiteDotCom"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// SiteDotCom is the metadata of an Experience Builder or Site.com site. Its
// content is the companion .site file, which is not XML.
type SiteDotCom struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"SiteDotCom"`
	Xmlns   string   `xml:"xmlns,attr"`
	Label   struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	SiteType struct {
		Text string `xml:",chardata"`
	} `xml:"siteType"`
}

func (c *SiteDotCom) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *SiteDotCom) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*SiteDotCom, error) {
	p := &SiteDotCom{}
	return p, metadata.ParseMetadataXml(p, path)
}
