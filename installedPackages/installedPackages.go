package installedPackages

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "InstalledPackage"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type InstalledPackage struct {
	internal.MetadataInfo
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

func (c *InstalledPackage) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *InstalledPackage) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*InstalledPackage, error) {
	p := &InstalledPackage{}
	return p, internal.ParseMetadataXml(p, path)
}
