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

type RoleAndSubordinates struct {
	RoleAndSubordinate []string `xml:"roleAndSubordinate"`
}

type Roles struct {
	Role []string `xml:"role"`
}

type PublicGroups struct {
	PublicGroup []string `xml:"publicGroup"`
}

type Users struct {
	User []string `xml:"user"`
}

type Queue struct {
	metadata.MetadataInfo
	XMLName                xml.Name `xml:"Queue"`
	Xmlns                  string   `xml:"xmlns,attr"`
	DoesSendEmailToMembers struct {
		Text string `xml:",chardata"`
	} `xml:"doesSendEmailToMembers"`
	Email *string `xml:"email"`
	Name  struct {
		Text string `xml:",chardata"`
	} `xml:"name"`
	QueueMembers struct {
		PublicGroups        *PublicGroups        `xml:"publicGroups"`
		RoleAndSubordinates *RoleAndSubordinates `xml:"roleAndSubordinates"`
		Roles               *Roles               `xml:"roles"`
		Users               *Users               `xml:"users"`
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
