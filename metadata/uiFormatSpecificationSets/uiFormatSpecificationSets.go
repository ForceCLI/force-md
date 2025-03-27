package uiFormatSpecificationSets

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "UiFormatSpecificationSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type UiFormatSpecificationSet struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"UiFormatSpecificationSet"`
	Xmlns   string   `xml:"xmlns,attr"`
	Field   struct {
		Text string `xml:",chardata"`
	} `xml:"field"`
	FormatType struct {
		Text string `xml:",chardata"`
	} `xml:"formatType"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	SobjectType struct {
		Text string `xml:",chardata"`
	} `xml:"sobjectType"`
	UiFormatSpecifications []struct {
		FormatProperties struct {
			Text string `xml:",chardata"`
		} `xml:"formatProperties"`
		FormatType struct {
			Text string `xml:",chardata"`
		} `xml:"formatType"`
		Order struct {
			Text string `xml:",chardata"`
		} `xml:"order"`
		VisibilityRule struct {
			Criteria []struct {
				LeftValue struct {
					Text string `xml:",chardata"`
				} `xml:"leftValue"`
				Operator struct {
					Text string `xml:",chardata"`
				} `xml:"operator"`
				RightValue struct {
					Text string `xml:",chardata"`
				} `xml:"rightValue"`
			} `xml:"criteria"`
			BooleanFilter struct {
				Text string `xml:",chardata"`
			} `xml:"booleanFilter"`
		} `xml:"visibilityRule"`
	} `xml:"uiFormatSpecifications"`
}

func (c *UiFormatSpecificationSet) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *UiFormatSpecificationSet) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*UiFormatSpecificationSet, error) {
	p := &UiFormatSpecificationSet{}
	return p, metadata.ParseMetadataXml(p, path)
}
