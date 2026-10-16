package territory2rule

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/territory2model"
	"github.com/ForceCLI/force-md/registry"
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
	Value *struct {
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

// NameFromPath names a rule as package.xml does, <Model>.<Rule>, since rules
// of different models may share a name.
func (c *Territory2Rule) NameFromPath(path string) metadata.MetadataObjectName {
	return territory2model.ComponentName(path)
}

// Files writes the rule into its model's directory:
// territory2Models/<Model>/rules/<Rule>.territory2Rule.
func (c *Territory2Rule) Files(format metadata.Format) (map[string][]byte, error) {
	path := string(c.MetadataInfo.Path())
	return territory2model.ModelComponentFiles(c, territory2model.ModelName(path), registry.GetMetadataDirectory(NAME), string(metadata.NameFromPath(path)), NAME, format)
}

func Open(path string) (*Territory2Rule, error) {
	p := &Territory2Rule{}
	return p, metadata.ParseMetadataXml(p, path)
}
