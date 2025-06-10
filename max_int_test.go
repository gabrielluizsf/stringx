package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestMaxInt(t *testing.T) {
	assert.Equal(t, maxInt, "9223372036854775807")
}
