package general

import (
	"encoding/xml"
	"testing"

	nbioxml "github.com/nbio/xml"
)

func TestTextLiteralReadsCharacterData(t *testing.T) {
	cases := []struct {
		name string
		xml  string
		want string
	}{
		{"entity references", `<e>a &lt; b &amp;&amp; c &apos;d&apos; &#65;</e>`, `a < b && c 'd' A`},
		{"cdata section", "<e><![CDATA[AND(\n    a < 1,\n    b >= 2\n)]]></e>", "AND(\n    a < 1,\n    b >= 2\n)"},
		{"references inside cdata are literal", `<e><![CDATA[x &amp; y]]></e>`, `x &amp; y`},
		{"text around cdata", `<e>a &amp; <![CDATA[<b>]]> c</e>`, `a & <b> c`},
		{"comment", `<e>a<!-- note -->b</e>`, `ab`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var lit TextLiteral
			if err := nbioxml.Unmarshal([]byte(tc.xml), &lit); err != nil {
				t.Fatal(err)
			}
			if got := lit.String(); got != tc.want {
				t.Errorf("String() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestTextLiteralRejectsMalformedXML(t *testing.T) {
	var lit TextLiteral
	if err := nbioxml.Unmarshal([]byte(`<e>a <b</e>`), &lit); err == nil {
		t.Error("Unmarshal accepted malformed XML")
	}
}

type textLiteralDocument struct {
	XMLName xml.Name     `xml:"doc"`
	Lit     TextLiteral  `xml:"lit"`
	Ptr     *TextLiteral `xml:"ptr"`
	Missing *TextLiteral `xml:"missing"`
}

func TestTextLiteralWritesInnerXMLAsRead(t *testing.T) {
	in := `<doc><lit><![CDATA[a < b]]></lit><ptr>c &amp; d</ptr></doc>`
	var doc textLiteralDocument
	if err := nbioxml.Unmarshal([]byte(in), &doc); err != nil {
		t.Fatal(err)
	}
	if doc.Missing != nil {
		t.Errorf("Missing = %#v, want nil", doc.Missing)
	}
	out, err := xml.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != in {
		t.Errorf("Marshal = %s, want %s", out, in)
	}
}

func TestTextLiteralSetEscapesCharacterData(t *testing.T) {
	doc := textLiteralDocument{Ptr: NewTextLiteral(`R&D <"x"> 'y'`)}
	doc.Lit.Set("a < b")
	out, err := xml.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	want := `<doc><lit>a &lt; b</lit><ptr>R&amp;D &lt;&quot;x&quot;&gt; &apos;y&apos;</ptr></doc>`
	if string(out) != want {
		t.Errorf("Marshal = %s, want %s", out, want)
	}
	if got := doc.Ptr.String(); got != `R&D <"x"> 'y'` {
		t.Errorf("String() = %q", got)
	}
}

func TestTextLiteralReformatReplacesCharacterData(t *testing.T) {
	var doc textLiteralDocument
	if err := nbioxml.Unmarshal([]byte(`<doc><lit><![CDATA[a<b]]></lit></doc>`), &doc); err != nil {
		t.Fatal(err)
	}
	if err := doc.Lit.Reformat(func(s string) (string, error) { return s + " && c", nil }); err != nil {
		t.Fatal(err)
	}
	if got := doc.Lit.String(); got != "a<b && c" {
		t.Errorf("String() = %q", got)
	}
	out, err := xml.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `<doc><lit>a&lt;b &amp;&amp; c</lit></doc>` {
		t.Errorf("Marshal = %s", out)
	}
}

func TestNillableTextAndTypedTextKeepAttributes(t *testing.T) {
	in := `<doc xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">` +
		`<nil xsi:nil="true"></nil>` +
		`<typed xsi:type="xsd:string">a &amp; b</typed>` +
		`</doc>`
	type document struct {
		XMLName xml.Name     `xml:"doc"`
		Xsi     string       `xml:"xmlns:xsi,attr"`
		Nil     NillableText `xml:"nil"`
		Typed   TypedText    `xml:"typed"`
	}
	var doc document
	if err := nbioxml.Unmarshal([]byte(in), &doc); err != nil {
		t.Fatal(err)
	}
	if doc.Nil.Nil != "true" || doc.Nil.String() != "" {
		t.Errorf("Nil = %#v", doc.Nil)
	}
	if doc.Typed.Type != "xsd:string" || doc.Typed.String() != "a & b" {
		t.Errorf("Typed = %#v", doc.Typed)
	}
	out, err := xml.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != in {
		t.Errorf("Marshal = %s, want %s", out, in)
	}
}
