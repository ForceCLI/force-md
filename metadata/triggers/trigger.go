package trigger

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "ApexTrigger"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ApexTrigger struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"ApexTrigger"`
	Xmlns      string   `xml:"xmlns,attr"`
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

	// Non-XML field to store the actual trigger code
	SourceCode []byte `xml:"-"`
	SourcePath string `xml:"-"`
}

func (c *ApexTrigger) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApexTrigger) Type() metadata.MetadataType {
	return NAME
}

func (c *ApexTrigger) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the trigger name from metadata
	triggerName := c.MetadataInfo.Name()
	if triggerName == "" {
		return nil, fmt.Errorf("trigger name is empty")
	}

	// Get the directory name for triggers
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Load source code if we haven't already
	if c.SourceCode == nil {
		c.SourceCode = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".trigger")
	}

	// Marshal the XML metadata using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal trigger metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for ApexTrigger
		return helpers.GenerateMetadataFilePaths(dirName, string(triggerName), ".trigger-meta.xml", ".trigger", xmlContent, c.SourceCode), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func Open(path string) (*ApexTrigger, error) {
	p := &ApexTrigger{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated code files
	p.SourcePath = path

	return p, nil
}
