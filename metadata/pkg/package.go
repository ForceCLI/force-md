package pkg

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"sort"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Package"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

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
	metadata.MetadataInfo
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

func (c *Package) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*Package, error) {
	p := &Package{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *Package) Type() metadata.MetadataType {
	return NAME
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

// Files implements the FilesGenerator interface
func (p *Package) Files(format metadata.Format) (map[string][]byte, error) {
	// Marshal the package to XML
	xmlContent, err := internal.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal package: %w", err)
	}

	files := make(map[string][]byte)

	// Get the original path to determine the package file name
	originalPath := string(p.MetadataInfo.Path())
	fileName := ""
	if originalPath != "" {
		fileName = filepath.Base(originalPath)
	}

	// If no path info or generic name, default to package.xml
	if fileName == "" || fileName == "Package" {
		fileName = "package.xml"
	}

	// Package files go in the root directory for both formats
	files[fileName] = xmlContent

	return files, nil
}
