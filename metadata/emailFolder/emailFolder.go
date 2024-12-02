package emailFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "EmailFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type EmailFolder struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"EmailFolder"`
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
	SharedTo struct {
	} `xml:"sharedTo"`
}

func (c *EmailFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EmailFolder) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*EmailFolder, error) {
	p := &EmailFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
