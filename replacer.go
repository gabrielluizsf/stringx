package stringx

import (
	"errors"
)

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

// Replace performs the replacement of old strings with new strings in the String.
func (s String) Replace(old, new string) String {
	buf := &Builder{}
	finder := makeStringFinder(old)
	i, matched := 0, false
	for {
		match := finder.next(s.String()[i:])
		if match == -1 {
			break
		}
		matched = true
		buf.Grow(match + len(new))
		buf.WriteString(s.String()[i : i+match])
		buf.WriteString(new)
		i += match + len(finder.pattern)
	}
	if !matched {
		return s
	}
	buf.WriteString(s.String()[i:])
	return Convert(buf)
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

	original := r.s
	for i, old := range r.old {
		original = String(original).Replace(old, r.new[i])
	}
	return original.String(), nil
}
