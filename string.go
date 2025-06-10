package stringx

// String is a functional string type.
type String string

// String returns the string representation of the String.
func (s String) String() string {
	return string(s)
}

// Concat concatenates the String with other Stringers.
func (s String) Concat(str ...Stringer) String {
	return s + ConvertMany(str...).Join(Empty.String())
}

// Bytes returns the byte slice representation of the String.
func (s String) Bytes() []byte {
	return []byte(s.String())
}

// Runes returns the rune slice representation of the String.
func (s String) Runes() []rune {
	return []rune(s.String())
}

// CharAt returns the character at the given index.
func (s String) CharAt(index int) String {
	return String(s.String()[index])
}
