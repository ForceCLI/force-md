package sharingSet

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "SharingSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type SharingSet struct {
	internal.MetadataInfo
	XMLName        xml.Name `xml:"SharingSet"`
	Xmlns          string   `xml:"xmlns,attr"`
	AccessMappings []struct {
		AccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"accessLevel"`
		Object struct {
			Text string `xml:",chardata"`
		} `xml:"object"`
		ObjectField struct {
			Text string `xml:",chardata"`
		} `xml:"objectField"`
		UserField struct {
			Text string `xml:",chardata"`
		} `xml:"userField"`
	} `xml:"accessMappings"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Profiles struct {
		Text string `xml:",chardata"`
	} `xml:"profiles"`
}

func (c *SharingSet) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *SharingSet) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*SharingSet, error) {
	p := &SharingSet{}
	return p, internal.ParseMetadataXml(p, path)
}
