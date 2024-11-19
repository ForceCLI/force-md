package dashboardFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "DashboardFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type DashboardFolder struct {
	internal.MetadataInfo
	XMLName      xml.Name `xml:"DashboardFolder"`
	Xmlns        string   `xml:"xmlns,attr"`
	FolderShares []struct {
		AccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"accessLevel"`
		SharedTo struct {
			Text string `xml:",chardata"`
		} `xml:"sharedTo"`
		SharedToType struct {
			Text string `xml:",chardata"`
		} `xml:"sharedToType"`
	} `xml:"folderShares"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

func (c *DashboardFolder) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DashboardFolder) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*DashboardFolder, error) {
	p := &DashboardFolder{}
	return p, internal.ParseMetadataXml(p, path)
}
