package customHelpMenuSection

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "CustomHelpMenuSection"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type CustomHelpMenuSection struct {
	internal.MetadataInfo
	XMLName             xml.Name `xml:"CustomHelpMenuSection"`
	Xmlns               string   `xml:"xmlns,attr"`
	CustomHelpMenuItems []struct {
		LinkUrl struct {
			Text string `xml:",chardata"`
		} `xml:"linkUrl"`
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
		SortOrder struct {
			Text string `xml:",chardata"`
		} `xml:"sortOrder"`
	} `xml:"customHelpMenuItems"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
}

func (c *CustomHelpMenuSection) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomHelpMenuSection) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*CustomHelpMenuSection, error) {
	p := &CustomHelpMenuSection{}
	return p, internal.ParseMetadataXml(p, path)
}
