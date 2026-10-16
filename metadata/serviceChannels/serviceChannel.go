package serviceChannels

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ServiceChannel"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ServiceChannel struct {
	metadata.MetadataInfo
	XMLName                           xml.Name                           `xml:"ServiceChannel"`
	Xmlns                             string                             `xml:"xmlns,attr"`
	AcwExtensionDuration              *IntegerText                       `xml:"acwExtensionDuration"`
	AfterConvoWorkMaxTime             *IntegerText                       `xml:"afterConvoWorkMaxTime"`
	CapacityModel                     *TextLiteral                       `xml:"capacityModel"`
	CustomRequestedDateTimeField      *TextLiteral                       `xml:"customRequestedDateTimeField"`
	DoesCheckCapOnOwnerChange         *BooleanText                       `xml:"doesCheckCapOnOwnerChange"`
	DoesCheckCapOnStatusChange        *BooleanText                       `xml:"doesCheckCapOnStatusChange"`
	DoesMinimizeWidgetOnAccept        *BooleanText                       `xml:"doesMinimizeWidgetOnAccept"`
	HasAcwExtensionEnabled            *BooleanText                       `xml:"hasAcwExtensionEnabled"`
	HasAfterConvoWorkTimer            *BooleanText                       `xml:"hasAfterConvoWorkTimer"`
	HasAutoAcceptEnabled              *BooleanText                       `xml:"hasAutoAcceptEnabled"`
	InteractionComponent              *TextLiteral                       `xml:"interactionComponent"`
	IsInterruptible                   *BooleanText                       `xml:"isInterruptible"`
	Label                             *TextLiteral                       `xml:"label"`
	MaxExtensions                     *IntegerText                       `xml:"maxExtensions"`
	RelatedEntityType                 *TextLiteral                       `xml:"relatedEntityType"`
	SecondaryRoutingPriorityField     *TextLiteral                       `xml:"secondaryRoutingPriorityField"`
	ServiceChannelFieldPriorities     []ServiceChannelFieldPriority      `xml:"serviceChannelFieldPriorities"`
	ServiceChannelStatusFieldMappings []ServiceChannelStatusFieldMapping `xml:"serviceChannelStatusFieldMappings"`
	StatusField                       *TextLiteral                       `xml:"statusField"`
}

type ServiceChannelFieldPriority struct {
	Priority *IntegerText `xml:"priority"`
	Value    *TextLiteral `xml:"value"`
}

type ServiceChannelStatusFieldMapping struct {
	Type  *TextLiteral `xml:"type"`
	Value *TextLiteral `xml:"value"`
}

func (c *ServiceChannel) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ServiceChannel) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ServiceChannel, error) {
	c := &ServiceChannel{}
	return c, metadata.ParseMetadataXml(c, path)
}
