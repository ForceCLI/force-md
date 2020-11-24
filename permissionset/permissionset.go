package permissionset

import (
	"encoding/xml"

	"github.com/octoberswimmer/force-md/internal"
)

type ApexClass struct {
	ApexClass string `xml:"apexClass"`
	Enabled   string `xml:"enabled"`
}

type License struct {
	Text string `xml:",chardata"`
}

type PermissionSet struct {
	XMLName       xml.Name    `xml:"PermissionSet"`
	Xmlns         string      `xml:"xmlns,attr"`
	ClassAccesses []ApexClass `xml:"classAccesses"`
	Description   struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	FieldPermissions []struct {
		Editable struct {
			Text string `xml:",chardata"`
		} `xml:"editable"`
		Field struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
		Readable struct {
			Text string `xml:",chardata"`
		} `xml:"readable"`
	} `xml:"fieldPermissions"`
	HasActivationRequired struct {
		Text string `xml:",chardata"`
	} `xml:"hasActivationRequired"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	ObjectPermissions []struct {
		AllowCreate struct {
			Text string `xml:",chardata"`
		} `xml:"allowCreate"`
		AllowDelete struct {
			Text string `xml:",chardata"`
		} `xml:"allowDelete"`
		AllowEdit struct {
			Text string `xml:",chardata"`
		} `xml:"allowEdit"`
		AllowRead struct {
			Text string `xml:",chardata"`
		} `xml:"allowRead"`
		ModifyAllRecords struct {
			Text string `xml:",chardata"`
		} `xml:"modifyAllRecords"`
		Object struct {
			Text string `xml:",chardata"`
		} `xml:"object"`
		ViewAllRecords struct {
			Text string `xml:",chardata"`
		} `xml:"viewAllRecords"`
	} `xml:"objectPermissions"`
	PageAccesses []struct {
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
	UserPermissions []struct {
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"userPermissions"`
	TabSettings []struct {
		Tab struct {
			Text string `xml:",chardata"`
		} `xml:"tab"`
		Visibility struct {
			Text string `xml:",chardata"`
		} `xml:"visibility"`
	} `xml:"tabSettings"`
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
