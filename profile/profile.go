package profile

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
	"github.com/octoberswimmer/force-md/permissionset"
)

type ApplicationVisibilityList []ApplicationVisibility

type TabVisibilityList []TabVisibility

type LayoutAssignmentList []LayoutAssignment

type LoginFlowsList []LoginFlow

type LoginIpRangeList []LoginIpRange

type TabVisibility struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

type LayoutAssignment struct {
	Layout     string      `xml:"layout"`
	RecordType *RecordType `xml:"recordType"`
}

type LoginFlow struct {
	Flow                *string     `xml:"flow"`
	FlowType            string      `xml:"flowType"`
	FriendlyName        string      `xml:"friendlyName"`
	UILoginFlowType     string      `xml:"uiLoginFlowType"`
	UseLightningRuntime BooleanText `xml:"useLightningRuntime"`
	VFFlowPage          *string     `xml:"vfFlowPage"`
	VFFlowPageTitle     *string     `xml:"vfFlowPageTitle"`
}

type LoginIpRange struct {
	EndAddress   string `xml:"endAddress"`
	StartAddress string `xml:"startAddress"`
}

type ApplicationVisibility struct {
	Application string      `xml:"application"`
	Default     BooleanText `xml:"default"`
	Visible     BooleanText `xml:"visible"`
}

type RecordTypeVisibility struct {
	Default              BooleanText  `xml:"default"`
	PersonAccountDefault *BooleanText `xml:"personAccountDefault"`
	permissionset.RecordTypeVisibility
}

type RecordTypeVisibilityList []RecordTypeVisibility

type RecordType struct {
	Text string `xml:",chardata"`
}

type FieldName struct {
	Text string `xml:",chardata"`
}

type PermissionName struct {
	Text string `xml:",chardata"`
}

type ObjectName struct {
	Text string `xml:",chardata"`
}

type Profile struct {
	XMLName                 xml.Name                    `xml:"Profile"`
	Xmlns                   string                      `xml:"xmlns,attr"`
	ApplicationVisibilities ApplicationVisibilityList   `xml:"applicationVisibilities"`
	ClassAccesses           permissionset.ApexClassList `xml:"classAccesses"`
	Custom                  struct {
		Text string `xml:",chardata"`
	} `xml:"custom"`
	CustomMetadataTypeAccesses permissionset.CustomMetadataTypeList `xml:"customMetadataTypeAccesses"`
	CustomPermissions          permissionset.CustomPermissionList   `xml:"customPermissions"`
	CustomSettingAccesses      permissionset.CustomSettingList      `xml:"customSettingAccesses"`
	Description                *string                              `xml:"description"`
	FieldPermissions           permissionset.FieldPermissionsList   `xml:"fieldPermissions"`
	FlowAccesses               permissionset.FlowAccessList         `xml:"flowAccesses"`
	LayoutAssignments          LayoutAssignmentList                 `xml:"layoutAssignments"`
	LoginFlows                 *LoginFlow                           `xml:"loginFlows"`
	LoginHours                 *struct {
		Text string `xml:",chardata"`
	} `xml:"loginHours"`
	LoginIPRanges          LoginIpRangeList                    `xml:"loginIpRanges"`
	ObjectPermissions      permissionset.ObjectPermissionsList `xml:"objectPermissions"`
	PageAccesses           permissionset.PageAccessList        `xml:"pageAccesses"`
	RecordTypeVisibilities RecordTypeVisibilityList            `xml:"recordTypeVisibilities"`
	TabVisibilities        TabVisibilityList                   `xml:"tabVisibilities"`
	UserLicense            string                              `xml:"userLicense"`
	UserPermissions        permissionset.UserPermissionList    `xml:"userPermissions"`
}

func NewBooleanText(val string) BooleanText {
	return BooleanText{
		Text: val,
	}
}

func (p *Profile) MetaCheck() {}

func Open(path string) (*Profile, error) {
	p := &Profile{}
	return p, internal.ParseMetadataXml(p, path)
}
