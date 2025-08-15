package wave

import (
	"encoding/xml"
	"fmt"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/helpers"
	"github.com/ForceCLI/force-md/registry"
)

const RECIPE_NAME = "WaveRecipe"

func init() {
	internal.TypeRegistry.Register(RECIPE_NAME, func(path string) (metadata.RegisterableMetadata, error) { return OpenRecipe(path) })
}

type WaveRecipe struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"WaveRecipe"`
	Xmlns   string   `xml:"xmlns,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Content struct {
		Nil string `xml:"nil,attr"`
	} `xml:"content"`
	Dataflow struct {
		Text string `xml:",chardata"`
	} `xml:"dataflow"`
	Format struct {
		Text string `xml:",chardata"`
	} `xml:"format"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`

	// Fields for handling companion files
	SourcePath    string `xml:"-"`
	SourceContent []byte `xml:"-"`
}

func (c *WaveRecipe) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveRecipe) Type() metadata.MetadataType {
	return RECIPE_NAME
}

func (c *WaveRecipe) Files(format metadata.Format) (map[string][]byte, error) {
	// Get the recipe name from metadata
	recipeName := c.MetadataInfo.Name()
	if recipeName == "" {
		return nil, fmt.Errorf("recipe name is empty")
	}

	// Get the directory name for wave recipes
	dirName := registry.GetCanonicalDirectoryName(RECIPE_NAME)

	// Load source content if we haven't already
	if c.SourceContent == nil {
		c.SourceContent = helpers.LoadCompanionFile(c.SourcePath, "-meta.xml", ".wdpr")
	}

	// Marshal the XML metadata using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal recipe metadata: %w", err)
	}

	switch format {
	case metadata.SourceFormat, metadata.MetadataFormat:
		// Both formats use the same file structure for WaveRecipe
		return helpers.GenerateMetadataFilePaths(dirName, string(recipeName), ".wdpr-meta.xml", ".wdpr", xmlContent, c.SourceContent), nil
	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}
}

func OpenRecipe(path string) (*WaveRecipe, error) {
	p := &WaveRecipe{}
	if err := metadata.ParseMetadataXml(p, path); err != nil {
		return nil, err
	}

	// Store the source path - Files() will use this to find associated content files
	p.SourcePath = path

	return p, nil
}
