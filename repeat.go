package stringx

import "math/bits"

// According to static analysis, spaces, dashes, zeros, equals, and tabs
// are the most commonly repeated string literal,
// often used for display on fixed-width terminal windows.
// Pre-declare constants for these for O(1) repetition in the common-case.
const (
	RepeatedSpaces = "" +
		"                                                                " +
		"                                                                "
	RepeatedDashes = "" +
		"----------------------------------------------------------------" +
		"----------------------------------------------------------------"
	RepeatedZeroes = "" +
		"0000000000000000000000000000000000000000000000000000000000000000"
	RepeatedEquals = "" +
		"================================================================" +
		"================================================================"
	RepeatedTabs = "" +
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t" +
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"
)

// Repeat returns a new String by repeating the specified `value` string 
// `count` times and concatenating the result to the original String `s`.
//
// example:
//   String("Go").Repeat("lang", 3) // returns "Golanglanglang"
//
// Parameters:
//   value - the string to repeat
//   count - the number of times to repeat the value
//
// Returns:
//   A new String with the repeated value appended to the original.
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

    switch value[0] {
    case ' ', '-', '0', '=', '\t':
        switch {
        case n <= len(RepeatedSpaces) && String(RepeatedSpaces).HasPrefix(value):
            return String(RepeatedSpaces[:n])
        case n <= len(RepeatedDashes) && String(RepeatedDashes).HasPrefix(value):
            return String(RepeatedDashes[:n])
        case n <= len(RepeatedZeroes) && String(RepeatedZeroes).HasPrefix(value):
            return String(RepeatedZeroes[:n])
        case n <= len(RepeatedEquals) && String(RepeatedEquals).HasPrefix(value):
            return String(RepeatedEquals[:n])
        case n <= len(RepeatedTabs) && String(RepeatedTabs).HasPrefix(value):
            return String(RepeatedTabs[:n])
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

