package stringx

import "unicode/utf8"

// Split splits the string s into substrings separated by sep.
func (s String) Split(sep string) []string {
	return split(s, sep, 0, -1)
}

func split(s String, sep string, sepSave, n int) []string {
	if !s.Includes(sep) {
		return []string{s.String()}
	}
	if n == 0 {
		return []string{}
	}
	if sep == "" {
		return explode(s.String(), n)
	}
	if n < 0 {
		n = String(s).Count(sep) + 1
	}

	if n > len(s)+1 {
		n = len(s) + 1
	}
	a := make([]string, n)
	n--
	i := 0
	for i < n {
		m := String(s).IndexOf(sep)
		if m < 0 {
			break
		}
		a[i] = s.String()[:m+sepSave]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s.String()
	return a[:i+1]
}

// explode splits s into a slice of UTF-8 strings,
// one string per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 bytes are sliced individually.
func explode(s string, n int) []string {
	l := utf8.RuneCountInString(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := range n - 1 {
		_, size := utf8.DecodeRuneInString(s)
		a[i] = s[:size]
		s = s[size:]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}
