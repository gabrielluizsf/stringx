package stringx

// HasPrefix checks if the string starts with the specified prefix.
func (s String) HasPrefix(prefix string) bool {
	return s.Length() >= len(prefix) && s[:len(prefix)].Equal(prefix)
}
