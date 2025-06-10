package stringx

// Empty is an empty string.
var Empty = String("")

// IsEmpty returns true if the string is empty.
func (s String) IsEmpty() bool {
	return s == Empty
}
