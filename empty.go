package stringx

// Empty is an empty string.
var Empty = String("")

// IsEmpty returns true if the string is empty.
func (s String) IsEmpty() bool {
	return IsEmpty(s.String())
}

// IsEmpty checks if the given string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}
