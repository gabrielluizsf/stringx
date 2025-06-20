package stringx

import (
	"math/bits"
	"strings"
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

	var (
		repeatedCount  = 64 * 2
		repeatedSpaces = strings.Repeat(" ", repeatedCount)
		repeatedDashes = strings.Repeat("-", repeatedCount)
		repeatedZeroes = strings.Repeat("0", repeatedCount)
		repeatedEquals = strings.Repeat("=", repeatedCount)
		repeatedTabs   = strings.Repeat("\t", repeatedCount)
	)
	switch value[0] {
	case ' ', '-', '0', '=', '\t':
		switch {
		case n <= String(repeatedSpaces).Length() && String(repeatedSpaces).HasPrefix(value):
			return String(repeatedSpaces)[:n]
		case n <= String(repeatedDashes).Length() && String(repeatedDashes).HasPrefix(value):
			return String(repeatedDashes)[:n]
		case n <= String(repeatedZeroes).Length() && String(repeatedZeroes).HasPrefix(value):
			return String(repeatedZeroes)[:n]
		case n <= String(repeatedEquals).Length() && String(repeatedEquals).HasPrefix(value):
			return String(repeatedEquals)[:n]
		case n <= String(repeatedTabs).Length() && String(repeatedTabs).HasPrefix(value):
			return String(repeatedTabs)[:n]
		}
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
