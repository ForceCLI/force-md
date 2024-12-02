package reportFolder

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ReportFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FolderShare struct {
	AccessLevel  string `xml:"accessLevel"`
	SharedTo     string `xml:"sharedTo"`
	SharedToType string `xml:"sharedToType"`
}

type ReportFolder struct {
	metadata.MetadataInfo
	XMLName      xml.Name      `xml:"ReportFolder"`
	Xmlns        string        `xml:"xmlns,attr"`
	FolderShares []FolderShare `xml:"folderShares"`
	Name         string        `xml:"name"`
}

func (c *ReportFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ReportFolder) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ReportFolder, error) {
	p := &ReportFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
