package stringx

import "math/bits"

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
	repeatedCount := 64 * 2

	var (
		repeatedSpaces = Empty.Repeat(" ", repeatedCount)
		repeatedDashes = Empty.Repeat("-", repeatedCount)
		repeatedZeroes = Empty.Repeat("0", repeatedCount)
		repeatedEquals = Empty.Repeat("=", repeatedCount)
		repeatedTabs   = Empty.Repeat("\t", repeatedCount)
	)
	switch value[0] {
	case ' ', '-', '0', '=', '\t':
		switch {
		case n <= repeatedSpaces.Length() && repeatedSpaces.HasPrefix(value):
			return repeatedSpaces[:n]
		case n <= repeatedDashes.Length() && repeatedDashes.HasPrefix(value):
			return repeatedDashes[:n]
		case n <= repeatedZeroes.Length() && repeatedZeroes.HasPrefix(value):
			return repeatedZeroes[:n]
		case n <= repeatedEquals.Length() && repeatedEquals.HasPrefix(value):
			return repeatedEquals[:n]
		case n <= repeatedTabs.Length() && repeatedTabs.HasPrefix(value):
			return repeatedTabs[:n]
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
