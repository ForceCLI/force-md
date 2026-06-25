package metadata

import "testing"

func TestEscapeBareAmpersands(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"no ampersand", `<a>plain</a>`, `<a>plain</a>`},
		{"bare in text", `<a>a & b</a>`, `<a>a &amp; b</a>`},
		{"bare no spaces", `<a>R&D</a>`, `<a>R&amp;D</a>`},
		{"already escaped amp", `<a>a &amp; b</a>`, `<a>a &amp; b</a>`},
		{"predefined entities untouched", `<a>&lt;&gt;&quot;&apos;</a>`, `<a>&lt;&gt;&quot;&apos;</a>`},
		{"decimal char ref untouched", `<a>&#169;</a>`, `<a>&#169;</a>`},
		{"hex char ref untouched", `<a>&#x1F600;</a>`, `<a>&#x1F600;</a>`},
		{"unknown named entity escaped", `<a>&copy;</a>`, `<a>&amp;copy;</a>`},
		{"trailing bare ampersand", `<a>x&</a>`, `<a>x&amp;</a>`},
		{"bare in attribute", `<a b="x & y"/>`, `<a b="x &amp; y"/>`},
		{"ampersand in CDATA preserved", `<a><![CDATA[x & y]]></a>`, `<a><![CDATA[x & y]]></a>`},
		{"ampersand in comment preserved", `<a><!-- x & y --></a>`, `<a><!-- x & y --></a>`},
		{"mixed", `<a>a & b &amp; c &#38; d</a>`, `<a>a &amp; b &amp; c &#38; d</a>`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := string(escapeBareAmpersands([]byte(tc.in)))
			if got != tc.want {
				t.Errorf("escapeBareAmpersands(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
