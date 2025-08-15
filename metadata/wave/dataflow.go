package wave

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const DATA_FLOW_NAME = "WaveDataflow"

func init() {
	internal.TypeRegistry.Register(DATA_FLOW_NAME, func(path string) (metadata.RegisterableMetadata, error) { return OpenDataflow(path) })
}

type WaveDataflow struct {
	metadata.MetadataInfo
	XMLName     xml.Name `xml:"WaveDataflow"`
	Xmlns       string   `xml:"xmlns,attr"`
	Xsi         string   `xml:"xsi,attr"`
	Application *struct {
		Text string `xml:",chardata"`
	} `xml:"application"`
	Content struct {
		Nil string `xml:"nil,attr"`
	} `xml:"content"`
	DataflowType struct {
		Text string `xml:",chardata"`
	} `xml:"dataflowType"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`

	// Non-XML fields to store the actual dataflow content
	SourceContent []byte `xml:"-"`
	SourcePath    string `xml:"-"`
}

func (c *WaveDataflow) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveDataflow) Type() metadata.MetadataType {
	return DATA_FLOW_NAME
}

func (c *WaveDataflow) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the dataflow name from metadata
	dataflowName := c.MetadataInfo.Name()
	if dataflowName == "" {
		return nil, fmt.Errorf("dataflow name is empty")
	}

	// Get the directory name for wave dataflows
	dirName := registry.GetCanonicalDirectoryName(DATA_FLOW_NAME)

	// Load source content if we haven't already
	if c.SourceContent == nil {
		c.SourceContent = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".wdf")
	}

	// Marshal the XML metadata using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dataflow metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for WaveDataflow
		return helpers.GenerateMetadataFilePaths(dirName, string(dataflowName), ".wdf-meta.xml", ".wdf", xmlContent, c.SourceContent), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func OpenDataflow(path string) (*WaveDataflow, error) {
	p := &WaveDataflow{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated content files
	p.SourcePath = path

	return p, nil
}
