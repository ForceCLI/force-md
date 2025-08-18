package dashboardFolder

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "DashboardFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DashboardFolder struct {
	metadata.MetadataInfo
	XMLName      xml.Name `xml:"DashboardFolder"`
	Xmlns        string   `xml:"xmlns,attr"`
	FolderShares []struct {
		AccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"accessLevel"`
		SharedTo struct {
			Text string `xml:",chardata"`
		} `xml:"sharedTo"`
		SharedToType struct {
			Text string `xml:",chardata"`
		} `xml:"sharedToType"`
	} `xml:"folderShares"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

func (c *DashboardFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DashboardFolder) Type() metadata.MetadataType {
	return NAME
}

func (c *DashboardFolder) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Get the directory name for dashboard folders
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Extract the relative path within the dashboards directory to preserve folder structure
	var relativePath string
	if strings.Contains(originalPath, "dashboards/") {
		// Extract everything after "dashboards/"
		parts := strings.Split(originalPath, "dashboards/")
		if len(parts) > 1 {
			relativePath = parts[1]
			// Remove the file name to get just the directory path
			relativePath = filepath.Dir(relativePath)
			// If it's just ".", we're in the root
			if relativePath == "." {
				relativePath = ""
			}
		}
	}

	// Get the folder name from metadata
	folderName := c.MetadataInfo.Name()
	if folderName == "" {
		return nil, fmt.Errorf("folder name is empty")
	}

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dashboard folder metadata: %w", err)
	}

	files := make(map[string][]byte)

	var fileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: FolderName.dashboardFolder-meta.xml
		fileName = string(folderName) + ".dashboardFolder-meta.xml"
	case metadata.MetadataFormat:
		// Metadata format: FolderName-meta.xml (no .dashboardFolder part)
		fileName = string(folderName) + "-meta.xml"
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	// If there's a subfolder structure, preserve it
	if relativePath != "" {
		fileName = filepath.Join(relativePath, fileName)
	}

	files[filepath.Join(dirName, fileName)] = xmlContent

	return files, nil
}

func Open(path string) (*DashboardFolder, error) {
	p := &DashboardFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
