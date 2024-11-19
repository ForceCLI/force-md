package wave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const RECIPE_NAME = "WaveRecipe"

func init() {
	internal.TypeRegistry.Register(RECIPE_NAME, func(path string) (internal.RegisterableMetadata, error) { return OpenRecipe(path) })
}

type WaveRecipe struct {
	internal.MetadataInfo
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

func (c *WaveRecipe) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveRecipe) Type() internal.MetadataType {
	return RECIPE_NAME
}

func OpenRecipe(path string) (*WaveRecipe, error) {
	p := &WaveRecipe{}
	return p, internal.ParseMetadataXml(p, path)
}
