package components

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "ApexComponent"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexComponent struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ApexComponent"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`

	// Non-XML fields
	ComponentCode []byte `xml:"-"` // The actual component code
	SourcePath    string `xml:"-"`
}

func (c *ApexComponent) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexComponent) Type() metadata.MetadataType {
	return NAME
}

func (c *ApexComponent) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the component name from metadata
	componentName := c.MetadataInfo.Name()
	if componentName == "" {
		return nil, fmt.Errorf("component name is empty")
	}

	// Get the directory name for components
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Load component code if we haven't already
	if c.ComponentCode == nil {
		c.ComponentCode = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".component")
	}

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal component metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for ApexComponent
		return helpers.GenerateMetadataFilePaths(dirName, string(componentName), ".component-meta.xml", ".component", xmlContent, c.ComponentCode), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func Open(path string) (*ApexComponent, error) {
	p := &ApexComponent{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated code files
	p.SourcePath = path

	return p, nil
}
