package territory2rule

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Territory2Rule"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RuleItem struct {
	Field struct {
		Text string `xml:",chardata"`
	} `xml:"field"`
	Operation struct {
		Text string `xml:",chardata"`
	} `xml:"operation"`
	Value struct {
		Text string `xml:",chardata"`
	} `xml:"value"`
}

type Territory2Rule struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Territory2Rule"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	BooleanFilter *struct {
		Text string `xml:",chardata"`
	} `xml:"booleanFilter"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	ObjectType struct {
		Text string `xml:",chardata"`
	} `xml:"objectType"`
	RuleItems []RuleItem `xml:"ruleItems"`
}

func (c *Territory2Rule) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Territory2Rule) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Territory2Rule, error) {
	p := &Territory2Rule{}
	return p, metadata.ParseMetadataXml(p, path)
}
