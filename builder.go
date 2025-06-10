package stringx

import (
	"unicode/utf8"
)


// Builder returns a new stringx.Builder.
func (s String) Builder() *Builder {
	b := NewBuilder([]byte(s.String())...)
	b.addr = b
	return b
}

// A Builder is used to efficiently build a string using Write methods.
type Builder struct {
	addr *Builder 

	buf []byte
}

// NewBuilder returns a new Builder.
func NewBuilder(buf ...byte) *Builder {
	if len(buf) == 0 {
		return &Builder{}
	}
	b := &Builder{
		addr: new(Builder),
		buf:  buf,
	}
	return b
}

// String returns the accumulated string.
func (b *Builder) String() string {
	return string(b.buf)
}

// Len returns the number of accumulated bytes
func (b *Builder) Len() int { return len(b.buf) }

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder) Cap() int { return cap(b.buf) }

// grow copies the buffer to a new, larger buffer so that there are at least n
// bytes of capacity beyond len(b.buf).
func (b *Builder) grow() {
	buf := b.buf[:len(b.buf)]
	copy(buf, b.buf)
	b.buf = buf
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (b *Builder) Grow(n int) {
	if n < 0 {
		panic("stringx.Builder.Grow: negative count")
	}
	if cap(b.buf)-len(b.buf) < n {
		b.grow()
	}
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (b *Builder) WriteByte(c byte) error {
	b.buf = append(b.buf, c)
	return nil
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
func (b *Builder) WriteRune(r rune) (int, error) {
	n := len(b.buf)
	b.buf = utf8.AppendRune(b.buf, r)
	return len(b.buf) - n, nil
}

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *Builder) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}
