package custompermission

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/metadata"
)

type CustomPermission struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"CustomPermission"`
	Xmlns       string   `xml:"xmlns,attr"`
	Label       string   `xml:"label"`
	Description *string  `xml:"description"`
}

func (c *CustomPermission) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*CustomPermission, error) {
	p := &CustomPermission{}
	return p, metadata.ParseMetadataXml(p, path)
}

func New(label string) CustomPermission {
	p := CustomPermission{
		Label: label,
		Xmlns: "http://soap.sforce.com/2006/04/metadata",
	}
	return p
}

func (p *CustomPermission) EditDescription(description string) {
	p.Description = &description
}
