package profile

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type FieldPermission struct {
	Editable struct {
		Text string `xml:",chardata"`
	} `xml:"editable"`
	Field struct {
		Text string `xml:",chardata"`
	} `xml:"field"`
	Readable struct {
		Text string `xml:",chardata"`
	} `xml:"readable"`
}

type ObjectPermissionsList []ObjectPermissions

type ObjectPermissions struct {
	AllowCreate      BooleanText `xml:"allowCreate"`
	AllowDelete      BooleanText `xml:"allowDelete"`
	AllowEdit        BooleanText `xml:"allowEdit"`
	AllowRead        BooleanText `xml:"allowRead"`
	ModifyAllRecords BooleanText `xml:"modifyAllRecords"`
	Object           ObjectName  `xml:"object"`
	ViewAllRecords   BooleanText `xml:"viewAllRecords"`
}

type RecordType struct {
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
	FieldPermissions []FieldPermission `xml:"fieldPermissions"`
	FlowAccesses     []struct {
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
		Flow struct {
			Text string `xml:",chardata"`
		} `xml:"flow"`
	} `xml:"flowAccesses"`
	LayoutAssignments []struct {
		Layout struct {
			Text string `xml:",chardata"`
		} `xml:"layout"`
		RecordType *RecordType `xml:"recordType"`
	} `xml:"layoutAssignments"`
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
	TabVisibilities []struct {
		Tab struct {
			Text string `xml:",chardata"`
		} `xml:"tab"`
		Visibility struct {
			Text string `xml:",chardata"`
		} `xml:"visibility"`
	} `xml:"tabVisibilities"`
	UserLicense struct {
		Text string `xml:",chardata"`
	} `xml:"userLicense"`
	UserPermissions []struct {
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"userPermissions"`
}

func (p *Profile) MetaCheck() {}

func Open(path string) (*Profile, error) {
	p := &Profile{}
	return p, internal.ParseMetadataXml(p, path)
}
