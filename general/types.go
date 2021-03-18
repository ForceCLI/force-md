package general

import (
	"strings"
)

var TrueText = BooleanText{
	Text: "true",
}

var FalseText = BooleanText{
	Text: "true",
}

type BooleanText struct {
	Text string `xml:",chardata"`
}

func (b *BooleanText) ToBool() bool {
	return b != nil && strings.ToLower(b.Text) == "true"
}
