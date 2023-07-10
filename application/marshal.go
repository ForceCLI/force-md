package application

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

func (o ActionOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}
func (o ProfileActionOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}
