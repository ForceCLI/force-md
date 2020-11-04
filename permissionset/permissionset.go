package permissionset

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type PermissionSet struct {
	XMLName       xml.Name `xml:"PermissionSet"`
	Xmlns         string   `xml:"xmlns,attr"`
	ClassAccesses []struct {
		ApexClass struct {
			Text string `xml:",chardata"`
		} `xml:"apexClass"`
		Enabled struct {
			Text string `xml:",chardata"`
		} `xml:"enabled"`
	} `xml:"classAccesses"`
	Description struct {
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
	License struct {
		Text string `xml:",chardata"`
	} `xml:"license"`
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

func (r *PermissionSet) Write(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	defer f.Close()
	fmt.Fprintln(f, `<?xml version="1.0" encoding="UTF-8"?>`)
	b, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return errors.Wrap(err, "serializing permission set")
	}
	_, err = f.Write(b)
	if err != nil {
		return errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(f, "")
	return nil
}

func ParsePermissionSet(permissionSetPath string) (*PermissionSet, error) {
	f, err := os.Open(permissionSetPath)
	if err != nil {
		return nil, errors.Wrap(err, "opening permission set")
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	var doc PermissionSet
	if err := dec.Decode(&doc); err != nil {
		return nil, errors.Wrap(err, "parsing xml")
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return &doc, errors.Wrap(err, "reading header")
	}

	return &doc, nil
}
