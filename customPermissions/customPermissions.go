package customPermissions

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "CustomPermission"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type CustomPermission struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"CustomPermission"`
	Xmlns      string   `xml:"xmlns,attr"`
	IsLicensed struct {
		Text string `xml:",chardata"`
	} `xml:"isLicensed"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *CustomPermission) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomPermission) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*CustomPermission, error) {
	p := &CustomPermission{}
	return p, internal.ParseMetadataXml(p, path)
}
