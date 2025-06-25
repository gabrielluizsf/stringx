package stringx

import (
	"regexp"
)

func (s String) Match(regex string) bool {
	return s.mustCompile(regex).MatchString(s.String())
}

func (s String) Search(regex string) []string {
	return s.mustCompile(regex).FindAllString(s.String(), -1)
}

func (s String) mustCompile(regex string) *regexp.Regexp {
	return regexp.MustCompile(regex)
}
