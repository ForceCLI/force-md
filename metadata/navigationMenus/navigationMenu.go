package navigationMenus

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "NavigationMenu"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type NavigationMenu struct {
	metadata.MetadataInfo
	XMLName             xml.Name             `xml:"NavigationMenu"`
	Xmlns               string               `xml:"xmlns,attr"`
	Container           *TextLiteral         `xml:"container"`
	ContainerType       *TextLiteral         `xml:"containerType"`
	Label               *TextLiteral         `xml:"label"`
	NavigationMenuItems []NavigationMenuItem `xml:"navigationMenuItem"`
}

type NavigationMenuItem struct {
	DefaultListViewId *TextLiteral `xml:"defaultListViewId"`
	Label             *TextLiteral `xml:"label"`
	MenuItemBranding  *struct {
		TileImage *TextLiteral `xml:"tileImage"`
	} `xml:"menuItemBranding"`
	Position          *TextLiteral       `xml:"position"`
	PubliclyAvailable *TextLiteral       `xml:"publiclyAvailable"`
	SubMenu           *NavigationSubMenu `xml:"subMenu"`
	Target            *TextLiteral       `xml:"target"`
	TargetPreference  *TextLiteral       `xml:"targetPreference"`
	Type              *TextLiteral       `xml:"type"`
}

type NavigationSubMenu struct {
	NavigationMenuItems []NavigationMenuItem `xml:"navigationMenuItem"`
}

func (n *NavigationMenu) SetMetadata(m metadata.MetadataInfo) {
	n.MetadataInfo = m
}

func (n *NavigationMenu) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*NavigationMenu, error) {
	n := &NavigationMenu{}
	return n, metadata.ParseMetadataXml(n, path)
}
