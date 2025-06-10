package stringx

import (
	"testing"
	"unicode/utf8"

	"github.com/i9si-sistemas/assert"
)

func TestBuilder(t *testing.T) {
	t.Run("NewBuilder with no args", func(t *testing.T) {
		b := NewBuilder()
		if b.Len() != 0 {
			t.Errorf("expected Len 0, got %d", b.Len())
		}
		if b.Cap() < 0 {
			t.Errorf("expected non-negative Cap, got %d", b.Cap())
		}
	})

	t.Run("NewBuilder with initial buffer", func(t *testing.T) {
		initial := []byte("init")
		b2 := NewBuilder(initial...)
		if got := b2.String(); got != "init" {
			t.Errorf("expected String 'init', got '%s'", got)
		}
		if b2.Len() != len(initial) {
			t.Errorf("expected Len %d, got %d", len(initial), b2.Len())
		}
	})
	exportedBuilder := NewBuilder()
	t.Run("Write", func(t *testing.T) {
		b := exportedBuilder
		data := []byte("hello")
		n, err := b.Write(data)
		if err != nil || n != len(data) {
			t.Errorf("Write failed: n=%d, err=%v", n, err)
		}
		if b.String() != "hello" {
			t.Errorf("expected String 'hello', got '%s'", b.String())
		}

	})
	t.Run("WriteByte", func(t *testing.T) {
		b := exportedBuilder
		err := b.WriteByte('!')
		if err != nil {
			t.Errorf("WriteByte failed: %v", err)
		}
		if b.String() != "hello!" {
			t.Errorf("expected String 'hello!', got '%s'", b.String())
		}
	})
	t.Run("WriteRune", func(t *testing.T) {
		b := exportedBuilder
		r := '世'
		n, err := b.WriteRune(r)
		if err != nil || utf8.RuneCountInString(string(r)) != 1 {
			t.Errorf("WriteRune failed: n=%d, err=%v", n, err)
		}
		if got := b.String(); got != "hello!世" {
			t.Errorf("expected String 'hello!世', got '%s'", got)
		}
	})
	t.Run("WriteString", func(t *testing.T) {
		b := exportedBuilder
		n, err := b.WriteString("界")
		if err != nil || n != len("界") {
			t.Errorf("WriteString failed: n=%d, err=%v", n, err)
		}
		if got := b.String(); got != "hello!世界" {
			t.Errorf("expected String 'hello!世界', got '%s'", got)
		}
	})
	t.Run("Grow", func(t *testing.T) {
		b := exportedBuilder
		initialCap := b.Cap()
		b.Grow(10)
		if b.Cap() < initialCap {
			t.Errorf("Grow should not decrease capacity")
		}
	})
	assert.NotNil(t, Empty.Builder())
}
