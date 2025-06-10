package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestMap(t *testing.T) {
	s := String("password")
	str := s.Map(func(r rune) rune {
		return 'x'
	}).String()
	expected := Empty.Repeat("x", s.Length())
	assert.Equal(t, str, expected)
}
