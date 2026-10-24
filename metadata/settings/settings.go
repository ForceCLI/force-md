package settings

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"golang.org/x/net/html/charset"
)

func init() {
	internal.TypeRegistry.RegisterSuffix("Settings", func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Settings struct {
	metadata.MetadataInfo
	XMLName xml.Name
	XMLNS   string `xml:"xmlns,attr"`
	Xsi     string `xml:"xmlns:xsi,attr,omitempty"`
	Content []byte `xml:",innerxml"`
}

func (c *Settings) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Settings) Type() metadata.MetadataType {
	return c.XMLName.Local
}

func (c *Settings) Files(metadataFormat metadata.Format) (map[string][]byte, error) {
	files := make(map[string][]byte)

	// Use the Type() method which returns XMLName.Local
	settingsType := c.Type()
	if settingsType == "" {
		// Fallback: try to extract from metadata name
		if c.MetadataInfo.Name() != "" {
			settingsType = string(c.MetadataInfo.Name())
			// Ensure it has Settings suffix
			if !strings.HasSuffix(settingsType, "Settings") {
				settingsType = settingsType + "Settings"
			}
		} else {
			return nil, fmt.Errorf("unable to determine settings type")
		}
	}

	// Extract the base name without "Settings" suffix for the filename
	baseName := settingsType
	if strings.HasSuffix(settingsType, "Settings") {
		baseName = strings.TrimSuffix(settingsType, "Settings")
	}

	// Build the XML manually to preserve the inner content structure.
	// Always use the full type name (with Settings suffix) in the XML.
	attrs := fmt.Sprintf(` xmlns="%s"`, c.XMLNS)
	if c.Xsi != "" {
		attrs += fmt.Sprintf(` xmlns:xsi="%s"`, c.Xsi)
	}
	var xmlContent []byte
	if len(bytes.TrimSpace(c.Content)) == 0 {
		// Salesforce writes settings without elements as a self-closed root
		// followed by a blank line
		xmlContent = []byte(fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<%s%s/>\n\n", settingsType, attrs))
	} else {
		xmlContent = []byte(fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<%s%s>%s</%s>\n", settingsType, attrs, string(c.Content), settingsType))
	}
	if metadataFormat == metadata.SourceFormat {
		files[fmt.Sprintf("settings/%s.settings-meta.xml", baseName)] = xmlContent
	} else {
		files[fmt.Sprintf("settings/%s.settings", baseName)] = xmlContent
	}

	return files, nil
}

func getRootElementName(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	r := bytes.NewReader(contents)
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel

	for {
		t, err := dec.Token()
		if err != nil {
			return "", err
		}
		if se, ok := t.(xml.StartElement); ok {
			return se.Name.Local, nil
		}
	}
}

func Open(path string) (*Settings, error) {
	p := &Settings{}
	err := metadata.ParseMetadataXml(p, path)
	// Ensure XMLName is preserved with proper namespace
	if p.XMLName.Local == "" {
		// Try to get the root element name from the file to determine the settings type
		elementName, _ := getRootElementName(path)
		if elementName != "" {
			p.XMLName.Local = elementName
		}
	}
	if p.XMLNS == "" {
		p.XMLNS = "http://soap.sforce.com/2006/04/metadata"
	}
	return p, err
}
