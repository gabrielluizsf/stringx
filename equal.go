package stringx

// Equal checks if the string is equal to the given value.
func (s String) Equal(value string) bool {
	return IsEqual(s.String(), value)
}

// IsEqual checks if two strings are equal.
func IsEqual(s, value string) bool {
	return s == value
}