package documentFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DocumentFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DocumentFolder struct {
	metadata.MetadataInfo
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

func (c *DocumentFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DocumentFolder) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DocumentFolder, error) {
	p := &DocumentFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
