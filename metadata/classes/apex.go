package apexClass

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "ApexClass"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexClass struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ApexClass"`
	Xmlns      string   `xml:"xmlns,attr"`
	Fqn        *string  `xml:"fqn,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	PackageVersions []struct {
		MajorNumber struct {
			Text string `xml:",chardata"`
		} `xml:"majorNumber"`
		MinorNumber struct {
			Text string `xml:",chardata"`
		} `xml:"minorNumber"`
		Namespace struct {
			Text string `xml:",chardata"`
		} `xml:"namespace"`
	} `xml:"packageVersions"`
	Status struct {
		Text string `xml:",chardata"`
	} `xml:"status"`

	// Non-XML field to store the actual class code
	SourceCode []byte `xml:"-"`
	SourcePath string `xml:"-"`
}

func (c *ApexClass) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexClass) Type() metadata.MetadataType {
	return NAME
}

func (c *ApexClass) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the class name from metadata
	className := c.MetadataInfo.Name()
	if className == "" {
		return nil, fmt.Errorf("class name is empty")
	}

	// Get the directory name for classes
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Load source code if we haven't already
	if c.SourceCode == nil {
		c.SourceCode = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".cls")
	}

	// Marshal the XML metadata using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal class metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for ApexClass
		return helpers.GenerateMetadataFilePaths(dirName, string(className), ".cls-meta.xml", ".cls", xmlContent, c.SourceCode), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func Open(path string) (*ApexClass, error) {
	p := &ApexClass{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated code files
	p.SourcePath = path

	return p, nil
}
