package assignmentRules

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "AssignmentRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type AssignmentRules struct {
	metadata.MetadataInfo
	XMLName        xml.Name `xml:"AssignmentRules"`
	Xmlns          string   `xml:"xmlns,attr"`
	AssignmentRule []struct {
		FullName struct {
			Text string `xml:",chardata"`
		} `xml:"fullName"`
		Active struct {
			Text string `xml:",chardata"`
		} `xml:"active"`
		RuleEntry []struct {
			AssignedTo struct {
				Text string `xml:",chardata"`
			} `xml:"assignedTo"`
			AssignedToType struct {
				Text string `xml:",chardata"`
			} `xml:"assignedToType"`
			CriteriaItems struct {
				Field struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
				Operation struct {
					Text string `xml:",chardata"`
				} `xml:"operation"`
				Value struct {
					Text string `xml:",chardata"`
				} `xml:"value"`
			} `xml:"criteriaItems"`
			Template struct {
				Text string `xml:",chardata"`
			} `xml:"template"`
		} `xml:"ruleEntry"`
	} `xml:"assignmentRule"`
}

func (c *AssignmentRules) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AssignmentRules) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*AssignmentRules, error) {
	p := &AssignmentRules{}
	return p, metadata.ParseMetadataXml(p, path)
}
