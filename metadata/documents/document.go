package document

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "Document"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Document struct {
	metadata.MetadataInfo
	XMLName         xml.Name     `xml:"Document"`
	Xmlns           string       `xml:"xmlns,attr"`
	Description     *TextLiteral `xml:"description"`
	InternalUseOnly *struct {
		Text string `xml:",chardata"`
	} `xml:"internalUseOnly"`
	Keywords *struct {
		Text string `xml:",chardata"`
	} `xml:"keywords"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Public struct {
		Text string `xml:",chardata"`
	} `xml:"public"`
	DocumentContent []byte `xml:"-"`
	DocumentExt     string `xml:"-"`
}

func (c *Document) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Document) Type() metadata.MetadataType {
	return NAME
}

func (c *Document) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Get the directory name for documents
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Extract the folder structure from the original path
	var relativePath string
	if strings.Contains(originalPath, "documents/") {
		// Extract everything after "documents/"
		parts := strings.Split(originalPath, "documents/")
		if len(parts) > 1 {
			relativePath = parts[1]
		}
	}

	if relativePath == "" {
		return nil, fmt.Errorf("could not extract document path from %s", originalPath)
	}

	// Get the directory part (e.g., "Images/")
	documentDir := filepath.Dir(relativePath)

	// Determine the base name and file extension based on the current path
	fileName := filepath.Base(relativePath)
	var baseName, actualFileExt string

	if strings.HasSuffix(fileName, ".document-meta.xml") {
		// Coming from source format: AEG_Transparent_Logo.document-meta.xml
		baseName = strings.TrimSuffix(fileName, ".document-meta.xml")
	} else if strings.HasSuffix(fileName, "-meta.xml") {
		// Coming from metadata format: AEG_Transparent_Logo.gif-meta.xml
		nameWithoutMeta := strings.TrimSuffix(fileName, "-meta.xml")
		actualFileExt = filepath.Ext(nameWithoutMeta)
		baseName = strings.TrimSuffix(nameWithoutMeta, actualFileExt)
	} else {
		return nil, fmt.Errorf("unrecognized document metadata file format: %s", fileName)
	}

	// Use the loaded document content and extension
	actualFileContent := c.DocumentContent
	if c.DocumentExt != "" {
		actualFileExt = c.DocumentExt
	}

	if actualFileExt == "" {
		return nil, fmt.Errorf("could not determine file extension for document %s", baseName)
	}

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document metadata: %w", err)
	}

	files := make(map[string][]byte)

	var metadataFileName, documentFileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: preserve folder structure
		// Metadata: documents/Images/AEG_Transparent_Logo.document-meta.xml
		// File: documents/Images/AEG_Transparent_Logo.gif
		metadataFileName = baseName + ".document-meta.xml"
		documentFileName = baseName + actualFileExt

	case metadata.MetadataFormat:
		// Metadata format: preserve folder structure
		// Metadata: documents/Images/AEG_Transparent_Logo.gif-meta.xml
		// File: documents/Images/AEG_Transparent_Logo.gif
		metadataFileName = baseName + actualFileExt + "-meta.xml"
		documentFileName = baseName + actualFileExt

	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	// Add metadata file
	if documentDir != "." && documentDir != "" {
		metadataFileName = filepath.Join(documentDir, metadataFileName)
	}
	files[filepath.Join(dirName, metadataFileName)] = xmlContent

	// Add actual document file if we found it
	if actualFileContent != nil {
		if documentDir != "." && documentDir != "" {
			documentFileName = filepath.Join(documentDir, documentFileName)
		}
		files[filepath.Join(dirName, documentFileName)] = actualFileContent
	}

	return files, nil
}

func Open(path string) (*Document, error) {
	p := &Document{}

	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Load the document file if it exists
	if strings.HasSuffix(path, "-meta.xml") {
		// This is a source format file - look for the corresponding document file
		metadataDir := filepath.Dir(path)
		fileName := filepath.Base(path)

		var baseName, actualFileExt string
		if strings.HasSuffix(fileName, ".document-meta.xml") {
			baseName = strings.TrimSuffix(fileName, ".document-meta.xml")
		} else if strings.HasSuffix(fileName, "-meta.xml") {
			// Format: DocumentName.ext-meta.xml
			nameWithoutMeta := strings.TrimSuffix(fileName, "-meta.xml")
			actualFileExt = filepath.Ext(nameWithoutMeta)
			baseName = strings.TrimSuffix(nameWithoutMeta, actualFileExt)
		}

		// First try with known extension
		if actualFileExt != "" {
			documentFilePath := filepath.Join(metadataDir, baseName+actualFileExt)
			if content, err := os.ReadFile(documentFilePath); err == nil {
				p.DocumentContent = content
				p.DocumentExt = actualFileExt
			}
		} else {
			// Scan directory for the document file
			if entries, err := os.ReadDir(metadataDir); err == nil {
				for _, entry := range entries {
					if entry.IsDir() {
						continue
					}

					entryName := entry.Name()
					entryBase := strings.TrimSuffix(entryName, filepath.Ext(entryName))

					// Check if this is the document file (not the metadata file)
					if entryBase == baseName && !strings.HasSuffix(entryName, "-meta.xml") && !strings.HasSuffix(entryName, ".document-meta.xml") {
						actualFileExt = filepath.Ext(entryName)
						actualFilePath := filepath.Join(metadataDir, entryName)
						if content, err := os.ReadFile(actualFilePath); err == nil {
							p.DocumentContent = content
							p.DocumentExt = actualFileExt
						}
						break
					}
				}
			}
		}
	}
	// For metadata format, we typically don't have separate document files

	return p, nil
}
