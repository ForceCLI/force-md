package queue

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "Queue"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type Queue struct {
	metadata.MetadataInfo
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

func (c *Queue) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *Queue) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*Queue, error) {
	p := &Queue{}
	return p, metadata.ParseMetadataXml(p, path)
}
