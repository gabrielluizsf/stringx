package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestRandom(t *testing.T) {
	emojis := []string{"😊", "😅", "🤖", "🙃", "⌛", "🧠", "💡"}
	r := NewRandomString(emojis...)
	s := r.Random()
	t.Log(s)
	r = random{
		s: ConvertStrings(emojis...),
		fn: func(len int) int {
			return 1
		},
	}
	s = r.Random()
	assert.Equal(t, s, "😅")
}
