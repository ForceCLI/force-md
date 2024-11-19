package wave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const DATA_SET_NAME = "WaveDataset"

func init() {
	internal.TypeRegistry.Register(DATA_SET_NAME, func(path string) (internal.RegisterableMetadata, error) { return OpenDataset(path) })
}

type WaveDataset struct {
	internal.MetadataInfo
	XMLName     xml.Name `xml:"WaveDataset"`
	Xmlns       string   `xml:"xmlns,attr"`
	Application struct {
		Text string `xml:",chardata"`
	} `xml:"application"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	DatasetType struct {
		Text string `xml:",chardata"`
	} `xml:"type"`
	TemplateAssetSourceName struct {
		Text string `xml:",chardata"`
	} `xml:"templateAssetSourceName"`
}

func (c *WaveDataset) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveDataset) Type() internal.MetadataType {
	return DATA_SET_NAME
}

func OpenDataset(path string) (*WaveDataset, error) {
	p := &WaveDataset{}
	return p, internal.ParseMetadataXml(p, path)
}
