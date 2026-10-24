package application

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "CustomApplication"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ProfileActionOverride struct {
	ActionName        string  `xml:"actionName"`
	Content           *string `xml:"content"`
	FormFactor        string  `xml:"formFactor"`
	PageOrSobjectType string  `xml:"pageOrSobjectType"`
	RecordType        *string `xml:"recordType"`
	Type              string  `xml:"type"`
	Profile           string  `xml:"profile"`
}

type ActionOverride struct {
	ActionName struct {
		Text string `xml:",chardata"`
	} `xml:"actionName"`
	Comment *struct {
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
}

type ProfileActionOverrideList []ProfileActionOverride

type ConsoleConfig struct {
	DetailPageRefreshMethod *TextLiteral `xml:"detailPageRefreshMethod"`
	FooterColor             *TextLiteral `xml:"footerColor"`
	HeaderColor             *TextLiteral `xml:"headerColor"`
	KeyboardShortcuts       *struct {
		CustomShortcuts []struct {
			Action      TextLiteral  `xml:"action"`
			Active      TextLiteral  `xml:"active"`
			Description *TextLiteral `xml:"description"`
			EventName   TextLiteral  `xml:"eventName"`
			KeyCommand  TextLiteral  `xml:"keyCommand"`
		} `xml:"customShortcuts"`
		DefaultShortcuts []struct {
			Action     TextLiteral `xml:"action"`
			Active     TextLiteral `xml:"active"`
			KeyCommand TextLiteral `xml:"keyCommand"`
		} `xml:"defaultShortcuts"`
	} `xml:"keyboardShortcuts"`
	ListPlacement *struct {
		Height   *TextLiteral `xml:"height"`
		Location TextLiteral  `xml:"location"`
		Units    *TextLiteral `xml:"units"`
		Width    *TextLiteral `xml:"width"`
	} `xml:"listPlacement"`
	ListRefreshMethod  *TextLiteral  `xml:"listRefreshMethod"`
	PrimaryTabColor    *TextLiteral  `xml:"primaryTabColor"`
	WhitelistedDomains []TextLiteral `xml:"whitelistedDomains"`
}

type Preferences struct {
	EnableCustomizeMyTabs        TextLiteral `xml:"enableCustomizeMyTabs"`
	EnableKeyboardShortcuts      TextLiteral `xml:"enableKeyboardShortcuts"`
	EnableListViewHover          TextLiteral `xml:"enableListViewHover"`
	EnableListViewReskin         TextLiteral `xml:"enableListViewReskin"`
	EnableMultiMonitorComponents TextLiteral `xml:"enableMultiMonitorComponents"`
	EnablePinTabs                TextLiteral `xml:"enablePinTabs"`
	EnableTabHover               TextLiteral `xml:"enableTabHover"`
	EnableTabLimits              TextLiteral `xml:"enableTabLimits"`
	SaveUserSessions             TextLiteral `xml:"saveUserSessions"`
}

type CustomApplication struct {
	metadata.MetadataInfo
	XMLName         xml.Name         `xml:"CustomApplication"`
	Xmlns           string           `xml:"xmlns,attr"`
	ActionOverrides []ActionOverride `xml:"actionOverrides"`
	Brand           *struct {
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
	ConsoleConfig     *ConsoleConfig `xml:"consoleConfig"`
	DefaultLandingTab *struct {
		Text string `xml:",chardata"`
	} `xml:"defaultLandingTab"`
	Description *struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
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
	IsOmniPinnedViewEnabled *TextLiteral `xml:"isOmniPinnedViewEnabled"`
	IsServiceCloudConsole   *TextLiteral `xml:"isServiceCloudConsole"`
	Label                   *struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	Logo *struct {
		Text string `xml:",chardata"`
	} `xml:"logo"`
	NavType *struct {
		Text string `xml:",chardata"`
	} `xml:"navType"`
	Preferences            *Preferences              `xml:"preferences"`
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

func (c *CustomApplication) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *CustomApplication) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*CustomApplication, error) {
	p := &CustomApplication{}
	return p, metadata.ParseMetadataXml(p, path)
}
