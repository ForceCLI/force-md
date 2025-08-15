package delegateGroup

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DelegateGroup"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DelegateGroup struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"DelegateGroup"`
	Xmlns   string   `xml:"xmlns,attr"`
	Groups  *struct {
		Text string `xml:",chardata"`
	} `xml:"groups"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LoginAccess struct {
		Text string `xml:",chardata"`
	} `xml:"loginAccess"`
	Roles []struct {
		Text string `xml:",chardata"`
	} `xml:"roles"`
}

func (c *DelegateGroup) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DelegateGroup) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DelegateGroup, error) {
	p := &DelegateGroup{}
	return p, metadata.ParseMetadataXml(p, path)
}
