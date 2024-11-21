package queue

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

const NAME = "Queue"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (internal.RegisterableMetadata, error) { return Open(path) })
}

type Queue struct {
	internal.MetadataInfo
	XMLName                xml.Name `xml:"Queue"`
	Xmlns                  string   `xml:"xmlns,attr"`
	DoesSendEmailToMembers struct {
		Text string `xml:",chardata"`
	} `xml:"doesSendEmailToMembers"`
	Email struct {
		Text string `xml:",chardata"`
	} `xml:"email"`
	Name struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	QueueMembers struct {
		RoleAndSubordinates struct {
			RoleAndSubordinate []struct {
				Text string `xml:",chardata"`
			} `xml:"roleAndSubordinate"`
		} `xml:"roleAndSubordinates"`
		Roles struct {
			Role []struct {
				Text string `xml:",chardata"`
			} `xml:"role"`
		} `xml:"roles"`
		PublicGroups struct {
			PublicGroup []struct {
				Text string `xml:",chardata"`
			} `xml:"publicGroup"`
		} `xml:"publicGroups"`
	} `xml:"queueMembers"`
	QueueSobject []struct {
		SobjectType struct {
			Text string `xml:",chardata"`
		} `xml:"sobjectType"`
	} `xml:"queueSobject"`
}

func (c *Queue) SetMetadata(m internal.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Queue) Type() internal.MetadataType {
	return NAME
}

func Open(path string) (*Queue, error) {
	p := &Queue{}
	return p, internal.ParseMetadataXml(p, path)
}
