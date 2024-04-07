package arrayshashing

import (
	"testing"
)

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func areAnagrams(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_values := make(map[rune]int, len(s))
	for _, char := range s {
		s_values[char] += 1
	}
	t_values := make(map[rune]int, len(t))
	for _, char := range t {
		t_values[char] += 1
	}

	if len(s_values) != len(t_values) {
		return false
	}
	for char, n_occurences := range s_values {
		if t_values[char] != n_occurences {
			return false
		}
	}
	return true
}

func TestIsAna(t *testing.T) {
	tests := []struct {
		s     string
		t     string
		isAna bool
	}{
		{"", "", true},
		{"ast", "sat", true},
		{"asdf", "ssdkfat", false},
		{"asdf", "ssdf", false},
	}
	for _, test := range tests {
		ans := areAnagrams(test.s, test.t)
		if ans != test.isAna {
			t.Errorf("Expected %v, got %v for %v, %v", test.isAna, ans, test.s, test.t)
		}
	}

}
