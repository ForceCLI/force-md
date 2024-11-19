package documentFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "DocumentFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type DocumentFolder struct {
	internal.MetadataInfo
	XMLName    xml.Name `xml:"DocumentFolder"`
	Xmlns      string   `xml:"xmlns,attr"`
	AccessType struct {
		Text string `xml:",chardata"`
	} `xml:"accessType"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	PublicFolderAccess struct {
		Text string `xml:",chardata"`
	} `xml:"publicFolderAccess"`
}

func (c *DocumentFolder) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DocumentFolder) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*DocumentFolder, error) {
	p := &DocumentFolder{}
	return p, internal.ParseMetadataXml(p, path)
}
