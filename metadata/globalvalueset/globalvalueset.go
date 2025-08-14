package globalvalueset

import (
	"encoding/xml"
	"sort"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "GlobalValueSet"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ValueFilter func(CustomValue) bool

type CustomValue struct {
	FullName string       `xml:"fullName"`
	Default  BooleanText  `xml:"default"`
	IsActive *BooleanText `xml:"isActive"`
	Label    string       `xml:"label"`
}

type GlobalValueSet struct {
	metadata.MetadataInfo
	XMLName     xml.Name      `xml:"GlobalValueSet"`
	Xmlns       string        `xml:"xmlns,attr"`
	CustomValue []CustomValue `xml:"customValue"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Sorted *BooleanText `xml:"sorted"`
}

func (c *GlobalValueSet) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*GlobalValueSet, error) {
	p := &GlobalValueSet{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *GlobalValueSet) Type() metadata.MetadataType {
	return NAME
}

func (p *GlobalValueSet) Tidy() {
	if p.Sorted != nil && p.Sorted.ToBool() {
		sort.Slice(p.CustomValue, func(i, j int) bool {
			return p.CustomValue[i].Label < p.CustomValue[j].Label
		})
	}
}
