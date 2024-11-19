package internal

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var NotXMLError = errors.New("Could not parse as XML")
var MetadataFileNotFound = errors.New("Could not identify metadata type")

type RelativePath = string
type AbsolutePath = string
type ForceMetadataFiles map[RelativePath][]byte

// Map relative paths to filesystem paths
type ForceMetadataFilePaths map[RelativePath]AbsolutePath

type MetadataType = string

// If the file in path contains metadata, return it.  Otherwise, try to find
// the corresponding file that contains metadata.

func metadataFileFromPath(path string) (string, error) {
	if IsMetadataFile(path) {
		return path, nil
	}
	if IsMetadataFile(path + "-meta.xml") {
		return path + "-meta.xml", nil
	}
	if m := metadataFileInSameFolder(path); m != "" && IsMetadataFile(m) {
		return m, nil
	}

	// Documents
	documentMetadata := strings.TrimSuffix(path, filepath.Ext(path)) + ".document-meta.xml"
	if IsMetadataFile(documentMetadata) {
		return documentMetadata, nil
	}

	// For static resources, walk up the filesystem to find the metadata file
	currentPath := path
	for {
		if currentPath == "" || currentPath == "." || currentPath == "/" {
			break
		}

		dirName := filepath.Base(currentPath)
		parentDir := filepath.Dir(currentPath)
		parentDirName := filepath.Base(parentDir)

		// Check if parent directory is 'staticresources'
		if parentDirName == "staticresources" {
			// Determine if currentPath is a file or directory
			var resourceName string
			fileInfo, err := os.Stat(currentPath)
			if err == nil && fileInfo.IsDir() {
				// It's a directory under 'staticresources'
				resourceName = dirName
			} else {
				// It's a file under 'staticresources', get file name without extension
				resourceName = strings.TrimSuffix(dirName, filepath.Ext(dirName))
			}
			resourceMetadata := filepath.Join(parentDir, resourceName+".resource-meta.xml")
			if IsMetadataFile(resourceMetadata) {
				return resourceMetadata, nil
			}
			// If not found, break as we have reached 'staticresources'
			break
		}

		currentPath = parentDir
	}

	return "", MetadataFileNotFound
}

// Look for a metadata file in the same folder as path to support bundled
// metadata (Aura and LWC)
func metadataFileInSameFolder(path string) string {
	info, err := os.Stat(path)
	if err != nil {
		return ""
	}
	if info.IsDir() {
		return ""
	}
	dirName := filepath.Dir(path)
	pattern := dirName + string(os.PathSeparator) + "*-meta.xml"
	files, err := filepath.Glob(pattern)
	if err != nil {
		return ""
	}
	if len(files) == 1 {
		return files[0]
	}
	return ""
}

func MetadataFromPath(path string) (RegisterableMetadata, error) {
	path, err := metadataFileFromPath(path)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(path)
	if err != nil {
		return nil, err
	}
	element, err := getRootElementName(path)
	if err != nil {
		return nil, err
	}
	if open, ok := TypeRegistry[element]; ok {
		return open(path)
	}
	return nil, fmt.Errorf("Could not find metadata")
}

func IsMetadataFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	element, err := getRootElementName(path)
	if err != nil {
		return false
	}
	if _, ok := TypeRegistry[element]; ok {
		return ok
	}
	return false
}

func RootElementName(xmlData []byte) (string, error) {
	decoder := xml.NewDecoder(io.NopCloser(bytes.NewReader(xmlData)))

	foundXML := false
	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("Error while parsing XML: %w", err)
		}

		switch element := t.(type) {
		case xml.ProcInst:
			// Check for the XML declaration and return the version if found
			if element.Target == "xml" {
				foundXML = true
			}
		case xml.StartElement:
			if !foundXML {
				return "", fmt.Errorf("%w: No XML declaration found", NotXMLError)
			}
			// Return the name of the root element
			return element.Name.Local, nil
		}
	}
	return "", fmt.Errorf("%w: No XML elements found", NotXMLError)
}

func getRootElementName(file string) (string, error) {
	xmlData, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Could not read XML file: %w", err)
	}

	return RootElementName(xmlData)
}
