package validationrule

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/objects/split"
)

const NAME = "ValidationRule"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type ValidationRule struct {
	internal.MetadataInfo
	XMLName xml.Name `xml:"ValidationRule"`
	Xmlns   string   `xml:"xmlns,attr"`
	Rule
}

type ValidationRuleList []Rule

type Rule struct {
	FullName string `xml:"fullName"`
	Active   struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	Description *struct {
		Text string `xml:",innerxml"`
	} `xml:"description"`
	ErrorConditionFormula *TextLiteral `xml:"errorConditionFormula"`
	ErrorDisplayField     *struct {
		Text string `xml:",chardata"`
	} `xml:"errorDisplayField"`
	ErrorMessage struct {
		Text string `xml:",innerxml"`
	} `xml:"errorMessage"`
}

func (c *ValidationRule) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ValidationRule) Type() internal.MetadataType {
	return NAME
}

func (c *ValidationRule) NameFromPath(path string) string {
	return split.NameFromPath(path)
}

func Open(path string) (*ValidationRule, error) {
	p := &ValidationRule{}
	return p, internal.ParseMetadataXml(p, path)
}
