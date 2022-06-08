package general

import (
	"strings"
)

type TextLiteral struct {
	Text string `xml:",innerxml"`
}

var TrueText = BooleanText{
	Text: "true",
}

var FalseText = BooleanText{
	Text: "false",
}

type BooleanText struct {
	Text string `xml:",chardata"`
}

type IntegerText struct {
	Text string `xml:",chardata"`
}

func (b *BooleanText) ToBool() bool {
	return b != nil && strings.ToLower(b.Text) == "true"
}

func (b *BooleanText) IsTrue() bool {
	return b != nil && strings.ToLower(b.Text) == "true"
}

func (b *BooleanText) IsFalse() bool {
	return b != nil && strings.ToLower(b.Text) == "false"
}
