package uibundles

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
)

const NAME = "UIBundle"
const directoryName = "uiBundles"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type UIBundle struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"UIBundle"`
	Xmlns       string   `xml:"xmlns,attr"`
	MasterLabel *struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	IsActive *struct {
		Text string `xml:",chardata"`
	} `xml:"isActive"`
	Version *struct {
		Text string `xml:",chardata"`
	} `xml:"version"`
	BundleFiles map[string][]byte `xml:"-"`
	SourcePath  string            `xml:"-"`
}

func (c *UIBundle) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *UIBundle) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*UIBundle, error) {
	p := &UIBundle{BundleFiles: make(map[string][]byte)}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}
	p.SourcePath = path
	return p, nil
}

func (c *UIBundle) Files(format metadata.Format) (map[string][]byte, error) {
	bundleName := c.MetadataInfo.Name()
	if bundleName == "" {
		return nil, fmt.Errorf("ui bundle name is empty")
	}
	if err := c.loadBundleFiles(); err != nil {
		return nil, fmt.Errorf("failed to load ui bundle files: %w", err)
	}
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ui bundle metadata: %w", err)
	}
	files := make(map[string][]byte)
	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
	bundleDir := filepath.Join(directoryName, string(bundleName))
	metadataFileName := string(bundleName) + ".uibundle-meta.xml"
	files[filepath.Join(bundleDir, metadataFileName)] = xmlContent
	for fileName, content := range c.BundleFiles {
		files[filepath.Join(bundleDir, fileName)] = content
	}
	return files, nil
}

func (c *UIBundle) loadBundleFiles() error {
	if c.SourcePath == "" || len(c.BundleFiles) > 0 {
		return nil
	}
	files, err := helpers.LoadBundleFiles(c.SourcePath, ".uibundle-meta.xml")
	if err != nil {
		return err
	}
	if files != nil {
		for fileName, content := range files {
			if strings.HasSuffix(fileName, ".uibundle-meta.xml") {
				continue
			}
			c.BundleFiles[fileName] = content
		}
	}
	return nil
}
