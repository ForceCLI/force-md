package staticresource

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "StaticResource"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type StaticResource struct {
	metadata.MetadataInfo
	XMLName      xml.Name `xml:"StaticResource"`
	Xmlns        *string  `xml:"xmlns,attr"`
	CacheControl struct {
		Text string `xml:",chardata"`
	} `xml:"cacheControl"`
	ContentType struct {
		Text string `xml:",chardata"`
	} `xml:"contentType"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *StaticResource) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *StaticResource) Type() metadata.MetadataType {
	return NAME
}

// zipDirectory creates a zip archive from a directory
func zipDirectory(dirPath string) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	defer zipWriter.Close()

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path from the base directory
		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if relPath == "." {
			return nil
		}

		// Create header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(relPath)

		// Write header
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If it's a directory, we're done
		if info.IsDir() {
			return nil
		}

		// If it's a file, write its contents
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		return nil, err
	}

	// Close the zip writer to flush any remaining data
	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *StaticResource) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the resource name from metadata
	resourceName := c.MetadataInfo.Name()
	if resourceName == "" {
		return nil, fmt.Errorf("resource name is empty")
	}

	// Get the directory name for static resources
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal static resource metadata: %w", err)
	}

	files := make(map[string][]byte)

	// Get the original path from metadata info to find the resource file
	originalPath := string(c.MetadataInfo.Path())
	metadataDir := filepath.Dir(originalPath)

	// Try to find and read the resource content
	var resourceContent []byte

	// First check if there's a directory with the resource name (unzipped static resource in SFDX)
	resourceDirPath := filepath.Join(metadataDir, string(resourceName))
	if info, err := os.Stat(resourceDirPath); err == nil && info.IsDir() {
		// It's a directory - we need to zip it for metadata format
		if format == metadata.MetadataFormat {
			zipContent, err := zipDirectory(resourceDirPath)
			if err != nil {
				return nil, fmt.Errorf("failed to zip directory %s: %w", resourceDirPath, err)
			}
			resourceContent = zipContent
		}
		// For source format, we would need to handle directory copying differently
		// but that's not typically how static resources work in source format
	} else {
		// Try to find .resource file
		resourcePath := filepath.Join(metadataDir, string(resourceName)+".resource")
		if content, err := ioutil.ReadFile(resourcePath); err == nil {
			resourceContent = content
		} else {
			// If .resource doesn't exist, look for the actual file with its real extension
			// Static resources in SFDX format can have their actual extensions (.css, .js, .zip, etc.)
			dirFiles, dirErr := ioutil.ReadDir(metadataDir)
			if dirErr == nil {
				for _, file := range dirFiles {
					if file.IsDir() {
						continue
					}
					fileName := file.Name()
					// Look for files that start with the resource name but aren't metadata files
					if strings.HasPrefix(fileName, string(resourceName)) &&
						!strings.HasSuffix(fileName, "-meta.xml") &&
						!strings.HasSuffix(fileName, ".resource-meta.xml") {
						// Found a potential resource file
						fullPath := filepath.Join(metadataDir, fileName)
						if content, err := ioutil.ReadFile(fullPath); err == nil {
							resourceContent = content
							break
						}
					}
				}
			}
		}
	}

	switch format {
	case metadata.SourceFormat:
		// Source format: .resource-meta.xml for metadata
		files[filepath.Join(dirName, string(resourceName)+".resource-meta.xml")] = xmlContent

		// Add the resource content if we found it
		if resourceContent != nil {
			files[filepath.Join(dirName, string(resourceName)+".resource")] = resourceContent
		}

	case metadata.MetadataFormat:
		// Metadata format: .resource-meta.xml for metadata (same as source format)
		files[filepath.Join(dirName, string(resourceName)+".resource-meta.xml")] = xmlContent

		// Add the resource content if we found it
		if resourceContent != nil {
			files[filepath.Join(dirName, string(resourceName)+".resource")] = resourceContent
		}

	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	return files, nil
}

func Open(path string) (*StaticResource, error) {
	p := &StaticResource{}
	return p, metadata.ParseMetadataXml(p, path)
}
