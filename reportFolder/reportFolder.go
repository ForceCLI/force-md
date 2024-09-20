package reportFolder

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type FolderShare struct {
	AccessLevel  string `xml:"accessLevel"`
	SharedTo     string `xml:"sharedTo"`
	SharedToType string `xml:"sharedToType"`
}

type ReportFolder struct {
	Metadata
	XMLName      xml.Name      `xml:"ReportFolder"`
	Xmlns        string        `xml:"xmlns,attr"`
	FolderShares []FolderShare `xml:"folderShares"`
	Name         string        `xml:"name"`
}

func (c *ReportFolder) SetMetadata(m Metadata) {
	c.Metadata = m
}

func Open(path string) (*ReportFolder, error) {
	p := &ReportFolder{}
	return p, internal.ParseMetadataXml(p, path)
}
