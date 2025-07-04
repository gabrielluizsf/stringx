package stringx

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestParser(t *testing.T) {
	mustBool := func(s string) bool {
		v, err := NewParser(s).Bool()
		assert.NoError(t, err)
		return v
	}
	mustInt := func(s string) int64 {
		v, err := NewParser(s).Int()
		assert.NoError(t, err)
		return v
	}
	mustFloat := func(s string) float64 {
		v, err := NewParser(s).Float()
		assert.NoError(t, err)
		return v
	}

	t.Run("Booleans", func(t *testing.T) {
		assert.True(t, mustBool(" '' == ''"))
		assert.True(t, mustBool(` "" == "" `))
		assert.True(t, mustBool(" '10' == 10"))
		assert.True(t, mustBool("11 != 10"))
		assert.False(t, mustBool("'10' != 10"))
		assert.True(t, mustBool("10 ==  10"))
		assert.True(t, mustBool("true"))
		assert.False(t, mustBool("false"))
		assert.False(t, mustBool("false == true"))
		assert.True(t, mustBool("true == true"))
		assert.False(t, mustBool("'10'> 11"))
		assert.True(t, mustBool("'a' < 'b'"))
		assert.True(t, mustBool("10 > 9"))
		assert.False(t, mustBool("'9' > 10"))
	})
	t.Run("Integers", func(t *testing.T) {
		assert.Equal(t, mustInt("10"), 10)
		assert.Equal(t, mustInt("   10 + 10"), 20)
		assert.Equal(t, mustInt("10 * 20"), 200)
		assert.Equal(t, mustInt("10 - 20"), -10)
		assert.Equal(t, mustInt("10 / 20    "), 0)
		assert.Equal(t, mustInt("10 % 20"), 10)
	})

	t.Run("Floats", func(t *testing.T) {
		assert.Equal(t, mustFloat("10.0"), 10.0)
		assert.Equal(t, mustFloat("10 + 10.0"), 20.0)
		assert.Equal(t, mustFloat("10.0 * 20.0"), 200.0)
		assert.Equal(t, mustFloat("10.0 - 20.0"), -10.0)
		assert.Equal(t, mustFloat("10.0 / 20"), 0.5)
		assert.Equal(t, mustFloat("10 % 20"), 10)
	})
}
