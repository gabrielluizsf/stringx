package stringx

// Slice returns a substring of the String from start to end indices.
func (s String) Slice(start, end int) String {
	if start < 0 || end < 0 || start > end || end > len(s.String()) {
		return fail("stringx: invalid Slice indices", s, Log)
	}
	return String(s.String()[start:end])
}
