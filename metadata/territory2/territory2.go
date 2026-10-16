package territory2

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/territory2model"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "Territory2"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type RuleAssociation struct {
	Inherited struct {
		Text string `xml:",chardata"`
	} `xml:"inherited"`
	RuleName struct {
		Text string `xml:",chardata"`
	} `xml:"ruleName"`
}

// ObjectAccessLevel is the access that users in a territory have to records
// of an object, such as Lead, assigned to the territory.
type ObjectAccessLevel struct {
	AccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"accessLevel"`
	ObjectType struct {
		Text string `xml:",chardata"`
	} `xml:"objectType"`
}

type Territory2 struct {
	metadata.MetadataInfo
	XMLName            xml.Name `xml:"Territory2"`
	Xmlns              string   `xml:"xmlns,attr"`
	Xsi                string   `xml:"xmlns:xsi,attr,omitempty"`
	Xsd                string   `xml:"xmlns:xsd,attr,omitempty"`
	AccountAccessLevel *struct {
		Text string `xml:",chardata"`
	} `xml:"accountAccessLevel"`
	CaseAccessLevel *struct {
		Text string `xml:",chardata"`
	} `xml:"caseAccessLevel"`
	ContactAccessLevel *struct {
		Text string `xml:",chardata"`
	} `xml:"contactAccessLevel"`
	CustomFields []territory2model.FieldValue `xml:"customFields"`
	Description  *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	ObjectAccessLevels     []ObjectAccessLevel `xml:"objectAccessLevels"`
	OpportunityAccessLevel *struct {
		Text string `xml:",chardata"`
	} `xml:"opportunityAccessLevel"`
	ParentTerritory *struct {
		Text string `xml:",chardata"`
	} `xml:"parentTerritory"`
	RuleAssociations []RuleAssociation `xml:"ruleAssociations"`
	Territory2Type   struct {
		Text string `xml:",chardata"`
	} `xml:"territory2Type"`
}

func (c *Territory2) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Territory2) Type() metadata.MetadataType {
	return NAME
}

// NameFromPath names a territory as package.xml does, <Model>.<Territory>,
// since territories of different models may share a name.
func (c *Territory2) NameFromPath(path string) metadata.MetadataObjectName {
	return territory2model.ComponentName(path)
}

// Files writes the territory into its model's directory:
// territory2Models/<Model>/territories/<Territory>.territory2.
func (c *Territory2) Files(format metadata.Format) (map[string][]byte, error) {
	path := string(c.MetadataInfo.Path())
	return territory2model.ModelComponentFiles(c, territory2model.ModelName(path), registry.GetMetadataDirectory(NAME), string(metadata.NameFromPath(path)), NAME, format)
}

func Open(path string) (*Territory2, error) {
	p := &Territory2{}
	return p, metadata.ParseMetadataXml(p, path)
}
