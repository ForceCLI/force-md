package sharingreason

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type SharingReasonMetadata struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"SharingReason"`
	Xmlns   string   `xml:"xmlns,attr"`
	SharingReason
}

type SharingReason struct {
	FullName string `xml:"fullName"`
	Label    string `xml:"label"`
}

func (c *SharingReasonMetadata) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*SharingReasonMetadata, error) {
	p := &SharingReasonMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
