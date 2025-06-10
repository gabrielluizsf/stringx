package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestRepeat(t *testing.T) {
	result := Empty.Repeat("x", 5)
	assert.Equal(t, result, String("xxxxx"))
}
