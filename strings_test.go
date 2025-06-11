package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestStrings(t *testing.T) {
	t.Run("ConvertMany with Stringers", func(t *testing.T) {
		assert.Equal(
			t,
			ConvertMany([]Stringer{String("hello"), String("world")}...),
			Strings{String("hello"), String("world")},
		)
	})
	t.Run("ConvertStrings with string", func(t *testing.T) {
		assert.Equal(
			t,
			ConvertStrings("hello", "world"),
			Strings{String("hello"), String("world")},
		)
	})
}
