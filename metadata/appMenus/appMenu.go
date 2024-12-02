package appMenu

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AppMenu"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type AppMenu struct {
	metadata.MetadataInfo
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

func (c *AppMenu) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AppMenu) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AppMenu, error) {
	p := &AppMenu{}
	return p, metadata.ParseMetadataXml(p, path)
}
