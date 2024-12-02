package apexClass

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ApexClass"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexClass struct {
	metadata.MetadataInfo
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

func (c *ApexClass) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexClass) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ApexClass, error) {
	p := &ApexClass{}
	return p, metadata.ParseMetadataXml(p, path)
}
