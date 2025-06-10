package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestRegex(t *testing.T) {
	s := String("Hello World")
	assert.True(t, s.Includes("World"))
	assert.False(t, s.Includes("!"))
	assert.True(t, s.Match(`\w+`))
	assert.Equal(t, s.Search(`\w+`), []string{"Hello", "World"})
	assert.Equal(t, s.IndexOf("World"), 6)
	assert.Equal(t, s.Replace("World", "Universe"), "Hello Universe")
	assert.Equal(t, s.Count("o"), 2)
}
