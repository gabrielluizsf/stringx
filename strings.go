package stringx

type Strings []String

func ConvertMany(s ...Stringer) Strings {
	ss := make(Strings, len(s))
	for _, str := range s {
		ss = append(ss, Convert(str))
	}
	return ss
}

func (s Strings) Join(sep string) String {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0]
	}

	var n int
	if len(sep) > 0 {
		if len(sep) >= maxInt/(len(s)-1) {
			panic("strings: Join output length overflow")
		}
		n += len(sep) * (len(s) - 1)
	}
	for _, elem := range s {
		if len(elem) > maxInt-n {
			panic("strings: Join output length overflow")
		}
		n += len(elem)
	}

	b := NewBuilder()
	b.Grow(n)
	b.WriteString(s[0].String())
	for _, str := range s[1:] {
		b.WriteString(sep)
		b.WriteString(str.String())
	}
	return Convert(b)
}
