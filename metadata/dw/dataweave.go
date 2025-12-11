package dataweave

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "DataWeaveResource"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DataWeaveResource struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"DataWeaveResource"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	IsGlobal struct {
		Text string `xml:",chardata"`
	} `xml:"isGlobal"`

	// Non-XML field to store the actual DataWeave code
	SourceCode []byte `xml:"-"`
	SourcePath string `xml:"-"`
}

func (c *DataWeaveResource) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DataWeaveResource) Type() metadata.MetadataType {
	return NAME
}

func (c *DataWeaveResource) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the resource name from metadata
	resourceName := c.MetadataInfo.Name()
	if resourceName == "" {
		return nil, fmt.Errorf("resource name is empty")
	}

	// Get the directory name for DataWeave resources
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Load source code if we haven't already
	if c.SourceCode == nil {
		c.SourceCode = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".dwl")
	}

	// Marshal the XML metadata using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DataWeave resource metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for DataWeaveResource
		return helpers.GenerateMetadataFilePaths(dirName, string(resourceName), ".dwl-meta.xml", ".dwl", xmlContent, c.SourceCode), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func Open(path string) (*DataWeaveResource, error) {
	p := &DataWeaveResource{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated code files
	p.SourcePath = path

	return p, nil
}
