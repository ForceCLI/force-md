package standardvalueset

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type StandardValue struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Default struct {
		Text string `xml:",chardata"`
	} `xml:"default"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	CssExposed struct {
		Text string `xml:",chardata"`
	} `xml:"cssExposed"`
	Closed struct {
		Text string `xml:",chardata"`
	} `xml:"closed"`
	GroupingString struct {
		Text string `xml:",chardata"`
	} `xml:"groupingString"`
	Converted struct {
		Text string `xml:",chardata"`
	} `xml:"converted"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	ForecastCategory struct {
		Text string `xml:",chardata"`
	} `xml:"forecastCategory"`
	Probability struct {
		Text string `xml:",chardata"`
	} `xml:"probability"`
	Won struct {
		Text string `xml:",chardata"`
	} `xml:"won"`
	ReverseRole struct {
		Text string `xml:",chardata"`
	} `xml:"reverseRole"`
	AllowEmail struct {
		Text string `xml:",chardata"`
	} `xml:"allowEmail"`
	HighPriority struct {
		Text string `xml:",chardata"`
	} `xml:"highPriority"`
}

type StandardValueSet struct {
	XMLName xml.Name `xml:"StandardValueSet"`
	Xmlns   string   `xml:"xmlns,attr"`
	Sorted  struct {
		Text string `xml:",chardata"`
	} `xml:"sorted"`
	StandardValue      []StandardValue `xml:"standardValue"`
	GroupingStringEnum struct {
		Text string `xml:",chardata"`
	} `xml:"groupingStringEnum"`
}

func (p *StandardValueSet) MetaCheck() {}

func Open(path string) (*StandardValueSet, error) {
	p := &StandardValueSet{}
	return p, internal.ParseMetadataXml(p, path)
}

func (s *StandardValueSet) GetValues() []StandardValue {
	return s.StandardValue
}
