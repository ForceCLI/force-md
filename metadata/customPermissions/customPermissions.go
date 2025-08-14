package customPermissions

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomPermission"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CustomPermission struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"CustomPermission"`
	Xmlns       string   `xml:"xmlns,attr"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	IsLicensed struct {
		Text string `xml:",chardata"`
	} `xml:"isLicensed"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

func (c *CustomPermission) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomPermission) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomPermission, error) {
	p := &CustomPermission{}
	return p, metadata.ParseMetadataXml(p, path)
}
