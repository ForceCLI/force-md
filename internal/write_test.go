package internal_test

import (
	"testing"

	. "github.com/ForceCLI/force-md/internal"
)

func TestSelfClosing(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte(`<x></x>`),
			expected: []byte(`<x/>`),
		},
		{
			input:    []byte(`<x attr="blah"></x>`),
			expected: []byte(`<x attr="blah"/>`),
		},
		{
			input:    []byte(`<x>contents</x>`),
			expected: []byte(`<x>contents</x>`),
		},
	}

	for _, test := range tests {
		result := SelfClosing(test.input)
		if string(result) != string(test.expected) {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, result)
		}
	}
}
