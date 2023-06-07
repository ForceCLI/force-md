package custommetadata

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (p *CustomMetadata) UpdateFieldValue(field string, value string) error {
	found := false
	for i, u := range p.Values {
		if strings.ToLower(u.Field) == strings.ToLower(field) {
			found = true
			p.Values[i].Value.Text = value
		}
	}
	if !found {
		return errors.New("field not found")
	}
	return nil
}

func (m *CustomMetadata) AddValue(key string, value any) {
	var valueType, stringValue string
	switch t := value.(type) {
	case bool:
		valueType = "xsd:boolean"
		stringValue = strconv.FormatBool(t)
	case float32, float64:
		valueType = "xsd:double"
		stringValue = value.(string)
	case int:
		valueType = "xsd:int"
		stringValue = strconv.Itoa(t)
	default:
		valueType = "xsd:string"
		stringValue = fmt.Sprintf("%s", t)
	}
	if stringValue == "" {
		m.Values = append(m.Values, Value{Field: key, Value: TypedValue{Nil: "true"}})
	} else {
		m.Values = append(m.Values, Value{Field: key, Value: TypedValue{Text: stringValue, Type: valueType}})
	}
}
