package wave

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const DATA_FLOW_NAME = "WaveDataflow"

func init() {
	internal.TypeRegistry.Register(DATA_FLOW_NAME, func(path string) (metadata.RegisterableMetadata, error) { return OpenDataflow(path) })
}

type WaveDataflow struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"WaveDataflow"`
	Xmlns   string   `xml:"xmlns,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Content struct {
		Nil string `xml:"nil,attr"`
	} `xml:"content"`
	DataflowType struct {
		Text string `xml:",chardata"`
	} `xml:"dataflowType"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	Application struct {
		Text string `xml:",chardata"`
	} `xml:"application"`
}

func (c *WaveDataflow) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *WaveDataflow) Type() metadata.MetadataType {
	return DATA_FLOW_NAME
}

func OpenDataflow(path string) (*WaveDataflow, error) {
	p := &WaveDataflow{}
	return p, metadata.ParseMetadataXml(p, path)
}
