package profile

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

func (o ApplicationVisibility) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o LayoutAssignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o RecordTypeVisibility) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o TabVisibility) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}
