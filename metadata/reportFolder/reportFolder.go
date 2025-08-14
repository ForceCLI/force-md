package reportFolder

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "ReportFolder"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type FolderShare struct {
	AccessLevel  string `xml:"accessLevel"`
	SharedTo     string `xml:"sharedTo"`
	SharedToType string `xml:"sharedToType"`
}

type ReportFolder struct {
	metadata.MetadataInfo
	XMLName      xml.Name      `xml:"ReportFolder"`
	Xmlns        string        `xml:"xmlns,attr"`
	FolderShares []FolderShare `xml:"folderShares"`
	Name         string        `xml:"name"`
}

func (c *ReportFolder) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ReportFolder) Type() metadata.MetadataType {
	return NAME
}

func (c *ReportFolder) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Get the directory name for report folders
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Extract the folder structure from the original path
	// e.g., reports/CRM_Admin_Exception_Reports/SystemAdminExceptionDashboard/BD_Discharge_Reports_Master.reportFolder-meta.xml
	// Should preserve: CRM_Admin_Exception_Reports/SystemAdminExceptionDashboard/BD_Discharge_Reports_Master

	var relativePath string
	if strings.Contains(originalPath, "reports/") {
		// Extract everything after "reports/"
		parts := strings.Split(originalPath, "reports/")
		if len(parts) > 1 {
			relativePath = parts[1]
		}
	}

	if relativePath == "" {
		return nil, fmt.Errorf("could not extract report folder path from %s", originalPath)
	}

	// Remove the file extension and -meta.xml suffix to get the clean relative path
	relativePath = strings.TrimSuffix(relativePath, "-meta.xml")
	relativePath = strings.TrimSuffix(relativePath, ".reportFolder")

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal report folder metadata: %w", err)
	}

	files := make(map[string][]byte)

	var fileName string
	switch format {
	case metadata.SourceFormat:
		// Source format: preserve folder structure and add -meta.xml suffix
		fileName = relativePath + "-meta.xml"
	case metadata.MetadataFormat:
		// Metadata format: preserve folder structure, add -meta.xml suffix
		fileName = relativePath + "-meta.xml"
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	files[filepath.Join(dirName, fileName)] = xmlContent

	return files, nil
}

func Open(path string) (*ReportFolder, error) {
	p := &ReportFolder{}
	return p, metadata.ParseMetadataXml(p, path)
}
