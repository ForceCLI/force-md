package campaignInfluenceModels

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CampaignInfluenceModel"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// CampaignInfluenceModel sets how revenue is attributed to the campaigns
// that influenced an opportunity
type CampaignInfluenceModel struct {
	metadata.MetadataInfo
	XMLName          xml.Name     `xml:"CampaignInfluenceModel"`
	Xmlns            string       `xml:"xmlns,attr"`
	IsActive         *TextLiteral `xml:"isActive"`
	IsDefaultModel   TextLiteral  `xml:"isDefaultModel"`
	IsModelLocked    TextLiteral  `xml:"isModelLocked"`
	ModelDescription *TextLiteral `xml:"modelDescription"`
	Name             TextLiteral  `xml:"name"`
	RecordPreference *TextLiteral `xml:"recordPreference"`
}

func (c *CampaignInfluenceModel) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CampaignInfluenceModel) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CampaignInfluenceModel, error) {
	p := &CampaignInfluenceModel{}
	return p, metadata.ParseMetadataXml(p, path)
}
