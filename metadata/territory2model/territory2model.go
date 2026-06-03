package territory2model

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Territory2Model"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Territory2Model struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"Territory2Model"`
	Xmlns   string   `xml:"xmlns,attr"`
	Name    struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
}

func (c *Territory2Model) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Territory2Model) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Territory2Model, error) {
	p := &Territory2Model{}
	return p, metadata.ParseMetadataXml(p, path)
}
