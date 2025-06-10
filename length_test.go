package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestLength(t *testing.T) {
	s := String("Hello")
	assert.Equal(t, s.Length(), len(s))
}
