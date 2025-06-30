package stringx

import (
	"math/rand"
	"time"
)

// Random defines an interface for selecting a random element from a string collection.
type Random interface {
	// Index returns a valid random index based on the collection size.
	Index() int
	// MaxLength returns the maximum size of the string collection.
	MaxLength() int
	// Random returns a random string from the collection.
	Random() String
}

// Random selects a random string from the collection s using the index provided by r.
func (s Strings) Random(r Random) String {
	return s[r.Index()]
}

// random is an implementation of the Random interface using a random index generator function.
type random struct {
	s  Strings           // Collection of strings
	fn func(len int) int // Function that generates a random index given a length
}

// NewRandomString creates a new Random instance from a list of strings.
func NewRandomString(s ...string) Random {
	return random{
		s: ConvertStrings(s...), // Converts values to the Strings type
		fn: func(len int) int {
			rand.New(rand.NewSource(int64(len))).Seed(time.Now().UnixNano())
			return rand.Intn(len)
		}, // Uses the default function for random number generation
	}
}

// Index returns a random index within the bounds of the string collection.
func (r random) Index() int {
	return r.fn(r.MaxLength())
}

// MaxLength returns the length of the string collection.
func (r random) MaxLength() int {
	return len(r.s)
}

// Random returns a random string from the collection.
func (r random) Random() String {
	return r.s.Random(r)
}
