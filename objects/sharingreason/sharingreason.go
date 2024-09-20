package sharingreason

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type SharingReasonMetadata struct {
	Metadata
	XMLName xml.Name `xml:"SharingReason"`
	Xmlns   string   `xml:"xmlns,attr"`
	SharingReason
}

type SharingReason struct {
	FullName string `xml:"fullName"`
	Label    string `xml:"label"`
}

func (c *SharingReasonMetadata) SetMetadata(m Metadata) {
	c.Metadata = m
}

func Open(path string) (*SharingReasonMetadata, error) {
	p := &SharingReasonMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
