package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestSplit(t *testing.T) {
	assert.Equal(t, String("a, b, c").Split(", "), []string{"a", "b", "c"})
	assert.Equal(t, String("a, b, c").Split("; "), []string{"a, b, c"})
	assert.Equal(t, len(String("a, b, c").Split(Empty.String())), 7)
	result := String("github.com/i9si/atendi9api/internal/server/handler/whatsapp.processLicensePlate").Split(".")
	expected := []string{"github", "com/i9si/atendi9api/internal/server/handler/whatsapp", "processLicensePlate"}
	assert.Equal(t, result, expected)
}
