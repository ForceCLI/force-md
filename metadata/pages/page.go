package apexPage

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "ApexPage"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexPage struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ApexPage"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	AvailableInTouch *struct {
		Text string `xml:",chardata"`
	} `xml:"availableInTouch"`
	ConfirmationTokenRequired *struct {
		Text string `xml:",chardata"`
	} `xml:"confirmationTokenRequired"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`

	// Non-XML fields
	PageCode   []byte `xml:"-"` // The actual page code
	SourcePath string `xml:"-"`
}

func (c *ApexPage) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexPage) Type() metadata.MetadataType {
	return NAME
}

func (c *ApexPage) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the page name from metadata
	pageName := c.MetadataInfo.Name()
	if pageName == "" {
		return nil, fmt.Errorf("page name is empty")
	}

	// Get the directory name for pages
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Load page code if we haven't already
	if c.PageCode == nil {
		c.PageCode = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".page")
	}

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal page metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for ApexPage
		return helpers.GenerateMetadataFilePaths(dirName, string(pageName), ".page-meta.xml", ".page", xmlContent, c.PageCode), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func Open(path string) (*ApexPage, error) {
	p := &ApexPage{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated code files
	p.SourcePath = path

	return p, nil
}
