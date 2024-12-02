package labels

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomLabels"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CustomLabelList []CustomLabel

type CustomLabel struct {
	FullName         string `xml:"fullName"`
	Categories       string `xml:"categories"`
	Language         string `xml:"language"`
	Protected        string `xml:"protected"`
	ShortDescription string `xml:"shortDescription"`
	Value            string `xml:"value"`
}

type CustomLabels struct {
	metadata.MetadataInfo
	XMLName xml.Name        `xml:"CustomLabels"`
	Xmlns   string          `xml:"xmlns,attr"`
	Labels  CustomLabelList `xml:"labels"`
}

func (c *CustomLabels) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomLabels) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomLabels, error) {
	p := &CustomLabels{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (s *CustomLabels) GetLabels() CustomLabelList {
	return s.Labels
}
