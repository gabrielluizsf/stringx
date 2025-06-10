package stringx

import (
	"unicode/utf8"
)

func (s String) Map(fn func(r rune) rune) String {
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
