package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestFinder(t *testing.T) {
	s := String("Hello World")
	assert.Equal(t, s.IndexOf("World"), 6)
	assert.True(t, s.Includes("World"))
	assert.False(t, s.Includes("!"))
	assert.Equal(t, s.Count("o"), 2)
}
