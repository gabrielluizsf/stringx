package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, New("Hello World").String(), "Hello World")
	t.Run("Concat", func(t *testing.T) {
		assert.Equal(t, New("Hello").Concat(
			Space,
			String("World"),
		), "Hello World")
	})
	t.Run("ConcatStrings", func(t *testing.T) {
		assert.Equal(t, New("Hello").ConcatStrings(
			Space.String(),
			"World",
		), "Hello World")
	})
	t.Run("Bytes", func(t *testing.T) {
		assert.Equal(t, New("Hello World").Bytes(), []byte("Hello World"))
	})
	t.Run("Runes", func(t *testing.T) {
		assert.Equal(t, New("Hello World").Runes(), []rune("Hello World"))
	})
	t.Run("CharAt", func(t *testing.T) {
		assert.Equal(t, New("Hello World").CharAt(0), String("H"))
	})
}
