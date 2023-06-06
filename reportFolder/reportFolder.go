package reportFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type FolderShare struct {
	AccessLevel  string `xml:"accessLevel"`
	SharedTo     string `xml:"sharedTo"`
	SharedToType string `xml:"sharedToType"`
}

type ReportFolder struct {
	XMLName      xml.Name      `xml:"ReportFolder"`
	Xmlns        string        `xml:"xmlns,attr"`
	Name         string        `xml:"name"`
	FolderShares []FolderShare `xml:"folderShares"`
}

func (p *ReportFolder) MetaCheck() {}

func Open(path string) (*ReportFolder, error) {
	p := &ReportFolder{}
	return p, internal.ParseMetadataXml(p, path)
}
