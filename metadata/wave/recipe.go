package wave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
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
}

func (c *WaveRecipe) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveRecipe) Type() metadata.MetadataType {
	return RECIPE_NAME
}

func OpenRecipe(path string) (*WaveRecipe, error) {
	p := &WaveRecipe{}
	return p, metadata.ParseMetadataXml(p, path)
}
