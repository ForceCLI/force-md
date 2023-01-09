package custommetadata

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type Value struct {
	Field string `xml:"field"`
	Value struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"value"`
}

type CustomMetadata struct {
	XMLName   xml.Name `xml:"CustomMetadata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Xsi       string   `xml:"xsi,attr"`
	Xsd       string   `xml:"xsd,attr"`
	Label     string   `xml:"label"`
	Protected struct {
		Text string `xml:",chardata"`
	} `xml:"protected"`
	Values []Value `xml:"values"`
}

func (p *CustomMetadata) MetaCheck() {}

func Open(path string) (*CustomMetadata, error) {
	p := &CustomMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
