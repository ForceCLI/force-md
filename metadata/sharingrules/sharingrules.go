package sharingrules

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "SharingRules"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type CriteriaRuleList []CriteriaRule

type CriteriaRule struct {
	FullName                 string           `xml:"fullName"`
	AccessLevel              AccessLevel      `xml:"accessLevel"`
	AccountSettings          *AccountSettings `xml:"accountSettings"`
	Description              *Description     `xml:"description"`
	Label                    Label            `xml:"label"`
	SharedTo                 SharedTo         `xml:"sharedTo"`
	BooleanFilter            *BooleanFilter   `xml:"booleanFilter"`
	CriteriaItems            []CriteriaItem   `xml:"criteriaItems"`
	IncludeRecordsOwnedByAll *BooleanText     `xml:"includeRecordsOwnedByAll"`
}

type OwnerRuleList []OwnerRule

type OwnerRule struct {
	FullName        string           `xml:"fullName"`
	AccessLevel     AccessLevel      `xml:"accessLevel"`
	AccountSettings *AccountSettings `xml:"accountSettings"`
	Description     *Description     `xml:"description"`
	Label           Label            `xml:"label"`
	SharedTo        SharedTo         `xml:"sharedTo"`
	SharedFrom      SharedFrom       `xml:"sharedFrom"`
}

type GuestRuleList []GuestRule

type GuestRule struct {
	FullName               string       `xml:"fullName"`
	AccessLevel            AccessLevel  `xml:"accessLevel"`
	Description            *Description `xml:"description"`
	Label                  Label        `xml:"label"`
	SharedTo               SharedTo     `xml:"sharedTo"`
	CriteriaItems          CriteriaItem `xml:"criteriaItems"`
	IncludeHVUOwnedRecords *TextLiteral `xml:"includeHVUOwnedRecords"`
}

type SharingRules struct {
	metadata.MetadataInfo
	XMLName              xml.Name         `xml:"SharingRules"`
	Xmlns                string           `xml:"xmlns,attr"`
	SharingCriteriaRules CriteriaRuleList `xml:"sharingCriteriaRules"`
	SharingGuestRules    GuestRuleList    `xml:"sharingGuestRules"`
	SharingOwnerRules    OwnerRuleList    `xml:"sharingOwnerRules"`
}

func (c *SharingRules) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func Open(path string) (*SharingRules, error) {
	p := &SharingRules{}
	return p, metadata.ParseMetadataXml(p, path)
}

func (c *SharingRules) Type() metadata.MetadataType {
	return NAME
}

func (s *SharingRules) GetOwnerRules() []OwnerRule {
	return s.SharingOwnerRules
}

func (s *SharingRules) GetCriteriaRules() []CriteriaRule {
	return s.SharingCriteriaRules
}
