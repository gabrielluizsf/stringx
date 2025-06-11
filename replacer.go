package stringx

import "errors"

// Replacer is a struct that holds a String and two slices of strings for replacement.
type Replacer struct {
	s   String
	old []string
	new []string
}

// NewReplacer creates a new Replacer instance with the provided String and slices of old and new strings.
func NewReplacer(s String, old []string, new []string) *Replacer {
	return &Replacer{s: s, old: old, new: new}
}

var (
	// ErrOldOrNewCannotBeEmpty is returned when either old or new slice is empty.
	ErrOldOrNewCannotBeEmpty = errors.New("stringx: old or new cannot be empty")
	// ErrOldAndNewMustHaveSameLen is returned when old and new slice have different length.
	ErrOldAndNewMustHaveSameLen = errors.New("stringx: old and new must have the same length")
)

// Replace performs the replacement of old strings with new strings in the String.
func (r *Replacer) Replace() (string, error) {
	if len(r.old) != len(r.new) {
		return r.s.String(), ErrOldAndNewMustHaveSameLen
	}
	if len(r.old) == 0 || len(r.new) == 0 {
		return r.s.String(), ErrOldOrNewCannotBeEmpty
	}
	for i, old := range r.old {
		r.s = r.s.Replace(old, r.new[i])
	}
	return r.s.String(), nil
}
