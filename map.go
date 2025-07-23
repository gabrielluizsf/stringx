package stringx

import (
	"unicode/utf8"
)

// MapFn defines a function type that takes a rune and returns a rune.
type MapFn func(r rune) rune

// Map applies a function to each rune in the string, returning a new string
func (s String) Map(fn MapFn) String {
	return Map(s, fn)
}

// Map applies a function to each rune in the string, returning a new string
func Map(s String, fn MapFn) String {
	b := NewBuilder()

	for i, c := range s {
		r := fn(c)
		if r == c && c != utf8.RuneError {
			continue
		}

		var width int
		if c == utf8.RuneError {
			c, width = utf8.DecodeRuneInString(s.String()[i:])
			if width != 1 && r == c {
				continue
			}
		} else {
			width = utf8.RuneLen(c)
		}

		b.Grow(len(s) + utf8.UTFMax)
		b.WriteString(s.String()[:i])
		if r >= 0 {
			b.WriteRune(r)
		}

		s = s[i+width:]
		break
	}

	if b.Cap() == 0 {
		return s
	}

	for _, c := range s {
		r := fn(c)
		if r >= 0 {
			if r < utf8.RuneSelf {
				b.WriteByte(byte(r))
				continue
			}
			b.WriteRune(r)
		}
	}

	return Convert(b)
}