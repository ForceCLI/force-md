package pkg

import (
	"encoding/xml"
	"sort"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type Member string

func (n Member) GetName() string {
	return string(n)
}

type MetadataItems struct {
	Comment string   `xml:",comment"`
	Members []Member `xml:"members"`
	Name    string   `xml:"name"`
}

type Package struct {
	XMLName xml.Name        `xml:"Package"`
	Xmlns   string          `xml:"xmlns,attr"`
	Types   []MetadataItems `xml:"types"`
	Version string          `xml:"version"`
}

func NewPackage(version string) Package {
	p := Package{
		Version: version,
		Xmlns:   "http://soap.sforce.com/2006/04/metadata",
	}
	return p
}

func (p *Package) MetaCheck() {}

func Open(path string) (*Package, error) {
	p := &Package{}
	return p, internal.ParseMetadataXml(p, path)
}

func (p *Package) Tidy() {
	sort.Slice(p.Types, func(i, j int) bool {
		return p.Types[i].Name < p.Types[j].Name
	})
	for i := range p.Types {
		p.Types[i].Tidy()
	}
}

func (members *MetadataItems) Tidy() {
	sort.Slice(members.Members, func(i, j int) bool {
		return members.Members[i] < members.Members[j]
	})
	RemoveDuplicates(&members.Members)
}
