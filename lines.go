package stringx

import "iter"

// Lines returns an iterator over the newline-terminated lines in the string s.
func (s String) Lines() iter.Seq[string] {
	return func(yield func(string) bool) {
		for len(s) > 0 {
			var line string
			if i := s.IndexOf(NewLine.String()); i >= 0 {
				index := i + 1
				line, s = s.String()[:index], s[index:]
			} else {
				line, s = s.String(), Empty
			}
			if !yield(line) {
				return
			}
		}
	}
}
