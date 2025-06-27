package stringx

import (
	"fmt"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestTrim(t *testing.T) {
	var trimTests = []struct {
		f            string
		in, arg, out string
	}{
		{"Trim", "abba", "a", "bb"},
		{"Trim", "abba", "ab", Empty.String()},
		{"TrimLeft", "abba", "ab", Empty.String()},
		{"TrimRight", "abba", "ab", Empty.String()},
		{"TrimLeft", "abba", "a", "bba"},
		{"TrimLeft", "abba", "b", "abba"},
		{"TrimRight", "abba", "a", "abb"},
		{"TrimRight", "abba", "b", "abba"},
		{"Trim", "<tag>", "<>", "tag"},
		{"Trim", "* listitem", " *", "listitem"},
		{"Trim", `"quote"`, `"`, "quote"},
		{"Trim", "\u2C6F\u2C6F\u0250\u0250\u2C6F\u2C6F", "\u2C6F", "\u0250\u0250"},
		{"Trim", "\x80test\xff", "\xff", "test"},
		{"Trim", " Ġ ", " ", "Ġ"},
		{"Trim", " Ġİ0", "0 ", "Ġİ"},
		{"TrimPrefix", "aabb", "a", "abb"},
		{"TrimPrefix", "aabb", "b", "aabb"},
		{"TrimSuffix", "aabb", "a", "aabb"},
		{"TrimSuffix", "aabb", "b", "aab"},
		//empty string tests
		{"Trim", "abba", Empty.String(), "abba"},
		{"Trim", Empty.String(), "123", Empty.String()},
		{"Trim", Empty.String(), Empty.String(), Empty.String()},
		{"TrimLeft", "abba", Empty.String(), "abba"},
		{"TrimLeft", Empty.String(), "123", Empty.String()},
		{"TrimLeft", Empty.String(), Empty.String(), Empty.String()},
		{"TrimRight", "abba", Empty.String(), "abba"},
		{"TrimRight", Empty.String(), "123", Empty.String()},
		{"TrimRight", Empty.String(), Empty.String(), Empty.String()},
		{"TrimRight", "☺\xc0", "☺", "☺\xc0"},
	}
	for i, tc := range trimTests {
		name := tc.f
		var f func(string, string) string
		switch name {
		case "Trim":
			f = func(s1, s2 string) string {
				return String(s1).Trim(s2).String()
			}
		case "TrimLeft":
			f = func(s1, s2 string) string {
				return String(s1).TrimStart(s2).String()
			}
		case "TrimRight":
			f = func(s1, s2 string) string {
				return String(s1).TrimEnd(s2).String()
			}
		case "TrimPrefix":
			f = func(s1, s2 string) string {
				return String(s1).TrimPrefix(s2).String()
			}
		case "TrimSuffix":
			f = func(s1, s2 string) string {
				return String(s1).TrimSuffix(s2).String()
			}
		default:
			t.Errorf("Undefined trim function %s", name)
		}
		actual := f(tc.in, tc.arg)
		assert.Equal(t, actual, tc.out, fmt.Sprintf("test case %d: %s(%q, %q)", i+1, name, tc.in, tc.arg))
	}
}
