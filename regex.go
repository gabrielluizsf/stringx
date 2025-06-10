package stringx

import "regexp"

func (s String) Includes(substr string) bool {
	return s.Match(regexp.QuoteMeta(substr))
}

func (s String) Match(regex string) bool {
	return s.mustCompile(regex).MatchString(s.String())
}

func (s String) Search(regex string) []string {
	return s.mustCompile(regex).FindAllString(s.String(), -1)
}

func (s String) IndexOf(substr string) int {
	return s.mustCompile(regexp.QuoteMeta(substr)).FindStringIndex(s.String())[0]
}

func (s String) Replace(old, new string) String {
	regexp := s.mustCompile(old)
	return String(regexp.ReplaceAllString(s.String(), new))
}

func (s String) mustCompile(regex string) *regexp.Regexp {
	return regexp.MustCompile(regex)
}

func (s String) Count(value string) int {
	return len(s.Search(value))
}