package duplicateRule

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "DuplicateRule"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DuplicateRule struct {
	metadata.MetadataInfo
	XMLName             xml.Name      `xml:"DuplicateRule"`
	Xmlns               string        `xml:"xmlns,attr"`
	Xsi                 string        `xml:"xmlns:xsi,attr,omitempty"`
	ActionOnInsert      TextLiteral   `xml:"actionOnInsert"`
	ActionOnUpdate      TextLiteral   `xml:"actionOnUpdate"`
	AlertText           *NillableText `xml:"alertText"`
	Description         *NillableText `xml:"description"`
	DuplicateRuleFilter *struct {
		Nil           string        `xml:"xsi:nil,attr,omitempty"`
		BooleanFilter *NillableText `xml:"booleanFilter"`
		Items         []struct {
			Field     TextLiteral `xml:"field"`
			Operation TextLiteral `xml:"operation"`
			Value     TextLiteral `xml:"value"`
			SortOrder TextLiteral `xml:"sortOrder"`
			Table     TextLiteral `xml:"table"`
		} `xml:"duplicateRuleFilterItems"`
	} `xml:"duplicateRuleFilter"`
	DuplicateRuleMatchRules []struct {
		MatchRuleSObjectType TextLiteral `xml:"matchRuleSObjectType"`
		MatchingRule         TextLiteral `xml:"matchingRule"`
		ObjectMapping        *struct {
			Nil           string       `xml:"xsi:nil,attr,omitempty"`
			InputObject   *TextLiteral `xml:"inputObject"`
			MappingFields []struct {
				InputField  TextLiteral `xml:"inputField"`
				OutputField TextLiteral `xml:"outputField"`
			} `xml:"mappingFields"`
			OutputObject *TextLiteral `xml:"outputObject"`
		} `xml:"objectMapping"`
	} `xml:"duplicateRuleMatchRules"`
	IsActive           TextLiteral   `xml:"isActive"`
	MasterLabel        TextLiteral   `xml:"masterLabel"`
	OperationsOnInsert []TextLiteral `xml:"operationsOnInsert"`
	OperationsOnUpdate []TextLiteral `xml:"operationsOnUpdate"`
	SecurityOption     TextLiteral   `xml:"securityOption"`
	SortOrder          TextLiteral   `xml:"sortOrder"`
}

func (c *DuplicateRule) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

// KeepEmptyElements marks duplicate rules as written with empty elements
// open, as Salesforce writes <value></value> in filter items.
func (c *DuplicateRule) KeepEmptyElements() {}

func (c *DuplicateRule) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*DuplicateRule, error) {
	p := &DuplicateRule{}
	return p, metadata.ParseMetadataXml(p, path)
}
