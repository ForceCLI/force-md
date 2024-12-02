package sharingSet

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "SharingSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type SharingSet struct {
	metadata.MetadataInfo
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

func (c *SharingSet) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *SharingSet) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*SharingSet, error) {
	p := &SharingSet{}
	return p, metadata.ParseMetadataXml(p, path)
}
