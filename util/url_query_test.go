package util

import (
	"fmt"
	"net/url"
	"testing"
)

func TestEncodeQueryRoundTrip(t *testing.T) {
	cases := []struct {
		name  string
		value string
	}{
		{"empty", ""},
		{"plain", "skill-name"},
		{"space", "hello world"},
		{"slash", "docs/SKILL.md"},
		{"question", "what?"},
		{"equals", "a=b"},
		{"unicode", "技能说明"},
		{"plus", "C++"},
		{"ampersand", "cats&dogs"},
		{"fragment", "docs/part#1.md"},
		{"percent", "100%"},
		{"encoded-looking", "%2F"},
		{"mixed", "a+b & c#d%?/=e"},
	}
	// Search, FetchFile and Download use these query parameter names.
	for _, key := range []string{"q", "path", "slug"} {
		t.Run(key, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					u, err := url.Parse(fmt.Sprintf("https://example.com/api?%s=%s&limit=25", key, EncodeQuery(tc.value)))
					if err != nil {
						t.Fatal(err)
					}
					values, err := url.ParseQuery(u.RawQuery)
					if err != nil {
						t.Fatal(err)
					}
					if values.Get(key) != tc.value || values.Get("limit") != "25" || len(values) != 2 || u.Fragment != "" {
						t.Fatalf("query round trip: values=%v fragment=%q, want %s=%q and limit=25", values, u.Fragment, key, tc.value)
					}
				})
			}
		})
	}
}
