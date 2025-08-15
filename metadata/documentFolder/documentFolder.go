package documentFolder

import (
	"encoding/xml"
	"fmt"
	"path/filepath"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "DocumentFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type DocumentFolder struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"DocumentFolder"`
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
}

func (c *DocumentFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *DocumentFolder) Type() metadata.MetadataType {
	return NAME
}

func (c *DocumentFolder) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the folder name from metadata
	folderName := c.MetadataInfo.Name()
	if folderName == "" {
		return nil, fmt.Errorf("folder name is empty")
	}

	// Get the directory name for document folders
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document folder metadata: %w", err)
	}

	files := make(map[string][]byte)

	var fileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: Document_Templates.documentFolder-meta.xml
		fileName = string(folderName) + ".documentFolder-meta.xml"
	case metadata.MetadataFormat:
		// Metadata format: Document_Templates-meta.xml (no .documentFolder part)
		fileName = string(folderName) + "-meta.xml"
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	files[filepath.Join(dirName, fileName)] = xmlContent

	return files, nil
}

func Open(path string) (*DocumentFolder, error) {
	p := &DocumentFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
