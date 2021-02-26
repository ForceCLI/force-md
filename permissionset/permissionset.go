package permissionset

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type ApexClass struct {
	ApexClass string `xml:"apexClass"`
	Enabled   string `xml:"enabled"`
}

type Description struct {
	Text string `xml:",chardata"`
}

type License struct {
	Text string `xml:",chardata"`
}

type TabSettings struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

type BooleanText struct {
	Text string `xml:",chardata"`
}

type ObjectPermissions struct {
	AllowCreate      BooleanText `xml:"allowCreate"`
	AllowDelete      BooleanText `xml:"allowDelete"`
	AllowEdit        BooleanText `xml:"allowEdit"`
	AllowRead        BooleanText `xml:"allowRead"`
	ModifyAllRecords BooleanText `xml:"modifyAllRecords"`
	Object           ObjectName  `xml:"object"`
	ViewAllRecords   BooleanText `xml:"viewAllRecords"`
}

type ObjectName struct {
	Text string `xml:",chardata"`
}

type FieldName struct {
	Text string `xml:",chardata"`
}

type ObjectPermissionsList []ObjectPermissions

type FieldPermissions struct {
	Editable BooleanText `xml:"editable"`
	Field    FieldName   `xml:"field"`
	Readable BooleanText `xml:"readable"`
}

type FieldPermissionsList []FieldPermissions

type PermissionSet struct {
	XMLName               xml.Name              `xml:"PermissionSet"`
	Xmlns                 string                `xml:"xmlns,attr"`
	ClassAccesses         []ApexClass           `xml:"classAccesses"`
	Description           *Description          `xml:"description"`
	FieldPermissions      FieldPermissionsList  `xml:"fieldPermissions"`
	HasActivationRequired BooleanText           `xml:"hasActivationRequired"`
	Label                 string                `xml:"label"`
	ObjectPermissions     ObjectPermissionsList `xml:"objectPermissions"`
	PageAccesses          []struct {
		ApexPage struct {
			Text string `xml:",chardata"`
		} `xml:"apexPage"`
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
	} `xml:"pageAccesses"`
	License           *License `xml:"license"`
	CustomPermissions []struct {
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"customPermissions"`
	TabSettings     []TabSettings `xml:"tabSettings"`
	UserPermissions []struct {
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"userPermissions"`
	ApplicationVisibilities []struct {
		Application struct {
			Text string `xml:",chardata"`
		} `xml:"application"`
		Visible struct {
			Text string `xml:",chardata"`
		} `xml:"visible"`
	} `xml:"applicationVisibilities"`
	RecordTypeVisibilities []struct {
		RecordType struct {
			Text string `xml:",chardata"`
		} `xml:"recordType"`
		Visible struct {
			Text string `xml:",chardata"`
		} `xml:"visible"`
	} `xml:"recordTypeVisibilities"`
}

func (p *PermissionSet) MetaCheck() {}

func Open(path string) (*PermissionSet, error) {
	p := &PermissionSet{}
	return p, internal.ParseMetadataXml(p, path)
}
