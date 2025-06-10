package stringx

// Case represents the case of a string.
type Case int

const (
	// Upper represents the upper case.
	Upper Case = iota
	// Lower represents the lower case.
	Lower
)

func changeCase(
	str String, 
	c Case, 
	fn func(rune) rune,
) String {
	s := str.String()

	unchanged := true
	for _, r := range s {
		switch c {
		case Upper:
			if 'a' <= r && r <= 'z' {
				unchanged = false
			}
		case Lower:
			if 'A' <= r && r <= 'Z' {
				unchanged = false
			}
		}
	}

	if unchanged {
		return str
	}

	return str.Map(fn)
}
