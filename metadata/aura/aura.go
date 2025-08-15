package aura

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "AuraDefinitionBundle"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type AuraDefinitionBundle struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"AuraDefinitionBundle"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MasterLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`

	// Non-XML fields to store bundle files
	BundleFiles map[string][]byte `xml:"-"` // filename -> content
	SourcePath  string            `xml:"-"`
}

func (c *AuraDefinitionBundle) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *AuraDefinitionBundle) Type() metadata.MetadataType {
	return NAME
}

func (c *AuraDefinitionBundle) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the bundle name from metadata
	bundleName := c.MetadataInfo.Name()
	if bundleName == "" {
		return nil, fmt.Errorf("bundle name is empty")
	}

	// Load bundle files if we haven't already
	if err := c.loadBundleFiles(); err != nil {
		return nil, fmt.Errorf("failed to load bundle files: %w", err)
	}

	// Get the directory name for aura
	dirName := registry.GetCanonicalDirectoryName(NAME)

	files := make(map[string][]byte)

	switch format {
	case metadata.SourceFormat:
		// Source format: bundle as a directory with multiple files
		// First, add the -meta.xml file using internal.Marshal to get proper formatting
		xmlContent, err := internal.Marshal(c)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal aura bundle metadata: %w", err)
		}

		// Determine the main file extension (.cmp for components, .app for apps, .evt for events)
		mainExt := ".cmp" // default to component
		for fileName := range c.BundleFiles {
			if strings.HasSuffix(fileName, ".app") {
				mainExt = ".app"
				break
			} else if strings.HasSuffix(fileName, ".evt") {
				mainExt = ".evt"
				break
			}
		}

		// Add the -meta.xml file
		metaPath := filepath.Join(dirName, string(bundleName), string(bundleName)+mainExt+"-meta.xml")
		files[metaPath] = xmlContent

		// Add all bundle files
		for fileName, content := range c.BundleFiles {
			filePath := filepath.Join(dirName, string(bundleName), fileName)
			files[filePath] = content
		}

	case metadata.MetadataFormat:
		// Metadata format: bundle as a directory with multiple files (same as source format but without -meta.xml suffix)
		xmlContent, err := internal.Marshal(c)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal aura bundle metadata: %w", err)
		}

		// Determine the main file extension (.cmp for components, .app for apps, .evt for events)
		mainExt := ".cmp" // default to component
		for fileName := range c.BundleFiles {
			if strings.HasSuffix(fileName, ".app") {
				mainExt = ".app"
				break
			} else if strings.HasSuffix(fileName, ".evt") {
				mainExt = ".evt"
				break
			}
		}

		// Add the -meta.xml file (in metadata format, it's still component.cmp-meta.xml)
		metaPath := filepath.Join(dirName, string(bundleName), string(bundleName)+mainExt+"-meta.xml")
		files[metaPath] = xmlContent

		// Add all bundle files
		for fileName, content := range c.BundleFiles {
			filePath := filepath.Join(dirName, string(bundleName), fileName)
			files[filePath] = content
		}

	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	return files, nil
}

func Open(path string) (*AuraDefinitionBundle, error) {
	p := &AuraDefinitionBundle{
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
func (c *AuraDefinitionBundle) loadBundleFiles() error {
	if c.SourcePath == "" || len(c.BundleFiles) > 0 {
		// No source path or files already loaded
		return nil
	}

	files, err := helpers.LoadBundleFiles(c.SourcePath, "-meta.xml")
	if err != nil {
		return err
	}

	// Aura stores files with just the filename, not the relative path
	if files != nil {
		c.BundleFiles = make(map[string][]byte)
		for path, content := range files {
			fileName := filepath.Base(path)
			c.BundleFiles[fileName] = content
		}
	}
	return nil
}
