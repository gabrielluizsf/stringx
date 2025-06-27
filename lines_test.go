package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestLines(t *testing.T) {
	s := String("hello\nworld\n")
	index := 0
	for line := range s.Lines() {
		switch index {
		case 0:
			assert.Equal(t, line, s[index:index+6])
		case 1:
			assert.Equal(t, line, s[index+5:])
		}
		index++
	}
}
