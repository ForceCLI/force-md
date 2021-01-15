package profile

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type FieldPermissions struct {
	Editable BooleanText `xml:"editable"`
	Field    FieldName   `xml:"field"`
	Readable BooleanText `xml:"readable"`
}

type FieldPermissionsList []FieldPermissions

type ObjectPermissionsList []ObjectPermissions

type TabVisibilityList []TabVisibility

type UserPermissionList []UserPermission

type LayoutAssignmentList []LayoutAssignment

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

type UserPermission struct {
	Enabled BooleanText `xml:"enabled"`
	Name    struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
}

type LayoutAssignment struct {
	Layout struct {
		Text string `xml:",chardata"`
	} `xml:"layout"`
	RecordType *RecordType `xml:"recordType"`
}

type RecordType struct {
	Text string `xml:",chardata"`
}

type FieldName struct {
	Text string `xml:",chardata"`
}

type ObjectName struct {
	Text string `xml:",chardata"`
}

type BooleanText struct {
	Text string `xml:",chardata"`
}

type PersonAccountDefault struct {
	Text string `xml:",chardata"`
}

type Profile struct {
	XMLName                 xml.Name `xml:"Profile"`
	Xmlns                   string   `xml:"xmlns,attr"`
	ApplicationVisibilities []struct {
		Application struct {
			Text string `xml:",chardata"`
		} `xml:"application"`
		Default struct {
			Text string `xml:",chardata"`
		} `xml:"default"`
		Visible struct {
			Text string `xml:",chardata"`
		} `xml:"visible"`
	} `xml:"applicationVisibilities"`
	ClassAccesses []struct {
		ApexClass struct {
			Text string `xml:",chardata"`
		} `xml:"apexClass"`
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
	} `xml:"classAccesses"`
	Custom struct {
		Text string `xml:",chardata"`
	} `xml:"custom"`
	FieldPermissions FieldPermissionsList `xml:"fieldPermissions"`
	FlowAccesses     []struct {
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
		Flow struct {
			Text string `xml:",chardata"`
		} `xml:"flow"`
	} `xml:"flowAccesses"`
	LayoutAssignments LayoutAssignmentList  `xml:"layoutAssignments"`
	ObjectPermissions ObjectPermissionsList `xml:"objectPermissions"`
	PageAccesses      []struct {
		ApexPage struct {
			Text string `xml:",chardata"`
		} `xml:"apexPage"`
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
	} `xml:"pageAccesses"`
	RecordTypeVisibilities []struct {
		Default struct {
			Text string `xml:",chardata"`
		} `xml:"default"`
		PersonAccountDefault *PersonAccountDefault `xml:"personAccountDefault"`
		RecordType           struct {
			Text string `xml:",chardata"`
		} `xml:"recordType"`
		Visible struct {
			Text string `xml:",chardata"`
		} `xml:"visible"`
	} `xml:"recordTypeVisibilities"`
	TabVisibilities TabVisibilityList `xml:"tabVisibilities"`
	UserLicense     struct {
		Text string `xml:",chardata"`
	} `xml:"userLicense"`
	UserPermissions UserPermissionList `xml:"userPermissions"`
}

func (p *Profile) MetaCheck() {}

func Open(path string) (*Profile, error) {
	p := &Profile{}
	return p, internal.ParseMetadataXml(p, path)
}
