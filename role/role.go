package role

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "Role"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type Role struct {
	internal.MetadataInfo
	XMLName         xml.Name `xml:"Role"`
	Xmlns           string   `xml:"xmlns,attr"`
	CaseAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"caseAccessLevel"`
	ContactAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"contactAccessLevel"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MayForecastManagerShare struct {
		Text string `xml:",chardata"`
	} `xml:"mayForecastManagerShare"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	OpportunityAccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"opportunityAccessLevel"`
	ParentRole struct {
		Text string `xml:",chardata"`
	} `xml:"parentRole"`
}

func (c *Role) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Role) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Role, error) {
	p := &Role{}
	return p, internal.ParseMetadataXml(p, path)
}