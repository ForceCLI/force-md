package businessprocess

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

type BusinessProcessMetadata struct {
	XMLName xml.Name `xml:"BusinessProcess"`
	Xmlns   string   `xml:"xmlns,attr"`
	BusinessProcess
}

type BusinessProcess struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	IsActive struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	Values []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Default struct {
			Text string `xml:",chardata"`
		} `xml:"default"`
	} `xml:"values"`
}

func (p *BusinessProcessMetadata) MetaCheck() {}

func Open(path string) (*BusinessProcessMetadata, error) {
	p := &BusinessProcessMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
