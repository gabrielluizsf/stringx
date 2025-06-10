package stringx

// Stringer is an interface that can be converted to a string.
type Stringer interface {
	// String returns the string representation of the object.
	String() string
}

// Convert converts a Stringer to a String.
// 
// Example:
//		type Slice []any
//		func (s Slice) String() string {
//			return fmt.Sprintf("%v", s)
//		}
//		slice := Slice{1, 2, 3}
//		s := Convert(slice)
func Convert(s Stringer) String {
	return String(s.String())
}
