package permissionsetgroup

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "PermissionSetGroup"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type PermissionSet struct {
	Text string `xml:",chardata"`
}

type PermissionSetList []PermissionSet

type PermissionSetGroup struct {
	internal.MetadataInfo
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

func (c *PermissionSetGroup) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *PermissionSetGroup) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*PermissionSetGroup, error) {
	p := &PermissionSetGroup{}
	return p, internal.ParseMetadataXml(p, path)
}
