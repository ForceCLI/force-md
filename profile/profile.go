package profile

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type FieldPermissions struct {
	Editable BooleanText `xml:"editable"`
	Field    FieldName   `xml:"field"`
	Readable BooleanText `xml:"readable"`
}

type CustomPermissions struct {
	Enabled BooleanText    `xml:"enabled"`
	Name    PermissionName `xml:"name"`
}

type CustomPermissionsList []CustomPermissions

type FieldPermissionsList []FieldPermissions

type FlowAccessList []FlowAccess

type ObjectPermissionsList []ObjectPermissions

type ApplicationVisibilityList []ApplicationVisibility

type TabVisibilityList []TabVisibility

type UserPermissionList []UserPermission

type LayoutAssignmentList []LayoutAssignment

type LoginFlowsList []LoginFlow

type LoginIpRangeList []LoginIpRange

type PageAccessList []PageAccess

type RecordTypeVisibilityList []RecordTypeVisibility

type ObjectPermissions struct {
	AllowCreate      BooleanText `xml:"allowCreate"`
	AllowDelete      BooleanText `xml:"allowDelete"`
	AllowEdit        BooleanText `xml:"allowEdit"`
	AllowRead        BooleanText `xml:"allowRead"`
	ModifyAllRecords BooleanText `xml:"modifyAllRecords"`
	Object           ObjectName  `xml:"object"`
	ViewAllRecords   BooleanText `xml:"viewAllRecords"`
}

type TabVisibility struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

type PageAccess struct {
	ApexPage string      `xml:"apexPage"`
	Enabled  BooleanText `xml:"enabled"`
}

type UserPermission struct {
	Enabled BooleanText `xml:"enabled"`
	Name    string      `xml:"name"`
}

type LayoutAssignment struct {
	Layout     string      `xml:"layout"`
	RecordType *RecordType `xml:"recordType"`
}

type FlowAccess struct {
	Enabled BooleanText `xml:"enabled"`
	Flow    string      `xml:"flow"`
}

type LoginFlow struct {
	Flow                string      `xml:"flow"`
	FlowType            string      `xml:"flowType"`
	FriendlyName        string      `xml:"friendlyName"`
	UILoginFlowType     string      `xml:"uiLoginFlowType"`
	UseLightningRuntime BooleanText `xml:"useLightningRuntime"`
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
	Default struct {
		Text string `xml:",chardata"`
	} `xml:"default"`
	PersonAccountDefault *PersonAccountDefault `xml:"personAccountDefault"`
	RecordType           string                `xml:"recordType"`
	Visible              BooleanText           `xml:"visible"`
}

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

type PersonAccountDefault struct {
	Text string `xml:",chardata"`
}

type ApexClass struct {
	ApexClass string      `xml:"apexClass"`
	Enabled   BooleanText `xml:"enabled"`
}

type ApexClassList []ApexClass

type Profile struct {
	XMLName                 xml.Name                  `xml:"Profile"`
	Xmlns                   string                    `xml:"xmlns,attr"`
	ApplicationVisibilities ApplicationVisibilityList `xml:"applicationVisibilities"`
	ClassAccesses           ApexClassList             `xml:"classAccesses"`
	Custom                  struct {
		Text string `xml:",chardata"`
	} `xml:"custom"`
	CustomPermissions CustomPermissionsList `xml:"customPermissions"`
	Description       *string               `xml:"description"`
	FieldPermissions  FieldPermissionsList  `xml:"fieldPermissions"`
	FlowAccesses      FlowAccessList        `xml:"flowAccesses"`
	LayoutAssignments LayoutAssignmentList  `xml:"layoutAssignments"`
	LoginFlows        LoginFlowsList        `xml:"loginFlows"`
	LoginHours        *struct {
		Text string `xml:",chardata"`
	} `xml:"loginHours"`
	LoginIPRanges          LoginIpRangeList         `xml:"loginIpRanges"`
	ObjectPermissions      ObjectPermissionsList    `xml:"objectPermissions"`
	PageAccesses           PageAccessList           `xml:"pageAccesses"`
	RecordTypeVisibilities RecordTypeVisibilityList `xml:"recordTypeVisibilities"`
	TabVisibilities        TabVisibilityList        `xml:"tabVisibilities"`
	UserLicense            string                   `xml:"userLicense"`
	UserPermissions        UserPermissionList       `xml:"userPermissions"`
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
