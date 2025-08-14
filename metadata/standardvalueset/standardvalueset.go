package standardvalueset

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "StandardValueSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type StandardValue struct {
	metadata.MetadataInfo
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	Default struct {
		Text string `xml:",chardata"`
	} `xml:"default"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	CssExposed *struct {
		Text string `xml:",chardata"`
	} `xml:"cssExposed"`
	Closed *struct {
		Text string `xml:",chardata"`
	} `xml:"closed"`
	GroupingString *struct {
		Text string `xml:",chardata"`
	} `xml:"groupingString"`
	Converted *struct {
		Text string `xml:",chardata"`
	} `xml:"converted"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	ForecastCategory *struct {
		Text string `xml:",chardata"`
	} `xml:"forecastCategory"`
	Probability *struct {
		Text string `xml:",chardata"`
	} `xml:"probability"`
	Won *struct {
		Text string `xml:",chardata"`
	} `xml:"won"`
	ReverseRole *struct {
		Text string `xml:",chardata"`
	} `xml:"reverseRole"`
	AllowEmail *struct {
		Text string `xml:",chardata"`
	} `xml:"allowEmail"`
	HighPriority *struct {
		Text string `xml:",chardata"`
	} `xml:"highPriority"`
}

type StandardValueSet struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"StandardValueSet"`
	Xmlns   string   `xml:"xmlns,attr"`
	Sorted  struct {
		Text string `xml:",chardata"`
	} `xml:"sorted"`
	StandardValue      []StandardValue `xml:"standardValue"`
	GroupingStringEnum *struct {
		Text string `xml:",chardata"`
	} `xml:"groupingStringEnum"`
}

func (c *StandardValueSet) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*StandardValueSet, error) {
	p := &StandardValueSet{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *StandardValueSet) Type() metadata.MetadataType {
	return NAME
}

func (s *StandardValueSet) GetValues() []StandardValue {
	return s.StandardValue
}
