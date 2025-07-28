package stringx

// String is a functional string type.
type String string

// New creates a new String from any type that implements Stringer or is a string.
func New(s any) String {
	if v, ok := s.(Stringer); ok {
		return Convert(v)
	}
	if v, ok := s.(string); ok {
		return String(v)
	}
	return fail("stringx: cannot convert to String", Empty, Log)
}

// String returns the string representation of the String.
func (s String) String() string {
	return string(s)
}

// Concat concatenates the String with other Stringers.
func (s String) Concat(str ...Stringer) String {
	return s + ConvertMany(str...).Join(Empty.String())
}

// ConcatStrings concatenates the String with other strings.
func (s String) ConcatStrings(str ...string) String {
	return s + String(ConvertStrings(str...).Join(Empty.String()))
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
