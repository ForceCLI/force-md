package lwc

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "LightningComponentBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type LightningComponentBundle struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"LightningComponentBundle"`
	Xmlns      string   `xml:"xmlns,attr"`
	Fqn        *string  `xml:"fqn,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	IsExposed struct {
		Text string `xml:",chardata"`
	} `xml:"isExposed"`
	MasterLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	RuntimeNamespace *struct {
		Text string `xml:",chardata"`
	} `xml:"runtimeNamespace"`
	Targets *struct {
		Target []struct {
			Text string `xml:",chardata"`
		} `xml:"target"`
	} `xml:"targets"`
	TargetConfigs *struct {
		TargetConfig []struct {
			Xmlns    *string `xml:"xmlns,attr"`
			Targets  string  `xml:"targets,attr"`
			Property []struct {
				Name        string  `xml:"name,attr"`
				Label       *string `xml:"label,attr"`
				Type        string  `xml:"type,attr"`
				Role        *string `xml:"role,attr"`
				Description *string `xml:"description,attr"`
				Datasource  *string `xml:"datasource,attr"`
				Default     *string `xml:"default,attr"`
				Required    *string `xml:"required,attr"`
			} `xml:"property"`
			SupportedFormFactors *struct {
				SupportedFormFactor []struct {
					Type string `xml:"type,attr"`
				} `xml:"supportedFormFactor"`
			} `xml:"supportedFormFactors"`
			ActionType *struct {
				Text string `xml:",chardata"`
			} `xml:"actionType"`
		} `xml:"targetConfig"`
	} `xml:"targetConfigs"`
	BundleFiles map[string][]byte `xml:"-"`
	SourcePath  string            `xml:"-"`
}

func (c *LightningComponentBundle) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *LightningComponentBundle) Type() metadata.MetadataType {
	return NAME
}

// Tidy implements the general.Tidyable interface
// It sorts targets and targetConfigs to ensure consistent output
func (c *LightningComponentBundle) Tidy() {
	// Sort targets if they exist
	if c.Targets != nil && len(c.Targets.Target) > 0 {
		sort.Slice(c.Targets.Target, func(i, j int) bool {
			return c.Targets.Target[i].Text < c.Targets.Target[j].Text
		})
	}

	// Sort targetConfigs if they exist
	if c.TargetConfigs != nil && len(c.TargetConfigs.TargetConfig) > 0 {
		sort.Slice(c.TargetConfigs.TargetConfig, func(i, j int) bool {
			return c.TargetConfigs.TargetConfig[i].Targets < c.TargetConfigs.TargetConfig[j].Targets
		})

		// Sort properties within each targetConfig
		for idx := range c.TargetConfigs.TargetConfig {
			tc := &c.TargetConfigs.TargetConfig[idx]
			if len(tc.Property) > 0 {
				sort.Slice(tc.Property, func(i, j int) bool {
					return tc.Property[i].Name < tc.Property[j].Name
				})
			}

			// Sort supportedFormFactors if they exist
			if tc.SupportedFormFactors != nil && len(tc.SupportedFormFactors.SupportedFormFactor) > 0 {
				sort.Slice(tc.SupportedFormFactors.SupportedFormFactor, func(i, j int) bool {
					return tc.SupportedFormFactors.SupportedFormFactor[i].Type < tc.SupportedFormFactors.SupportedFormFactor[j].Type
				})
			}
		}
	}
}

func (c *LightningComponentBundle) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the component name from metadata
	componentName := c.MetadataInfo.Name()
	if componentName == "" {
		return nil, fmt.Errorf("component name is empty")
	}

	// Get the directory name for LWC components
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Load bundle files if we haven't already
	if err := c.loadBundleFiles(); err != nil {
		return nil, fmt.Errorf("failed to load bundle files: %w", err)
	}

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal LWC metadata: %w", err)
	}

	files := make(map[string][]byte)

	// Validate format
	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same bundle structure for LWC
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	// LWC bundles should be in a directory with all component files
	// Both source and metadata formats use the same bundle structure
	componentDir := filepath.Join(dirName, string(componentName))

	// Add the metadata file
	metadataFileName := string(componentName) + ".js-meta.xml"
	files[filepath.Join(componentDir, metadataFileName)] = xmlContent

	// Add all bundle files (excluding __tests__ which are already filtered in Open)
	for fileName, content := range c.BundleFiles {
		// Skip any __tests__ files that might have been loaded
		// (this shouldn't happen if Open() is working correctly, but being defensive)
		if strings.Contains(fileName, "__tests__") {
			continue
		}
		files[filepath.Join(componentDir, fileName)] = content
	}

	// If no other files were found (could happen during round-trip),
	// this is okay - we at least have the metadata file

	return files, nil
}

func Open(path string) (*LightningComponentBundle, error) {
	p := &LightningComponentBundle{
		BundleFiles: make(map[string][]byte),
	}

	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find bundle files
	p.SourcePath = path

	// Note: We intentionally don't load bundle files here.
	// The Files() method will handle loading them when needed.
	// This keeps Open() focused on just parsing the metadata XML.

	return p, nil
}

// loadBundleFiles loads all the bundle files from the source directory
func (c *LightningComponentBundle) loadBundleFiles() error {
	if c.SourcePath == "" || len(c.BundleFiles) > 0 {
		// No source path or files already loaded
		return nil
	}

	files, err := helpers.LoadBundleFiles(c.SourcePath, "__tests__", "-meta.xml")
	if err != nil {
		return err
	}
	if files != nil {
		c.BundleFiles = files
	}
	return nil
}
