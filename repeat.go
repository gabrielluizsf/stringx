package stringx

import (
	"math/bits"
)

// Repeat returns a new String by repeating the specified `value` string
// `count` times and concatenating the result to the original String `s`.
//
// example:
//
//	String("Go").Repeat("lang", 3) // returns "Golanglanglang"
//
// Parameters:
//
//	value - the string to repeat
//	count - the number of times to repeat the value
//
// Returns:
//
//	A new String with the repeated value appended to the original.
func (s String) Repeat(value string, count int) String {
	return s.Concat(repeat(value, count))
}

var (
	// Space is the space character.
	Space = String(" ")

	// Dash is the dash character.
	Dash = String("-")

	// Zero is the zero character
	Zero = String("0")

	// Equals is the equals character
	Equals = String("=")

	// Tab is the tab character.
	Tab = String("\t")

	// Plus is the plus character.
	Plus = String("+")
)

func repeat(value string, count int) String {
	switch count {
	case 0:
		return Empty
	case 1:
		return String(value)
	}
	if count < 0 {
		panic("stringx: negative Repeat count")
	}
	concatLength := len(value) * count
	hi, lo := bits.Mul(uint(len(value)), uint(count))
	if hi > 0 || lo > uint(maxInt) {
		panic("stringx: Repeat output length overflow")
	}
	n := int(lo)

	if concatLength == 0 {
		return Empty
	}
	const chunkLimit = 8 * 1024
	chunkMax := n
	if n > chunkLimit {
		chunkMax = chunkLimit / len(value) * len(value)
		if chunkMax == 0 {
			chunkMax = len(value)
		}
	}

	b := NewBuilder()
	b.Grow(n)
	b.WriteString(value)
	for b.Len() < n {
		chunk := min(n-b.Len(), b.Len(), chunkMax)
		b.WriteString(b.String()[:chunk])
	}
	return Convert(b)
}
