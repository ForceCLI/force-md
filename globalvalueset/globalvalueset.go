package globalvalueset

import (
	"encoding/xml"
	"sort"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type GlobalValueSet struct {
	XMLName     xml.Name `xml:"GlobalValueSet"`
	Xmlns       string   `xml:"xmlns,attr"`
	CustomValue []struct {
		FullName string       `xml:"fullName"`
		Default  BooleanText  `xml:"default"`
		Label    string       `xml:"label"`
		IsActive *BooleanText `xml:"isActive"`
	} `xml:"customValue"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Sorted *BooleanText `xml:"sorted"`
}

func (p *GlobalValueSet) MetaCheck() {}

func Open(path string) (*GlobalValueSet, error) {
	p := &GlobalValueSet{}
	return p, internal.ParseMetadataXml(p, path)
}

func (p *GlobalValueSet) Tidy() {
	if p.Sorted.Text == "true" {
		sort.Slice(p.CustomValue, func(i, j int) bool {
			return p.CustomValue[i].Label < p.CustomValue[j].Label
		})
	}
}
