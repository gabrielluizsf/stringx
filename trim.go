package stringx

import (
	"unicode/utf8"
)

// Trim removes all leading and trailing characters in cutset from the string.
func (s String) Trim(cutset string) (result String) {
	byteFn := func() String {
		return String(
			trimLeftByte(
				trimRightByte(s.String(), cutset[0]),
				cutset[0],
			),
		)
	}
	asciiFn := func(as asciiSet) String {
		return String(
			trimLeftASCII(
				trimRightASCII(s.String(), &as),
				&as,
			),
		)
	}
	resultFn := func() String {
		return String(
			trimLeftUnicode(
				trimRightUnicode(s.String(), cutset),
				cutset,
			),
		)
	}
	result = trim(
		s, cutset,
		byteFn,
		asciiFn,
		resultFn,
	)
	return
}

// TrimStart removes all leading characters in cutset from the string.
func (s String) TrimStart(cutset string) (result String) {
	byteFn := func() String {
		return String(
			trimLeftByte(
				s.String(),
				cutset[0],
			),
		)
	}
	asciiFn := func(as asciiSet) String {
		return String(
			trimLeftASCII(
				s.String(),
				&as,
			),
		)
	}
	resultFn := func() String {
		return String(
			trimLeftUnicode(
				s.String(),
				cutset,
			),
		)
	}
	result = trim(
		s, cutset,
		byteFn,
		asciiFn,
		resultFn,
	)
	return
}

// TrimEnd removes all trailing characters in cutset from the string.
func (s String) TrimEnd(cutset string) (result String) {
	byteFn := func() String {
		return String(
			trimRightByte(
				s.String(),
				cutset[0],
			),
		)
	}
	asciiFn := func(as asciiSet) String {
		return String(
			trimRightASCII(
				s.String(),
				&as,
			),
		)
	}
	resultFn := func() String {
		return String(
			trimRightUnicode(
				s.String(),
				cutset,
			),
		)
	}
	result = trim(
		s, cutset,
		byteFn,
		asciiFn,
		resultFn,
	)

	return
}

type trimFn func() String

func trim(
	s String,
	cutset string,
	byteFn trimFn,
	asciiFn func(asciiSet) String,
	resultFn trimFn,
) String {
	if s == Empty || String(cutset) == Empty {
		return s
	}
	if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
		return byteFn()
	}
	if as, ok := makeASCIISet(cutset); ok {
		return asciiFn(as)
	}
	return resultFn()
}

func trimLeftByte(s string, c byte) string {
	for len(s) > 0 && s[0] == c {
		s = s[1:]
	}
	return s
}

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
// This allocates a total of 32 bytes even though the upper half
// is unused to avoid bounds checks in asciiSet.contains.
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := range len(chars) {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c/32] |= 1 << (c % 32)
	}
	return as, true
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
	return (as[c/32] & (1 << (c % 32))) != 0
}

func trimLeftASCII(s string, as *asciiSet) string {
	for len(s) > 0 {
		if !as.contains(s[0]) {
			break
		}
		s = s[1:]
	}
	return s
}

func trimLeftUnicode(s, cutset string) string {
	for len(s) > 0 {
		r, n := rune(s[0]), 1
		if r >= utf8.RuneSelf {
			r, n = utf8.DecodeRuneInString(s)
		}
		if !includes(cutset, r) {
			break
		}
		s = s[n:]
	}
	return s
}

func trimRightByte(s string, c byte) string {
	for len(s) > 0 && s[len(s)-1] == c {
		s = s[:len(s)-1]
	}
	return s
}

func trimRightASCII(s string, as *asciiSet) string {
	for len(s) > 0 {
		if !as.contains(s[len(s)-1]) {
			break
		}
		s = s[:len(s)-1]
	}
	return s
}

func trimRightUnicode(s, cutset string) string {
	for len(s) > 0 {
		r, n := rune(s[len(s)-1]), 1
		if r >= utf8.RuneSelf {
			r, n = utf8.DecodeLastRuneInString(s)
		}
		if !includes(cutset, r) {
			break
		}
		s = s[:len(s)-n]
	}
	return s
}

func includes(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
