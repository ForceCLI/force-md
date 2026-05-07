package viewdefinitions

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ViewDefinition"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// ViewDefinition represents the ViewDefinition metadata type used by
// Slack apps. The accompanying .view file (YAML, holding the block-kit
// component tree) sits next to the .view-meta.xml file in the
// viewdefinitions directory and is read separately by consumers.
type ViewDefinition struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"ViewDefinition"`
	Xmlns   string   `xml:"xmlns,attr"`

	ApiVersion *struct {
		Text string `xml:",chardata"`
	} `xml:"apiVersion,omitempty"`
	IsProtected *struct {
		Text string `xml:",chardata"`
	} `xml:"isProtected,omitempty"`
	MasterLabel struct {
		Text string `xml:",chardata"`
	} `xml:"masterLabel"`
	TargetType *struct {
		Text string `xml:",chardata"`
	} `xml:"targetType,omitempty"`
}

func (v *ViewDefinition) SetMetadata(m metadata.MetadataInfo) {
	v.MetadataInfo = m
}

func Open(path string) (*ViewDefinition, error) {
	p := &ViewDefinition{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (v *ViewDefinition) Type() metadata.MetadataType {
	return NAME
}
