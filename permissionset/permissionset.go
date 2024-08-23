package permissionset

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
)

type Description struct {
	Text string `xml:",innerxml"`
}

type License struct {
	Text string `xml:",chardata"`
}

type ApexClass struct {
	ApexClass string      `xml:"apexClass"`
	Enabled   BooleanText `xml:"enabled"`
}

type ApexClassList []ApexClass

type ObjectPermissions struct {
	AllowCreate      BooleanText `xml:"allowCreate"`
	AllowDelete      BooleanText `xml:"allowDelete"`
	AllowEdit        BooleanText `xml:"allowEdit"`
	AllowRead        BooleanText `xml:"allowRead"`
	ModifyAllRecords BooleanText `xml:"modifyAllRecords"`
	Object           string      `xml:"object"`
	ViewAllRecords   BooleanText `xml:"viewAllRecords"`
}

type ObjectPermissionsList []ObjectPermissions

type FieldPermissions struct {
	Editable BooleanText `xml:"editable"`
	Field    string      `xml:"field"`
	Readable BooleanText `xml:"readable"`
}

type FieldPermissionsList []FieldPermissions

type TabSettings struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

type TabSettingsList []TabSettings

type UserPermission struct {
	Enabled BooleanText `xml:"enabled"`
	Name    string      `xml:"name"`
}

type UserPermissionList []UserPermission

type CustomPermission struct {
	Enabled BooleanText `xml:"enabled"`
	Name    string      `xml:"name"`
}

type CustomPermissionList []CustomPermission

type CustomMetadataType struct {
	Enabled BooleanText `xml:"enabled"`
	Name    string      `xml:"name"`
}

type CustomMetadataTypeList []CustomMetadataType

type CustomSetting struct {
	Enabled BooleanText `xml:"enabled"`
	Name    string      `xml:"name"`
}

type CustomSettingList []CustomSetting

type PageAccess struct {
	ApexPage string      `xml:"apexPage"`
	Enabled  BooleanText `xml:"enabled"`
}

type PageAccessList []PageAccess

type ApplicationVisibility struct {
	Application string      `xml:"application"`
	Visible     BooleanText `xml:"visible"`
}

type ApplicationVisibilityList []ApplicationVisibility

type ExternalCredentialPrincipalAccess struct {
	Enabled                     BooleanText `xml:"enabled"`
	ExternalCredentialPrincipal string      `xml:"externalCredentialPrincipal"`
}

type ExternalCredentialPrincipalAccessList []ExternalCredentialPrincipalAccess

type FlowAccess struct {
	Enabled BooleanText `xml:"enabled"`
	Flow    string      `xml:"flow"`
}

type FlowAccessList []FlowAccess

type RecordTypeVisibility struct {
	RecordType string      `xml:"recordType"`
	Visible    BooleanText `xml:"visible"`
}

type RecordTypeVisibilityList []RecordTypeVisibility

type PermissionSet struct {
	XMLName                             xml.Name                              `xml:"PermissionSet"`
	Xmlns                               string                                `xml:"xmlns,attr"`
	ApplicationVisibilities             ApplicationVisibilityList             `xml:"applicationVisibilities"`
	ClassAccesses                       ApexClassList                         `xml:"classAccesses"`
	CustomMetadataTypeAccesses          CustomMetadataTypeList                `xml:"customMetadataTypeAccesses"`
	CustomPermissions                   CustomPermissionList                  `xml:"customPermissions"`
	CustomSettingAccesses               CustomSettingList                     `xml:"customSettingAccesses"`
	Description                         *Description                          `xml:"description"`
	FieldPermissions                    FieldPermissionsList                  `xml:"fieldPermissions"`
	FlowAccesses                        FlowAccessList                        `xml:"flowAccesses"`
	ExternalCredentialPrincipalAccesses ExternalCredentialPrincipalAccessList `xml:"externalCredentialPrincipalAccesses"`
	HasActivationRequired               BooleanText                           `xml:"hasActivationRequired"`
	Label                               string                                `xml:"label"`
	License                             *License                              `xml:"license"`
	ObjectPermissions                   ObjectPermissionsList                 `xml:"objectPermissions"`
	PageAccesses                        PageAccessList                        `xml:"pageAccesses"`
	RecordTypeVisibilities              RecordTypeVisibilityList              `xml:"recordTypeVisibilities"`
	TabSettings                         TabSettingsList                       `xml:"tabSettings"`
	UserPermissions                     UserPermissionList                    `xml:"userPermissions"`
}

func (p *PermissionSet) MetaCheck() {}

func Open(path string) (*PermissionSet, error) {
	p := &PermissionSet{}
	return p, internal.ParseMetadataXml(p, path)
}
