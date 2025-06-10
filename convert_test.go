package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestConvert(t *testing.T) {
	sb := NewBuilder()
	msg := "Hello"
	sb.WriteString(msg)
	s := Convert(sb)
	assert.Equal(t, s.String(), msg)
}
