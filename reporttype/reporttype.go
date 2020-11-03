package reporttype

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type ReportType struct {
	Declaration string   `xml:"-"`
	XMLName     xml.Name `xml:"ReportType"`
	Xmlns       string   `xml:"xmlns,attr"`
	BaseObject  struct {
		Text string `xml:",chardata"`
	} `xml:"baseObject"`
	Category struct {
		Text string `xml:",chardata"`
	} `xml:"category"`
	Deployed struct {
		Text string `xml:",chardata"`
	} `xml:"deployed"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Join struct {
		Join struct {
			Join struct {
				OuterJoin struct {
					Text string `xml:",chardata"`
				} `xml:"outerJoin"`
				Relationship struct {
					Text string `xml:",chardata"`
				} `xml:"relationship"`
			} `xml:"join"`
			OuterJoin struct {
				Text string `xml:",chardata"`
			} `xml:"outerJoin"`
			Relationship struct {
				Text string `xml:",chardata"`
			} `xml:"relationship"`
		} `xml:"join"`
		OuterJoin struct {
			Text string `xml:",chardata"`
		} `xml:"outerJoin"`
		Relationship struct {
			Text string `xml:",chardata"`
		} `xml:"relationship"`
	} `xml:"join"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Sections []struct {
		Columns []struct {
			CheckedByDefault struct {
				Text string `xml:",chardata"`
			} `xml:"checkedByDefault"`
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Table struct {
				Text string `xml:",chardata"`
			} `xml:"table"`
		} `xml:"columns"`
		MasterLabel struct {
			Text string `xml:",chardata"`
		} `xml:"masterLabel"`
	} `xml:"sections"`
}

func (p *ReportType) MetaCheck() {}

func Open(path string) (*ReportType, error) {
	p := &ReportType{}
	return p, internal.ParseMetadataXml(p, path)
}
