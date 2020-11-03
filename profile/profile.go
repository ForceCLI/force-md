package profile

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
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

type Profile struct {
	Declaration             string   `xml:"-"`
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
	FieldPermissions  []FieldPermission `xml:"fieldPermissions"`
	LayoutAssignments []struct {
		Layout struct {
			Text string `xml:",chardata"`
		} `xml:"layout"`
	} `xml:"layoutAssignments"`
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

func (p *Profile) Write(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	defer f.Close()
	fmt.Fprintln(f, p.Declaration)
	b, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		return errors.Wrap(err, "serializing profile")
	}
	_, err = f.Write(b)
	if err != nil {
		return errors.Wrap(err, "writing xml")
	}
	fmt.Fprintln(f, "")
	return nil
}

func ParseProfile(profilePath string) (*Profile, error) {
	f, err := os.Open(profilePath)
	if err != nil {
		return nil, errors.Wrap(err, "opening profile")
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	var doc Profile
	if err := dec.Decode(&doc); err != nil {
		return nil, errors.Wrap(err, "parsing xml")
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return &doc, errors.Wrap(err, "reading header")
	}
	b := bufio.NewReader(f)
	declaration, isPrefix, err := b.ReadLine()
	if isPrefix {
		return &doc, errors.New("reading xml declaration")
	}
	if err != nil {
		return &doc, errors.Wrap(err, "reading header")
	}
	doc.Declaration = string(declaration)

	return &doc, nil
}
