package sharingrules

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
)

func (o CriteriaRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}

func (o OwnerRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.MarshalXml(o, e, start)
}
