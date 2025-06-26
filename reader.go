package stringx

import (
	"errors"
	"io"
	"unicode/utf8"
)

// NewReader returns a new [github.com/i9si-sistemas/stringx.Reader] reading from s.
// It is similar to [bytes.NewBufferString] but more efficient and non-writable.
func NewReader(s string) *Reader {
	return &Reader{s, 0, -1}
}

// A Reader implements the [io.Reader], [io.ReaderAt], [io.ByteReader], [io.ByteScanner],
// [io.RuneReader], [io.RuneScanner], [io.Seeker], and [io.WriterTo] interfaces by reading
// from a string.
// The zero value for Reader operates like a Reader of an empty string.
type Reader struct {
	s string
	i int64
	p int
}

// Len returns the number of bytes of the unread portion of the
// string.
func (r *Reader) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

// Size returns the original length of the underlying string.
// Size is the number of bytes available for reading via [github.com/i9si-sistemas/stringx.Reader.ReadAt].
// The returned value is always the same and is not affected by calls
// to any other method.
func (r *Reader) Size() int64 { return int64(len(r.s)) }

// ErrReaderEOF is returned when the reader reaches the end of the string.
var ErrReaderEOF = io.EOF

// Read implements the [io.Reader] interface.
func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, ErrReaderEOF
	}
	r.p = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

// ErrReadAtNegativeOffset is returned by [github.com/i9si-sistemas/stringx.Reader.ReadAt] when the offset is negative.
var ErrReadAtNegativeOffset = errors.New("stringx.Reader.ReadAt: negative offset")

// ReadAt implements the [io.ReaderAt] interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, ErrReadAtNegativeOffset
	}
	if off >= int64(len(r.s)) {
		return 0, ErrReaderEOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = ErrReaderEOF
	}
	return
}

// ReadByte implements the [io.ByteReader] interface.
func (r *Reader) ReadByte() (byte, error) {
	r.p = -1
	if r.i >= int64(len(r.s)) {
		return 0, ErrReaderEOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

// ErrUnreadByte is returned by [github.com/i9si-sistemas/stringx.Reader.UnreadByte] when the reader is at the beginning of the string.
var ErrUnreadByte = errors.New("stringx.Reader.UnreadByte: at beginning of string")

// UnreadByte implements the [io.ByteScanner] interface.
func (r *Reader) UnreadByte() error {
	if r.i <= 0 {
		return ErrUnreadByte
	}
	r.p = -1
	r.i--
	return nil
}

// ReadRune implements the [io.RuneReader] interface.
func (r *Reader) ReadRune() (ch rune, size int, err error) {
	if r.i >= int64(len(r.s)) {
		r.p = -1
		return 0, 0, ErrReaderEOF
	}
	r.p = int(r.i)
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRuneInString(r.s[r.i:])
	r.i += int64(size)
	return
}

var (
	// ErrUnreadRuneAtBeginning is returned by [github.com/i9si-sistemas/stringx.Reader.UnreadRune] when the reader is at the beginning of the string.
	ErrUnreadRuneAtBeginning = errors.New("stringx.Reader.UnreadRune: at beginning of string")

	// ErrUnreadRuneNotRead is returned by [github.com/i9si-sistemas/stringx.Reader.UnreadRune] when the previous operation was not ReadRune.
	ErrUnreadRuneNotRead = errors.New("stringx.Reader.UnreadRune: previous operation was not ReadRune")
)

// UnreadRune implements the [io.RuneScanner] interface.
func (r *Reader) UnreadRune() error {
	if r.i <= 0 {
		return ErrUnreadRuneAtBeginning
	}
	if r.p < 0 {
		return ErrUnreadRuneNotRead
	}
	r.i = int64(r.p)
	r.p = -1
	return nil
}

var (
	// ErrSeekInvalidWhence is returned by [github.com/i9si-sistemas/stringx.Reader.Seek] when the whence argument is invalid.
	ErrSeekInvalidWhence = errors.New("stringx.Reader.Seek: invalid whence")

	// ErrSeekNegativePosition is returned by [github.com/i9si-sistemas/stringx.Reader.Seek] when the position is negative.
	ErrSeekNegativePosition = errors.New("stringx.Reader.Seek: negative position")
)

// Seek implements the [io.Seeker] interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	r.p = -1
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.i + offset
	case io.SeekEnd:
		abs = int64(len(r.s)) + offset
	default:
		return 0, ErrSeekInvalidWhence
	}
	if abs < 0 {
		return 0, ErrSeekNegativePosition
	}
	r.i = abs
	return abs, nil
}

// ErrWriteToInvalidWriteStringCount is returned by [github.com/i9si-sistemas/stringx.Reader.WriteTo] when the number of bytes written is not equal to the length of the string.
var ErrWriteToInvalidWriteStringCount = errors.New(
	"stringx.Reader.WriteTo: invalid WriteString count",
)

// WriteTo implements the [io.WriterTo] interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) {
	r.p = -1
	if r.i >= int64(len(r.s)) {
		return 0, nil
	}
	s := r.s[r.i:]
	var writeString func(s string) (int, error)
	switch w := w.(type) {
	case io.StringWriter:
		writeString = func(s string) (int, error) {
			return w.WriteString(s)
		}
	default:
		writeString = func(s string) (int, error) {
			return w.Write([]byte(s))
		}
	}
	m, err := writeString(s)
	if m > len(s) {
		return 0, ErrWriteToInvalidWriteStringCount
	}
	r.i += int64(m)
	n = int64(m)
	if m != len(s) && err == nil {
		err = io.ErrShortWrite
	}
	return
}

// Reset resets the [github.com/i9si-sistemas/stringx.Reader] to be reading from s.
func (r *Reader) Reset(s string) {
	r.s = s
	r.i = 0
	r.p = -1
}
