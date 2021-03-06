package permissionset

import (
	"encoding/xml"

	. "github.com/octoberswimmer/force-md/general"
	"github.com/octoberswimmer/force-md/internal"
)

type Description struct {
	Text string `xml:",chardata"`
}

type License struct {
	Text string `xml:",chardata"`
}

type ObjectName struct {
	Text string `xml:",chardata"`
}

type FieldName struct {
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
	Object           ObjectName  `xml:"object"`
	ViewAllRecords   BooleanText `xml:"viewAllRecords"`
}

type ObjectPermissionsList []ObjectPermissions

type FieldPermissions struct {
	Editable BooleanText `xml:"editable"`
	Field    FieldName   `xml:"field"`
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

type RecordType struct {
	RecordType string      `xml:"recordType"`
	Visible    BooleanText `xml:"visible"`
}

type RecordTypeList []RecordType

type CustomPermission struct {
	Enabled BooleanText `xml:"enabled"`
	Name    string      `xml:"name"`
}

type CustomPermissionList []CustomPermission

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

type PermissionSet struct {
	XMLName                 xml.Name                  `xml:"PermissionSet"`
	Xmlns                   string                    `xml:"xmlns,attr"`
	ClassAccesses           ApexClassList             `xml:"classAccesses"`
	Description             *Description              `xml:"description"`
	FieldPermissions        FieldPermissionsList      `xml:"fieldPermissions"`
	HasActivationRequired   BooleanText               `xml:"hasActivationRequired"`
	Label                   string                    `xml:"label"`
	ObjectPermissions       ObjectPermissionsList     `xml:"objectPermissions"`
	PageAccesses            PageAccessList            `xml:"pageAccesses"`
	License                 *License                  `xml:"license"`
	CustomPermissions       CustomPermissionList      `xml:"customPermissions"`
	UserPermissions         UserPermissionList        `xml:"userPermissions"`
	ApplicationVisibilities ApplicationVisibilityList `xml:"applicationVisibilities"`
	RecordTypeVisibilities  RecordTypeList            `xml:"recordTypeVisibilities"`
	TabSettings             TabSettingsList           `xml:"tabSettings"`
}

func (p *PermissionSet) MetaCheck() {}

func Open(path string) (*PermissionSet, error) {
	p := &PermissionSet{}
	return p, internal.ParseMetadataXml(p, path)
}
