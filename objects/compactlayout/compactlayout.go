package compactlayout

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type CompactLayoutMetadata struct {
	XMLName xml.Name `xml:"CompactLayout"`
	Xmlns   string   `xml:"xmlns,attr"`
	CompactLayout
}

type CompactLayout struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Fields []struct {
		Text string `xml:",chardata"`
	} `xml:"fields"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
}

func (p *CompactLayoutMetadata) MetaCheck() {}

func Open(path string) (*CompactLayoutMetadata, error) {
	p := &CompactLayoutMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
