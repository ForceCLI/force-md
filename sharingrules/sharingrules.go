package sharingrules

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type CriteriaRule struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	AccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"accessLevel"`
	AccountSettings *struct {
		CaseAccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"caseAccessLevel"`
		ContactAccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"contactAccessLevel"`
		OpportunityAccessLevel struct {
			Text string `xml:",chardata"`
		} `xml:"opportunityAccessLevel"`
	} `xml:"accountSettings"`
	Description *struct {
		Text string `xml:",innerxml"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",innerxml"`
	} `xml:"label"`
	SharedTo struct {
		Group *struct {
			Text string `xml:",chardata"`
		} `xml:"group"`
		Role *struct {
			Text string `xml:",chardata"`
		} `xml:"role"`
		AllInternalUsers *struct {
		} `xml:"allInternalUsers"`
		RoleAndSubordinates *struct {
			Text string `xml:",chardata"`
		} `xml:"roleAndSubordinates"`
	} `xml:"sharedTo"`
	CriteriaItems []struct {
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Operation struct {
			Text string `xml:",chardata"`
		} `xml:"operation"`
		Value struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"criteriaItems"`
	BooleanFilter *struct {
		Text string `xml:",chardata"`
	} `xml:"booleanFilter"`
}

type OwnerRule struct {
	FullName struct {
		Text string `xml:",chardata"`
	} `xml:"fullName"`
	AccessLevel struct {
		Text string `xml:",chardata"`
	} `xml:"accessLevel"`
	Description *struct {
		Text string `xml:",innerxml"`
	} `xml:"description"`
	Label struct {
		Text string `xml:",innerxml"`
	} `xml:"label"`
	SharedTo struct {
		Role *struct {
			Text string `xml:",chardata"`
		} `xml:"role"`
		Group *struct {
			Text string `xml:",chardata"`
		} `xml:"group"`
		RoleAndSubordinates *struct {
			Text string `xml:",chardata"`
		} `xml:"roleAndSubordinates"`
	} `xml:"sharedTo"`
	SharedFrom struct {
		RoleAndSubordinates *struct {
			Text string `xml:",chardata"`
		} `xml:"roleAndSubordinates"`
		Group *struct {
			Text string `xml:",chardata"`
		} `xml:"group"`
		Queue *struct {
			Text string `xml:",chardata"`
		} `xml:"queue"`
		Role *struct {
			Text string `xml:",chardata"`
		} `xml:"role"`
		AllInternalUsers *struct {
		} `xml:"allInternalUsers"`
	} `xml:"sharedFrom"`
}

type SharingRules struct {
	XMLName              xml.Name       `xml:"SharingRules"`
	Xmlns                string         `xml:"xmlns,attr"`
	SharingCriteriaRules []CriteriaRule `xml:"sharingCriteriaRules"`
	SharingOwnerRules    []OwnerRule    `xml:"sharingOwnerRules"`
}

func (p *SharingRules) MetaCheck() {}

func Open(path string) (*SharingRules, error) {
	p := &SharingRules{}
	return p, internal.ParseMetadataXml(p, path)
}

func (s *SharingRules) GetOwnerRules() []OwnerRule {
	return s.SharingOwnerRules
}

func (s *SharingRules) GetCriteriaRules() []CriteriaRule {
	return s.SharingCriteriaRules
}
