package servicePresenceStatuses

import (
	"encoding/xml"

	. "github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ServicePresenceStatus"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

type ServicePresenceStatus struct {
	metadata.MetadataInfo
	XMLName  xml.Name                       `xml:"ServicePresenceStatus"`
	Xmlns    string                         `xml:"xmlns,attr"`
	Channels *ServicePresenceStatusChannels `xml:"channels"`
	Label    *TextLiteral                   `xml:"label"`
}

// ServicePresenceStatusChannels lists the service channels, by developer
// name, an agent in the status receives work from. A status with no channels
// is a busy status.
type ServicePresenceStatusChannels struct {
	Channel []TextLiteral `xml:"channel"`
}

func (s *ServicePresenceStatus) SetMetadata(m metadata.MetadataInfo) {
	s.MetadataInfo = m
}

func (s *ServicePresenceStatus) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ServicePresenceStatus, error) {
	s := &ServicePresenceStatus{}
	return s, metadata.ParseMetadataXml(s, path)
}
