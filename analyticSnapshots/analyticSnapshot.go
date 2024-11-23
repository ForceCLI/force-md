package analyticSnapshot

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "AnalyticSnapshot"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type AnalyticSnapshot struct {
	internal.MetadataInfo
	XMLName  xml.Name `xml:"AnalyticSnapshot"`
	Xmlns    string   `xml:"xmlns,attr"`
	Mappings []struct {
		SourceField struct {
			Text string `xml:",chardata"`
		} `xml:"sourceField"`
		SourceType struct {
			Text string `xml:",chardata"`
		} `xml:"sourceType"`
		TargetField struct {
			Text string `xml:",chardata"`
		} `xml:"targetField"`
	} `xml:"mappings"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	RunningUser struct {
		Text string `xml:",chardata"`
	} `xml:"runningUser"`
	SourceReport struct {
		Text string `xml:",chardata"`
	} `xml:"sourceReport"`
	TargetObject struct {
		Text string `xml:",chardata"`
	} `xml:"targetObject"`
}

func (c *AnalyticSnapshot) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AnalyticSnapshot) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*AnalyticSnapshot, error) {
	p := &AnalyticSnapshot{}
	return p, internal.ParseMetadataXml(p, path)
}
