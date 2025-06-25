package stringx

import "unicode/utf8"

// IndexOf returns the index of the first occurrence of substr in s, or -1 if substr is not present in s.
func (s String) IndexOf(substr string) int {
	return makeStringFinder(substr).next(s.String())
}

// Includes checks if the string s contains the substring substr.
func (s String) Includes(substr string) bool {
	return s.IndexOf(substr) != -1
}

// Count returns the number of non-overlapping instances of substr in s.
func (s String) Count(value string) int {
	if len(value) == 0 {
		return utf8.RuneCountInString(s.String()) + 1
	}
	n := 0
	for {
		i := s.IndexOf(value)
		if i == -1 {
			return n
		}
		n++
		s = s[i+len(value):]
	}
}

// stringFinder efficiently finds strings in a source text.
type stringFinder struct {
	pattern string
	bcs     [256]int
	gss     []int
}

// next returns the index in text of the first occurrence of the pattern. If
// the pattern is not found, it returns -1.
func (f *stringFinder) next(text string) int {
	i := len(f.pattern) - 1
	for i < len(text) {
		j := len(f.pattern) - 1
		for j >= 0 && text[i] == f.pattern[j] {
			i--
			j--
		}
		if j < 0 {
			return i + 1
		}
		i += max(f.bcs[text[i]], f.gss[j])
	}
	return -1
}

func makeStringFinder(pattern string) *stringFinder {
	f := &stringFinder{
		pattern: pattern,
		gss:     make([]int, len(pattern)),
	}
	last := len(pattern) - 1
	for i := range f.bcs {
		f.bcs[i] = len(pattern)
	}
	for i := range last {
		f.bcs[pattern[i]] = last - i
	}
	lastPrefix := last
	for i := last; i >= 0; i-- {
		if String(pattern).HasPrefix(pattern[i+1:]) {
			lastPrefix = i + 1
		}
		f.gss[i] = lastPrefix + last - i
	}
	for i := range last {
		lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
			f.gss[last-lenSuffix] = lenSuffix + last - i
		}
	}
	return f
}

func longestCommonSuffix(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}
