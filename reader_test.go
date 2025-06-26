package stringx

import (
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/i9si-sistemas/assert"
)

var generateUnderscoreString = func(max int) string { return Empty.Repeat("_", max).String() }

func TestReader(t *testing.T) {
	s := generateUnderscoreString(10)
	r := NewReader(s)
	tests := []struct {
		off       int64
		seek      int
		n         int
		want      string
		wantPos   int64
		readerErr error
		seekErr   string
	}{
		{seek: io.SeekStart, off: 0, n: 10, want: s},
		{seek: io.SeekStart, off: 1, n: 1, want: string(s[1])},
		{seek: io.SeekCurrent, off: 1, wantPos: 3, n: 2, want: s[2:4]},
		{seek: io.SeekStart, off: -1, seekErr: ErrSeekNegativePosition.Error()},
		{seek: io.SeekStart, off: 1 << 33, wantPos: 1 << 33, readerErr: ErrReaderEOF},
		{seek: io.SeekCurrent, off: 1, wantPos: 1<<33 + 1, readerErr: ErrReaderEOF},
		{seek: io.SeekStart, n: 5, want: s[:5]},
		{seek: io.SeekCurrent, n: 5, want: s[5:10]},
		{seek: io.SeekEnd, off: -1, n: 1, wantPos: 9, want: s[len(s)-1:]},
	}

	for index, testCase := range tests {
		pos, err := r.Seek(testCase.off, testCase.seek)
		testCaseIndex := index + 1
		if err == nil && !Empty.Equal(testCase.seekErr) {
			t.Errorf("%d. want seek error %q", testCaseIndex, testCase.seekErr)
			continue
		}
		if err != nil && !String(err.Error()).Equal(testCase.seekErr) {
			t.Errorf("%d. seek error = %q; want %q", testCaseIndex, err.Error(), testCase.seekErr)
			continue
		}
		assert.False(t, testCase.wantPos != 0 && testCase.wantPos != pos)
		buf := make([]byte, testCase.n)
		n, err := r.Read(buf)
		if err != testCase.readerErr {
			t.Errorf("%d. read = %v; want %v", testCaseIndex, err, testCase.readerErr)
			continue
		}
		got := string(buf[:n])
		assert.Equal(t, got, testCase.want, fmt.Sprintf("test case %d", testCaseIndex))
	}
}

func TestReadAfterBigSeek(t *testing.T) {
	r := NewReader(generateUnderscoreString(100))
	if _, err := r.Seek(1<<31+5, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	if n, err := r.Read(make([]byte, 100)); n != 0 || err != ErrReaderEOF {
		t.Errorf("Read = %d, %v; want 0, EOF", n, err)
	}
}

func TestReaderAt(t *testing.T) {
	s := generateUnderscoreString(10)
	r := NewReader(s)
	tests := []struct {
		off     int64
		n       int
		want    string
		wantErr any
	}{
		{0, 10, s, nil},
		{1, 10, s[1:], ErrReaderEOF},
		{1, 9, s[1:], nil},
		{11, 10, Empty.String(), ErrReaderEOF},
		{0, 0, Empty.String(), nil},
		{-1, 0, Empty.String(), ErrReadAtNegativeOffset},
	}
	for index, testCase := range tests {
		b := make([]byte, testCase.n)
		rn, err := r.ReadAt(b, testCase.off)
		got := string(b[:rn])
		assert.Equal(t, got, testCase.want, fmt.Sprintf("test case %d", index+1))
		assert.StrictEqual(t, err, testCase.wantErr)
	}
	t.Run("race detector", func(t *testing.T) {
		r := NewReader(generateUnderscoreString(10))
		var wg sync.WaitGroup
		for i := range 5 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				var buf [1]byte
				r.ReadAt(buf[:], int64(i))
			}(i)
		}
		wg.Wait()
	})
	t.Run("race detector with empty string", func(t *testing.T) {
		r := NewReader(Empty.String())
		var wg sync.WaitGroup
		for range 5 {
			wg.Add(2)
			go func() {
				defer wg.Done()
				var buf [1]byte
				r.Read(buf[:])
			}()
			go func() {
				defer wg.Done()
				r.Read(nil)
			}()
		}
		wg.Wait()
	})
}

func TestWriteTo(t *testing.T) {
	const str = "1_2_3_4_5_6_7_8_9"
	for i := range len(str) {
		s := str[i:]
		r := NewReader(s)
		builder := NewBuilder()
		n, err := r.WriteTo(builder)
		expect := int64(len(s))
		assert.Nil(t, err)
		assert.Equal(t, n, expect)
		assert.True(t, Convert(builder).Equal(s))
		assert.Zero(t, r.Len())
	}
}

func TestReaderReset(t *testing.T) {
	r := NewReader("世界")
	if _, _, err := r.ReadRune(); err != nil {
		t.Errorf("ReadRune: unexpected error: %v", err)
	}

	const want = "Hello, 世界"
	r.Reset(want)
	if err := r.UnreadRune(); err == nil {
		t.Errorf("UnreadRune: expected error, got nil")
	}
	b, err := io.ReadAll(r)
	assert.Nil(t, err)
	assert.Equal(t, string(b), want)
}
