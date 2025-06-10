package stringx

import (
	"unicode"
)

// ToLowerCase converts the string to lower case.
// It uses the unicode.ToLower function to ensure proper handling of Unicode characters.
func (s String) ToLowerCase() String {
	return changeCase(
		s,
		Lower,
		unicode.ToLower,
	)
}
