package stringx

// HasSuffix reports whether s ends with suffix.
func (s String) HasSuffix(suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):].Equal(suffix)
}
