package internal

import "strings"

// Trim suffix plus anything that follows it.
func TrimSuffixToEnd(s, suffix string) string {
	return s[0:strings.LastIndex(s, suffix)]
}
