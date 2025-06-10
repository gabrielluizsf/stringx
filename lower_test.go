package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestLowerCase(t *testing.T) {
	str := String("HELLO WORLD")
	assert.Equal(t, str.ToLowerCase().String(), "hello world")
}
