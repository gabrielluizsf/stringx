package stringx

func (s String) Reverse() (result String) {
	for i := len(s) - 1; i >= 0; i-- {
		result = result.Concat(String(s[i]))
	}
	return
}

type ReverseFn func(index int) String

func (s String) ReverseFn(fn ReverseFn) (result String) {
	for i := len(s) - 1; i >= 0; i-- {
		result = result.Concat(fn(i))
	}
	return
}
