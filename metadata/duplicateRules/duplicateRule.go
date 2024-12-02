package duplicateRule

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DuplicateRule"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DuplicateRule struct {
	metadata.MetadataInfo
	XMLName        xml.Name `xml:"DuplicateRule"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	ActionOnInsert struct {
		Text string `xml:",chardata"`
	} `xml:"actionOnInsert"`
	ActionOnUpdate struct {
		Text string `xml:",chardata"`
	} `xml:"actionOnUpdate"`
	AlertText struct {
		Text string `xml:",chardata"`
	} `xml:"alertText"`
	Description struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"description"`
	DuplicateRuleFilter struct {
		Nil           string `xml:"nil,attr"`
		BooleanFilter struct {
			Nil string `xml:"nil,attr"`
		} `xml:"booleanFilter"`
		DuplicateRuleFilterItems []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operation struct {
				Text string `xml:",chardata"`
			} `xml:"operation"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
			SortOrder struct {
				Text string `xml:",chardata"`
			} `xml:"sortOrder"`
			Table struct {
				Text string `xml:",chardata"`
			} `xml:"table"`
		} `xml:"duplicateRuleFilterItems"`
	} `xml:"duplicateRuleFilter"`
	DuplicateRuleMatchRules struct {
		MatchRuleSObjectType struct {
			Text string `xml:",chardata"`
		} `xml:"matchRuleSObjectType"`
		MatchingRule struct {
			Text string `xml:",chardata"`
		} `xml:"matchingRule"`
		ObjectMapping struct {
			Nil         string `xml:"nil,attr"`
			InputObject struct {
				Text string `xml:",chardata"`
			} `xml:"inputObject"`
			MappingFields []struct {
				InputField struct {
					Text string `xml:",chardata"`
				} `xml:"inputField"`
				OutputField struct {
					Text string `xml:",chardata"`
				} `xml:"outputField"`
			} `xml:"mappingFields"`
			OutputObject struct {
				Text string `xml:",chardata"`
			} `xml:"outputObject"`
		} `xml:"objectMapping"`
	} `xml:"duplicateRuleMatchRules"`
	IsActive struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	OperationsOnUpdate []struct {
		Text string `xml:",chardata"`
	} `xml:"operationsOnUpdate"`
	SecurityOption struct {
		Text string `xml:",chardata"`
	} `xml:"securityOption"`
	SortOrder struct {
		Text string `xml:",chardata"`
	} `xml:"sortOrder"`
	OperationsOnInsert []struct {
		Text string `xml:",chardata"`
	} `xml:"operationsOnInsert"`
}

func (c *DuplicateRule) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DuplicateRule) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DuplicateRule, error) {
	p := &DuplicateRule{}
	return p, metadata.ParseMetadataXml(p, path)
}
