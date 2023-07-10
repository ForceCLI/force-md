package permissionset

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

func (o FieldPermissions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o ObjectPermissions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o PageAccess) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o UserPermission) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}
