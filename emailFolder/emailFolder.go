package emailFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "EmailFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type EmailFolder struct {
	internal.MetadataInfo
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

func (c *EmailFolder) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EmailFolder) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*EmailFolder, error) {
	p := &EmailFolder{}
	return p, internal.ParseMetadataXml(p, path)
}
