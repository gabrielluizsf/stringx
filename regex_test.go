package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestRegex(t *testing.T) {
	s := String("Hello World")
	assert.True(t, s.Match(`\w+`))
	assert.Equal(t, s.Search(`\w+`), []string{"Hello", "World"})
}
