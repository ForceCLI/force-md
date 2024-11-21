package appMenu

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "AppMenu"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type AppMenu struct {
	internal.MetadataInfo
	XMLName      xml.Name `xml:"AppMenu"`
	Xmlns        string   `xml:"xmlns,attr"`
	AppMenuItems []struct {
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
	} `xml:"appMenuItems"`
}

func (c *AppMenu) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AppMenu) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*AppMenu, error) {
	p := &AppMenu{}
	return p, internal.ParseMetadataXml(p, path)
}
