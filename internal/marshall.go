package internal

import (
	"encoding/xml"
	"reflect"
)

var MarshalWide = false

func MarshalXml(o any, e *xml.Encoder, start xml.StartElement) error {
	var err error
	val := reflect.ValueOf(o)
	lt := reflect.TypeOf(o)
	if reflect.Zero(lt) == val {
		return nil
	}
	e.EncodeToken(start)
	if MarshalWide {
		DisableIndent(e)
	}

	for _, field := range reflect.VisibleFields(lt) {
		if alias, ok := field.Tag.Lookup("xml"); ok {
			if err = e.EncodeElement(val.FieldByName(field.Name).Interface(), xml.StartElement{Name: xml.Name{Local: alias}}); err != nil {
				return err
			}
		}
	}

	if MarshalWide {
		EnableIndent(e)
	}
	e.EncodeToken(start.End())
	e.Flush()
	return err
}
