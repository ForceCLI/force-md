package installedPackages

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "InstalledPackage"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type InstalledPackage struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"InstalledPackage"`
	Xmlns       string   `xml:"xmlns,attr"`
	Xsi         string   `xml:"xsi,attr"`
	ActivateRSS struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"activateRSS"`
	VersionNumber struct {
		Text string `xml:",chardata"`
	} `xml:"versionNumber"`
}

func (c *InstalledPackage) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *InstalledPackage) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*InstalledPackage, error) {
	p := &InstalledPackage{}
	return p, metadata.ParseMetadataXml(p, path)
}
