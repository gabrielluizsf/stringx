package stringx

import "iter"

// Lines returns an iterator over the newline-terminated lines in the string s.
func (s String) Lines() iter.Seq[string] {
	return func(yield func(string) bool) {
		for len(s) > 0 {
			var line string
			if i := s.IndexOf(NewLine.String()); i >= 0 {
				line, s = s.String()[:i+1], s[i+1:]
			} else {
				line, s = s.String(), ""
			}
			if !yield(line) {
				return
			}
		}
	}
}
