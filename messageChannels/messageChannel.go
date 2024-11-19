package messageChannel

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "LightningMessageChannel"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type LightningMessageChannel struct {
	internal.MetadataInfo
	XMLName     xml.Name `xml:"LightningMessageChannel"`
	Xmlns       string   `xml:"xmlns,attr"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	IsExposed struct {
		Text string `xml:",chardata"`
	} `xml:"isExposed"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *LightningMessageChannel) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *LightningMessageChannel) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*LightningMessageChannel, error) {
	p := &LightningMessageChannel{}
	return p, internal.ParseMetadataXml(p, path)
}
