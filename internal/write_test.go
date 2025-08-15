package internal_test

import (
	"testing"

	. "github.com/ForceCLI/force-md/internal"
)

func TestSelfClosing(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "simple empty tag",
			input:    []byte(`<x></x>`),
			expected: []byte(`<x/>`),
		},
		{
			name:     "empty tag with attributes",
			input:    []byte(`<x attr="blah"></x>`),
			expected: []byte(`<x attr="blah"/>`),
		},
		{
			name:     "non-empty tag should not change",
			input:    []byte(`<x>contents</x>`),
			expected: []byte(`<x>contents</x>`),
		},
		{
			name:     "empty tag with whitespace",
			input:    []byte(`<tag>  </tag>`),
			expected: []byte(`<tag/>`),
		},
		{
			name:     "empty tag with newline",
			input:    []byte("<tag>\n</tag>"),
			expected: []byte(`<tag/>`),
		},
		{
			name:     "tag with namespace",
			input:    []byte(`<ns:tag></ns:tag>`),
			expected: []byte(`<ns:tag/>`),
		},
		{
			name:     "tag with namespace and attributes",
			input:    []byte(`<ns:tag attr="value"></ns:tag>`),
			expected: []byte(`<ns:tag attr="value"/>`),
		},
		{
			name:     "tag with dots in name",
			input:    []byte(`<tag.name></tag.name>`),
			expected: []byte(`<tag.name/>`),
		},
		{
			name:     "tag with dashes in name",
			input:    []byte(`<tag-name></tag-name>`),
			expected: []byte(`<tag-name/>`),
		},
		{
			name:     "nested empty tags",
			input:    []byte(`<outer><inner></inner></outer>`),
			expected: []byte(`<outer><inner/></outer>`),
		},
		{
			name:     "multiple empty tags",
			input:    []byte(`<tag1></tag1><tag2></tag2>`),
			expected: []byte(`<tag1/><tag2/>`),
		},
		{
			name:     "real Salesforce example - description",
			input:    []byte(`<description></description>`),
			expected: []byte(`<description/>`),
		},
		{
			name:     "real Salesforce example - masterLabel",
			input:    []byte(`<masterLabel></masterLabel>`),
			expected: []byte(`<masterLabel/>`),
		},
		{
			name:     "real Salesforce example - groupingSortProperties",
			input:    []byte(`<groupingSortProperties></groupingSortProperties>`),
			expected: []byte(`<groupingSortProperties/>`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SelfClosing(test.input)
			if string(result) != string(test.expected) {
				t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, result)
			}
		})
	}
}
