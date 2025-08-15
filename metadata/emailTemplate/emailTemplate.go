package emailTemplate

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "EmailTemplate"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type EmailTemplate struct {
	metadata.MetadataInfo
	XMLName    xml.Name `xml:"EmailTemplate"`
	Xmlns      string   `xml:"xmlns,attr"`
	ApiVersion *struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion"`
	Available struct {
		Text string `xml:",chardata"`
	} `xml:"available"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	EncodingKey struct {
		Text string `xml:",chardata"`
	} `xml:"encodingKey"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Style struct {
		Text string `xml:",chardata"`
	} `xml:"style"`
	Subject struct {
		Text string `xml:",chardata"`
	} `xml:"subject"`
	TextOnly     *TextLiteral `xml:"textOnly"`
	TemplateType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	UiType *struct {
		Text string `xml:",chardata"`
	} `xml:"uiType"`
	EmailContent []byte `xml:"-"`
	SourcePath   string `xml:"-"`
}

func (c *EmailTemplate) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *EmailTemplate) Type() metadata.MetadataType {
	return NAME
}

func (c *EmailTemplate) Files(format metadata.Format) (map[string][]byte, error) {
	// Load email content if we haven't already
	if c.EmailContent == nil {
		c.EmailContent = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".email")
	}

	// Get the original path from metadata info
	originalPath := string(c.MetadataInfo.Path())

	// Get the directory name for email templates
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Extract the folder structure from the original path
	var relativePath string
	if strings.Contains(originalPath, "/email/") {
		// Extract everything after "/email/"
		parts := strings.Split(originalPath, "/email/")
		if len(parts) > 1 {
			relativePath = parts[1]
		}
	} else if strings.HasPrefix(filepath.Base(filepath.Dir(originalPath)), "email") {
		// Handle case where path is like "test-email/email/Health_Cloud/file.email-meta.xml"
		// Extract relative path from the email directory
		pathParts := strings.Split(originalPath, string(filepath.Separator))
		emailIndex := -1
		for i, part := range pathParts {
			if part == "email" {
				emailIndex = i
				break
			}
		}
		if emailIndex >= 0 && emailIndex < len(pathParts)-1 {
			relativePath = strings.Join(pathParts[emailIndex+1:], string(filepath.Separator))
		}
	}

	if relativePath == "" {
		// Fallback: just use the filename if we can't determine the path structure
		relativePath = filepath.Base(originalPath)
	}

	// Get the directory part (e.g., "Health_Cloud/")
	emailDir := filepath.Dir(relativePath)

	// Determine the base name based on the current path
	fileName := filepath.Base(relativePath)
	var baseName string

	if strings.HasSuffix(fileName, ".email-meta.xml") {
		// Coming from source format: Cancel_Admission_Notification.email-meta.xml
		baseName = strings.TrimSuffix(fileName, ".email-meta.xml")
	} else if strings.HasSuffix(fileName, ".email") {
		// Coming from metadata format: Cancel_Admission_Notification.email
		baseName = strings.TrimSuffix(fileName, ".email")
	} else {
		return nil, fmt.Errorf("unrecognized email template file format: %s", fileName)
	}

	// Use the loaded email content
	emailContent := c.EmailContent

	files := make(map[string][]byte)

	switch format {
	case metadata.SourceFormat:
		// Source format: preserve folder structure
		// Metadata: email/Health_Cloud/Cancel_Admission_Notification.email-meta.xml
		// Template: email/Health_Cloud/Cancel_Admission_Notification.email

		// Marshal the metadata to XML using internal.Marshal to get proper formatting
		xmlContent, err := internal.Marshal(c)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal email template metadata: %w", err)
		}

		metadataFileName := baseName + ".email-meta.xml"
		emailFileName := baseName + ".email"

		if emailDir != "." && emailDir != "" {
			metadataFileName = filepath.Join(emailDir, metadataFileName)
			emailFileName = filepath.Join(emailDir, emailFileName)
		}

		files[filepath.Join(dirName, metadataFileName)] = xmlContent

		// Add email template file if we found it
		if emailContent != nil {
			files[filepath.Join(dirName, emailFileName)] = emailContent
		}

	case metadata.MetadataFormat:
		// Metadata format: preserve folder structure
		// Create both files:
		// - .email file with Visualforce content (HTML part)
		// - .email-meta.xml file with metadata and text content

		// Marshal the metadata to XML using internal.Marshal to get proper formatting
		xmlContent, err := internal.Marshal(c)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal email template metadata: %w", err)
		}

		metadataFileName := baseName + ".email-meta.xml"
		emailFileName := baseName + ".email"

		if emailDir != "." && emailDir != "" {
			metadataFileName = filepath.Join(emailDir, metadataFileName)
			emailFileName = filepath.Join(emailDir, emailFileName)
		}

		// Always create the metadata file
		files[filepath.Join(dirName, metadataFileName)] = xmlContent

		// Add email template file if we found Visualforce content
		if emailContent != nil {
			files[filepath.Join(dirName, emailFileName)] = emailContent
		}

	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	return files, nil
}

func Open(path string) (*EmailTemplate, error) {
	p := &EmailTemplate{}

	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find the email content file
	p.SourcePath = path

	// Note: We intentionally don't load the email content here.
	// The Files() method will handle loading it when needed.
	// This keeps Open() focused on just parsing the metadata XML.

	return p, nil
}
