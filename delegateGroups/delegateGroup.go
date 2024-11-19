package delegateGroup

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "DelegateGroup"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type DelegateGroup struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"DelegateGroup"`
	Xmlns   string   `xml:"xmlns,attr"`
	Label   struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LoginAccess struct {
		Text string `xml:",chardata"`
	} `xml:"loginAccess"`
	Roles []struct {
		Text string `xml:",chardata"`
	} `xml:"roles"`
	Groups struct {
		Text string `xml:",chardata"`
	} `xml:"groups"`
}

func (c *DelegateGroup) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DelegateGroup) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*DelegateGroup, error) {
	p := &DelegateGroup{}
	return p, internal.ParseMetadataXml(p, path)
}
