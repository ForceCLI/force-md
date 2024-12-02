package permissionsetgroup

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "PermissionSetGroup"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type PermissionSet struct {
	Text string `xml:",chardata"`
}

type PermissionSetList []PermissionSet

type PermissionSetGroup struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"PermissionSetGroup"`
	Xmlns       string   `xml:"xmlns,attr"`
	Description struct {
		Text string `xml:",innerxml"`
	} `xml:"description"`
	HasActivationRequired struct {
		Text string `xml:",chardata"`
	} `xml:"hasActivationRequired"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	PermissionSets PermissionSetList `xml:"permissionSets"`
	Status         struct {
		Text string `xml:",chardata"`
	} `xml:"status"`
}

func (c *PermissionSetGroup) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *PermissionSetGroup) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*PermissionSetGroup, error) {
	p := &PermissionSetGroup{}
	return p, metadata.ParseMetadataXml(p, path)
}
