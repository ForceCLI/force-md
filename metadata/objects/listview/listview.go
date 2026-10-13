package listview

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objects/split"
)

const NAME = "ListView"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ListViewMetadata struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"ListView"`
	Xmlns   string   `xml:"xmlns,attr"`
	ListView
}

type ListView struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	BooleanFilter *struct {
		Text string `xml:",chardata"`
	} `xml:"booleanFilter"`
	Columns []struct {
		Text string `xml:",chardata"`
	} `xml:"columns"`
	FilterScope struct {
		Text string `xml:",chardata"`
	} `xml:"filterScope"`
	Filters []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Value *struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"filters"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Queue *struct {
		Text string `xml:",chardata"`
	} `xml:"queue"`
	SharedTo *SharedTo `xml:"sharedTo"`
	Language *struct {
		Text string `xml:",chardata"`
	} `xml:"language"`
}

func (c *ListViewMetadata) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ListViewMetadata) NameFromPath(path string) metadata.MetadataObjectName {
	return split.NameFromPath(path)
}

func (c *ListViewMetadata) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ListViewMetadata, error) {
	p := &ListViewMetadata{}
	return p, metadata.ParseMetadataXml(p, path)
}

// SharedTo lists who a list view is shared with.
type SharedTo struct {
	AllCustomerPortalUsers      *TextLiteral  `xml:"allCustomerPortalUsers"`
	AllInternalUsers            *TextLiteral  `xml:"allInternalUsers"`
	AllPartnerUsers             *TextLiteral  `xml:"allPartnerUsers"`
	ChannelProgramGroup         []TextLiteral `xml:"channelProgramGroup"`
	ChannelProgramGroups        []TextLiteral `xml:"channelProgramGroups"`
	Group                       []TextLiteral `xml:"group"`
	Groups                      []TextLiteral `xml:"groups"`
	GuestUser                   []TextLiteral `xml:"guestUser"`
	ManagerSubordinates         []TextLiteral `xml:"managerSubordinates"`
	Managers                    []TextLiteral `xml:"managers"`
	PortalRole                  []TextLiteral `xml:"portalRole"`
	PortalRoleAndSubordinates   []TextLiteral `xml:"portalRoleAndSubordinates"`
	Queue                       []TextLiteral `xml:"queue"`
	Role                        []TextLiteral `xml:"role"`
	RoleAndSubordinates         []TextLiteral `xml:"roleAndSubordinates"`
	RoleAndSubordinatesInternal []TextLiteral `xml:"roleAndSubordinatesInternal"`
	Territories                 []TextLiteral `xml:"territories"`
	TerritoriesAndSubordinates  []TextLiteral `xml:"territoriesAndSubordinates"`
	Territory                   []TextLiteral `xml:"territory"`
	TerritoryAndSubordinates    []TextLiteral `xml:"territoryAndSubordinates"`
}
