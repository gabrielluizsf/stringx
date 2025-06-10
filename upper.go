package stringx

import (
	"unicode"
)

// ToUpperCase converts the string to upper case.
// It uses the unicode.ToUpper function to ensure proper handling of Unicode characters.
func (s String) ToUpperCase() String {
	return changeCase(
		s,
		Upper,
		unicode.ToUpper,
	)
}
