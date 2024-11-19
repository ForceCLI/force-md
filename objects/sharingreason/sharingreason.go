package sharingreason

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects/split"
)

const NAME = "SharingReason"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

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

func (c *SharingReasonMetadata) NameFromPath(path string) string {
	return split.NameFromPath(path)
}

func (c *SharingReasonMetadata) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*SharingReasonMetadata, error) {
	p := &SharingReasonMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
