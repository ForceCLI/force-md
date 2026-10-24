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

func TestSelfClosingWithAttributes(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "empty tag stays open",
			input:    []byte(`<value></value>`),
			expected: []byte(`<value></value>`),
		},
		{
			name:     "empty tag with attributes",
			input:    []byte(`<blockId xsi:nil="true"></blockId>`),
			expected: []byte(`<blockId xsi:nil="true"/>`),
		},
		{
			name:     "non-empty tag with attributes should not change",
			input:    []byte(`<value xsi:type="xsd:string">x</value>`),
			expected: []byte(`<value xsi:type="xsd:string">x</value>`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SelfClosingWithAttributes(test.input)
			if string(result) != string(test.expected) {
				t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, result)
			}
		})
	}
}

func TestMarshalEmptyRoot(t *testing.T) {
	type matchingRules struct {
		XMLName struct{} `xml:"MatchingRules"`
		Xmlns   string   `xml:"xmlns,attr"`
		Rules   []string `xml:"matchingRules"`
	}
	out, err := Marshal(matchingRules{Xmlns: "http://soap.sforce.com/2006/04/metadata"})
	if err != nil {
		t.Fatal(err)
	}
	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MatchingRules xmlns=\"http://soap.sforce.com/2006/04/metadata\"/>\n\n"
	if string(out) != expected {
		t.Errorf("Expected: %q\nGot: %q", expected, out)
	}

	out, err = Marshal(matchingRules{Xmlns: "http://soap.sforce.com/2006/04/metadata", Rules: []string{"Rule"}})
	if err != nil {
		t.Fatal(err)
	}
	expected = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MatchingRules xmlns=\"http://soap.sforce.com/2006/04/metadata\">\n    <matchingRules>Rule</matchingRules>\n</MatchingRules>\n"
	if string(out) != expected {
		t.Errorf("Expected: %q\nGot: %q", expected, out)
	}
}

func TestSelfClosingNamed(t *testing.T) {
	input := []byte("<value>\n    </value><stringValue></stringValue><values></values>")
	expected := "<value/><stringValue></stringValue><values></values>"
	if got := string(SelfClosingNamed(input, []string{"value"})); got != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, got)
	}
	if got := string(SelfClosingNamed(input, nil)); got != string(input) {
		t.Errorf("Expected no change with no names, got: %s", got)
	}
}

func TestMarshalWritesNewlinesAndTabsLiterally(t *testing.T) {
	type field struct {
		XMLName  struct{} `xml:"CustomField"`
		HelpText string   `xml:"inlineHelpText"`
	}
	out, err := Marshal(field{HelpText: "First line\n\tSecond line"})
	if err != nil {
		t.Fatal(err)
	}
	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<CustomField>\n    <inlineHelpText>First line\n\tSecond line</inlineHelpText>\n</CustomField>\n"
	if string(out) != expected {
		t.Errorf("Expected: %q\nGot: %q", expected, out)
	}
}
