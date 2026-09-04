package tab

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomTab"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// ActionOverride is one <actionOverrides> entry: the Lightning page assigned
// to a tab action. The Home tab's org-default page assignment is recorded as
// an override of its Tab action on standard-home.
type ActionOverride struct {
	ActionName           string       `xml:"actionName"`
	Comment              *string      `xml:"comment"`
	Content              *string      `xml:"content"`
	FormFactor           *string      `xml:"formFactor"`
	SkipRecordTypeSelect *BooleanText `xml:"skipRecordTypeSelect"`
	Type                 string       `xml:"type"`
}

type CustomTab struct {
	metadata.MetadataInfo
	XMLName         xml.Name         `xml:"CustomTab"`
	Xmlns           string           `xml:"xmlns,attr"`
	ActionOverrides []ActionOverride `xml:"actionOverrides"`
	CustomObject    *struct {
		Text string `xml:",chardata"`
	} `xml:"customObject"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	FlexiPage *struct {
		Text string `xml:",chardata"`
	} `xml:"flexiPage"`
	FrameHeight *struct {
		Text string `xml:",chardata"`
	} `xml:"frameHeight"`
	HasSidebar *struct {
		Text string `xml:",chardata"`
	} `xml:"hasSidebar"`
	Icon *struct {
		Text string `xml:",chardata"`
	} `xml:"icon"`
	Label *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	LwcComponent *struct {
		Text string `xml:",chardata"`
	} `xml:"lwcComponent"`
	Motif struct {
		Text string `xml:",chardata"`
	} `xml:"motif"`
	Page *struct {
		Text string `xml:",chardata"`
	} `xml:"page"`
	URL *struct {
		Text string `xml:",chardata"`
	} `xml:"url"`
	UrlEncodingKey *struct {
		Text string `xml:",chardata"`
	} `xml:"urlEncodingKey"`
}

func (c *CustomTab) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomTab) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomTab, error) {
	p := &CustomTab{}
	return p, metadata.ParseMetadataXml(p, path)
}
