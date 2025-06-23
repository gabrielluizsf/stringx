package stringx

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      String
		value  string
		expect bool
	}{
		{"Equal", String("test"), "test", true},
		{"NotEqual", String("test"), "TEST", false},
		{"EmptyString", Empty, Empty.String(), true},
		{"EmptyValue", String("test"), Empty.String(), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.Equal(tc.value); got != tc.expect {
				t.Errorf("String.Equal() = %v, want %v", got, tc.expect)
			}
		})
	}
}
