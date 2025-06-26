package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestSplit(t *testing.T) {
	s := String("a, b, c")
	assert.Equal(t, s.Split(", "), []string{"a", "b", "c"})
	assert.Equal(t, s.Split("; "), []string{"a, b, c"})
	assert.Equal(t, len(s.Split(Empty.String())), 7)
	result := String("github.com/i9si/atendi9api/internal/server/handler/whatsapp.processLicensePlate").Split(".")
	expected := []string{"github", "com/i9si/atendi9api/internal/server/handler/whatsapp", "processLicensePlate"}
	assert.Equal(t, result, expected)
	assert.Equal(t, s.SplitN(", ", 2), []string{"a", "b, c"})
}
