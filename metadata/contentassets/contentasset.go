package contentasset

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/registry"
)

const NAME = "ContentAsset"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ContentAsset struct {
	metadata.MetadataInfo
	XMLName                  xml.Name `xml:"ContentAsset"`
	Xmlns                    string   `xml:"xmlns,attr"`
	IsVisibleByExternalUsers struct {
		Text string `xml:",chardata"`
	} `xml:"isVisibleByExternalUsers"`
	Language struct {
		Text string `xml:",chardata"`
	} `xml:"language"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Relationships struct {
		Organization struct {
			Access struct {
				Text string `xml:",chardata"`
			} `xml:"access"`
		} `xml:"organization"`
	} `xml:"relationships"`
	Versions struct {
		Version struct {
			Number struct {
				Text string `xml:",chardata"`
			} `xml:"number"`
			PathOnClient struct {
				Text string `xml:",chardata"`
			} `xml:"pathOnClient"`
		} `xml:"version"`
	} `xml:"versions"`
}

func (c *ContentAsset) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ContentAsset) Type() metadata.MetadataType {
	return NAME
}

func (c *ContentAsset) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the asset name from metadata
	assetName := c.MetadataInfo.Name()
	if assetName == "" {
		return nil, fmt.Errorf("asset name is empty")
	}

	// Get the directory name for content assets
	dirName := registry.GetCanonicalDirectoryName(NAME)

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal content asset metadata: %w", err)
	}

	files := make(map[string][]byte)

	// Get the original path from metadata info to find the asset file
	originalPath := string(c.MetadataInfo.Path())
	metadataDir := filepath.Dir(originalPath)

	// Try to find and read the asset content file
	var assetContent []byte
	var assetPath string

	if strings.HasSuffix(originalPath, "-meta.xml") {
		// Source format: .asset-meta.xml -> .asset
		assetPath = filepath.Join(metadataDir, string(assetName)+".asset")
	} else if strings.HasSuffix(originalPath, ".xml") {
		// Metadata format: might be .asset.xml -> .asset
		assetPath = filepath.Join(metadataDir, string(assetName)+".asset")
	}

	// Try to read the asset content file
	if assetPath != "" {
		if content, err := ioutil.ReadFile(assetPath); err == nil {
			assetContent = content
		}
	}

	switch format {
	case metadata.SourceFormat:
		// Source format: .asset-meta.xml for metadata
		files[filepath.Join(dirName, string(assetName)+".asset-meta.xml")] = xmlContent

		// Add the asset content if we found it
		if assetContent != nil {
			files[filepath.Join(dirName, string(assetName)+".asset")] = assetContent
		}

	case metadata.MetadataFormat:
		// Metadata format: .asset-meta.xml for metadata (same as source format)
		files[filepath.Join(dirName, string(assetName)+".asset-meta.xml")] = xmlContent

		// Add the asset content if we found it
		if assetContent != nil {
			files[filepath.Join(dirName, string(assetName)+".asset")] = assetContent
		}

	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	return files, nil
}

func Open(path string) (*ContentAsset, error) {
	p := &ContentAsset{}
	return p, metadata.ParseMetadataXml(p, path)
}
