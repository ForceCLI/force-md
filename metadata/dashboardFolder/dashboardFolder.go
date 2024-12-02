package dashboardFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DashboardFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DashboardFolder struct {
	metadata.MetadataInfo
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

func (c *DashboardFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DashboardFolder) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DashboardFolder, error) {
	p := &DashboardFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
