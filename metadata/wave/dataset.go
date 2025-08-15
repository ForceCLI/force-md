package wave

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const DATA_SET_NAME = "WaveDataset"

func init() {
	internal.TypeRegistry.Register(DATA_SET_NAME, func(path string) (metadata.RegisterableMetadata, error) { return OpenDataset(path) })
}

type WaveDataset struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"WaveDataset"`
	Xmlns       string   `xml:"xmlns,attr"`
	Application struct {
		Text string `xml:",chardata"`
	} `xml:"application"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	DatasetType *struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	TemplateAssetSourceName *struct {
		Text string `xml:",chardata"`
	} `xml:"templateAssetSourceName"`

	// Fields for handling companion files
	SourcePath    string `xml:"-"`
	SourceContent []byte `xml:"-"`
}

func (c *WaveDataset) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveDataset) Type() metadata.MetadataType {
	return DATA_SET_NAME
}

func (c *WaveDataset) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the dataset name from metadata
	datasetName := c.MetadataInfo.Name()
	if datasetName == "" {
		return nil, fmt.Errorf("dataset name is empty")
	}

	// Get the directory name for wave datasets
	dirName := registry.GetCanonicalDirectoryName(DATA_SET_NAME)

	// Load source content if we haven't already
	if c.SourceContent == nil {
		c.SourceContent = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".wds")
	}

	// Marshal the XML metadata using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dataset metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for WaveDataset
		return helpers.GenerateMetadataFilePaths(dirName, string(datasetName), ".wds-meta.xml", ".wds", xmlContent, c.SourceContent), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func OpenDataset(path string) (*WaveDataset, error) {
	p := &WaveDataset{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated content files
	p.SourcePath = path

	return p, nil
}
