package permissionsetgroup

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type PermissionSet struct {
	Text string `xml:",chardata"`
}

type PermissionSetList []PermissionSet

type PermissionSetGroup struct {
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

func (p *PermissionSetGroup) MetaCheck() {}

func Open(path string) (*PermissionSetGroup, error) {
	p := &PermissionSetGroup{}
	return p, internal.ParseMetadataXml(p, path)
}
