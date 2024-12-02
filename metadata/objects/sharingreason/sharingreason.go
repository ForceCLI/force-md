package sharingreason

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "SharingReason"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type SharingReasonMetadata struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"SharingReason"`
	Xmlns   string   `xml:"xmlns,attr"`
	SharingReason
}

type SharingReason struct {
	FullName string `xml:"fullName"`
	Label    string `xml:"label"`
}

func (c *SharingReasonMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *SharingReasonMetadata) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *SharingReasonMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*SharingReasonMetadata, error) {
	p := &SharingReasonMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}
