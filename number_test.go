package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestParseNumber(t *testing.T) {
	mult := String("   10 * 10   ")
	assert.Equal(t, ParseNumber(mult.String()).Float(), 100.0)
	assert.Equal(t, ParseNumber(mult.Concat(Space).String()).Float32(), 100.0)
	assert.Equal(t, ParseNumber(mult.String()).Int(), 100)
	assert.Equal(t, ParseNumber(mult.String()).Int32(), 100)
	assert.Equal(t, ParseNumber(mult.String()).Int16(), 100)
	assert.Equal(t, ParseNumber(mult.String()).Int8(), 100)
	sub := String("10         - 20   ")
	assert.Equal(t, ParseNumber(sub.Concat(Space).String()).Uint(), uint(18446744073709551606))
	assert.Equal(t, ParseNumber(sub.String()).Uint64(), uint(18446744073709551606))
	assert.Equal(t, ParseNumber(sub.String()).Uint32(), uint(4294967286))
	assert.Equal(t, ParseNumber(sub.String()).Uint16(), uint(65526))
	assert.Equal(t, ParseNumber(sub.String()).Uint8(), uint(246))
}
