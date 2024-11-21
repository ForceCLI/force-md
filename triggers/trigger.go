package trigger

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "ApexTrigger"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ApexTrigger struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"ApexTrigger"`
	Xmlns      string   `xml:"xmlns,attr"`
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

func (c *ApexTrigger) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexTrigger) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*ApexTrigger, error) {
	p := &ApexTrigger{}
	return p, internal.ParseMetadataXml(p, path)
}
