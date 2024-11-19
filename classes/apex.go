package apexClass

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "ApexClass"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ApexClass struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"ApexClass"`
	Xmlns      string   `xml:"xmlns,attr"`
	Fqn        string   `xml:"fqn,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
	PackageVersions struct {
		MajorNumber struct {
			Text string `xml:",chardata"`
		} `xml:"majorNumber"`
		MinorNumber struct {
			Text string `xml:",chardata"`
		} `xml:"minorNumber"`
		Namespace struct {
			Text string `xml:",chardata"`
		} `xml:"namespace"`
	} `xml:"packageVersions"`
}

func (c *ApexClass) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexClass) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*ApexClass, error) {
	p := &ApexClass{}
	return p, internal.ParseMetadataXml(p, path)
}
