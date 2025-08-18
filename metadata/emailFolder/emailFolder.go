package emailFolder

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "EmailFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type EmailFolder struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"EmailFolder"`
	Xmlns      string   `xml:"xmlns,attr"`
	AccessType struct {
		Text string `xml:",chardata"`
	} `xml:"accessType"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	PublicFolderAccess struct {
		Text string `xml:",chardata"`
	} `xml:"publicFolderAccess"`
	SharedTo []struct {
	} `xml:"sharedTo"`
}

func (c *EmailFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EmailFolder) Type() metadata.MetadataType {
	return NAME
}

func (c *EmailFolder) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Get the directory name for email folders
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Extract the relative path within the email directory to preserve folder structure
	var relativePath string
	if strings.Contains(originalPath, "email/") {
		// Extract everything after "email/"
		parts := strings.Split(originalPath, "email/")
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
		return nil, fmt.Errorf("failed to marshal email folder metadata: %w", err)
	}

	files := make(map[string][]byte)

	var fileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: FolderName.emailFolder-meta.xml
		fileName = string(folderName) + ".emailFolder-meta.xml"
	case metadata.MetadataFormat:
		// Metadata format: FolderName-meta.xml (no .emailFolder part)
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

func Open(path string) (*EmailFolder, error) {
	p := &EmailFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
