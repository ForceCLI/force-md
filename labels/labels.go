package labels

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

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
	internal.MetadataInfo
	XMLName xml.Name        `xml:"CustomLabels"`
	Xmlns   string          `xml:"xmlns,attr"`
	Labels  CustomLabelList `xml:"labels"`
}

func (c *CustomLabels) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*CustomLabels, error) {
	p := &CustomLabels{}
	return p, internal.ParseMetadataXml(p, path)
}

func (s *CustomLabels) GetLabels() CustomLabelList {
	return s.Labels
}
