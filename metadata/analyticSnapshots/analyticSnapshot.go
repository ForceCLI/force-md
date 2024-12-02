package analyticSnapshot

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AnalyticSnapshot"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type AnalyticSnapshot struct {
	metadata.MetadataInfo
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

func (c *AnalyticSnapshot) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AnalyticSnapshot) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AnalyticSnapshot, error) {
	p := &AnalyticSnapshot{}
	return p, metadata.ParseMetadataXml(p, path)
}
