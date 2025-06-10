package stringx

// Slice returns a substring of the String from start to end indices.
// If the indices are out of bounds, it will panic.
func (s String) Slice(start, end int) String {
	return String(s.String()[start:end])
}
