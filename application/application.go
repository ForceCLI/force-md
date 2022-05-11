package application

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type ProfileActionOverride struct {
	ActionName        string  `xml:"actionName"`
	Content           string  `xml:"content"`
	FormFactor        string  `xml:"formFactor"`
	PageOrSobjectType string  `xml:"pageOrSobjectType"`
	RecordType        *string `xml:"recordType"`
	Type              string  `xml:"type"`
	Profile           string  `xml:"profile"`
}

type ProfileActionOverrideList []ProfileActionOverride

type CustomApplication struct {
	XMLName xml.Name `xml:"CustomApplication"`
	Xmlns   string   `xml:"xmlns,attr"`
	Brand   *struct {
		HeaderColor struct {
			Text string `xml:",chardata"`
		} `xml:"headerColor"`
		Logo *struct {
			Text string `xml:",chardata"`
		} `xml:"logo"`
		LogoVersion *struct {
			Text string `xml:",chardata"`
		} `xml:"logoVersion"`
		ShouldOverrideOrgTheme struct {
			Text string `xml:",chardata"`
		} `xml:"shouldOverrideOrgTheme"`
	} `xml:"brand"`
	DefaultLandingTab *struct {
		Text string `xml:",chardata"`
	} `xml:"defaultLandingTab"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	ActionOverrides []struct {
		ActionName struct {
			Text string `xml:",chardata"`
		} `xml:"actionName"`
		Comment struct {
			Text string `xml:",chardata"`
		} `xml:"comment"`
		Content struct {
			Text string `xml:",chardata"`
		} `xml:"content"`
		FormFactor struct {
			Text string `xml:",chardata"`
		} `xml:"formFactor"`
		SkipRecordTypeSelect struct {
			Text string `xml:",chardata"`
		} `xml:"skipRecordTypeSelect"`
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
		PageOrSobjectType struct {
			Text string `xml:",chardata"`
		} `xml:"pageOrSobjectType"`
	} `xml:"actionOverrides"`
	FormFactors []struct {
		Text string `xml:",chardata"`
	} `xml:"formFactors"`
	IsNavAutoTempTabsDisabled *struct {
		Text string `xml:",chardata"`
	} `xml:"isNavAutoTempTabsDisabled"`
	IsNavPersonalizationDisabled *struct {
		Text string `xml:",chardata"`
	} `xml:"isNavPersonalizationDisabled"`
	IsNavTabPersistenceDisabled *struct {
		Text string `xml:",chardata"`
	} `xml:"isNavTabPersistenceDisabled"`
	Label *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Logo *struct {
		Text string `xml:",chardata"`
	} `xml:"logo"`
	NavType *struct {
		Text string `xml:",chardata"`
	} `xml:"navType"`
	ProfileActionOverrides ProfileActionOverrideList `xml:"profileActionOverrides"`
	SetupExperience        *struct {
		Text string `xml:",chardata"`
	} `xml:"setupExperience"`
	Tabs   []TextLiteral `xml:"tabs"`
	UiType *struct {
		Text string `xml:",chardata"`
	} `xml:"uiType"`
	UtilityBar *struct {
		Text string `xml:",chardata"`
	} `xml:"utilityBar"`
	WorkspaceConfig *struct {
		Mappings []struct {
			FieldName *string `xml:"fieldName"`
			Tab       *string `xml:"tab"`
		} `xml:"mappings"`
	} `xml:"workspaceConfig"`
}

func (p *CustomApplication) MetaCheck() {}

func Open(path string) (*CustomApplication, error) {
	p := &CustomApplication{}
	return p, internal.ParseMetadataXml(p, path)
}
