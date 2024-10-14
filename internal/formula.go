package internal

import "strings"

var FormulaEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&apos;",
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&quot;",
)
