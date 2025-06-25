package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestReplacer(t *testing.T) {
	t.Run("Replace", func(t *testing.T) {
		replacer := NewReplacer(String("hello world"), []string{"hello", "world"}, []string{"hi", "universe"})
		replaced, err := replacer.Replace()
		if err != nil {
			t.Fatal(err)
		}
		assert.StrictEqual(t, replaced, "hi universe")
		assert.Equal(t, String("Hello World").Replace("World", "Universe"), "Hello Universe")
	})
}
