package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestReverse(t *testing.T) {
	s := String("2014-06-08")
	splitedValues := s.Split(Dash.String())
	var reversed String
	for i := len(splitedValues) - 1; i >= 0; i-- {
		n := splitedValues[i]
		if i == 0 {
			reversed = reversed.Concat(String(n).Reverse().Reverse())
			break
		}
		reversed = reversed.Concat(String(n).Reverse().Reverse(), Dash)
	}
	assert.Equal(t, reversed.String(), "08-06-2014")
	s = String("Hello World")
	assert.Equal(t, s.ReverseFn(func(i int) String {
		if Space.Equal(string(s[i])) {
			return Dash
		}
		return String(string(s[i]))
	}), "dlroW-olleH")
}
