package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestUpperCase(t *testing.T) {
	str := String("hello world")
	assert.Equal(t, str.ToUpperCase().String(), "HELLO WORLD")
}
