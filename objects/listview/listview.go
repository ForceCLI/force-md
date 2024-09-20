package listview

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type ListViewMetadata struct {
	Metadata
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
	SharedTo *struct {
		Group []struct {
			Text string `xml:",chardata"`
		} `xml:"group"`
		Rol []struct {
			Text string `xml:",chardata"`
		} `xml:"role"`
		RoleAndSubordinates []struct {
			Text string `xml:",chardata"`
		} `xml:"roleAndSubordinates"`
		AllInternalUsers *struct {
			Text string `xml:",chardata"`
		} `xml:"allInternalUsers"`
	} `xml:"sharedTo"`
	Language *struct {
		Text string `xml:",chardata"`
	} `xml:"language"`
}

func (c *ListViewMetadata) SetMetadata(m Metadata) {
	c.Metadata = m
}

func Open(path string) (*ListViewMetadata, error) {
	p := &ListViewMetadata{}
	return p, internal.ParseMetadataXml(p, path)
}
