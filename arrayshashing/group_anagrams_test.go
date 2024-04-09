package arrayshashing

import (
	"sort"
	"testing"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func groupAnagrams(strs []string) [][]string {
	anagramGroupLookup := make(map[string][]string, len(strs))
	for _, str := range strs {
		sorted_str := []byte(str)
		sort.Slice(sorted_str, func(i, j int) bool {
			return sorted_str[i] < sorted_str[j]
		})
		key := string(sorted_str)
		anagramGroupLookup[key] = append(anagramGroupLookup[key], str)
	}
	ans := make([][]string, len(anagramGroupLookup))
	index := 0
	for _, values := range anagramGroupLookup {
		ans[index] = values
		index++
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input  []string
		output [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, test := range tests {
		ans := groupAnagrams(test.input)
		if len(test.output) != len(ans) {
			t.Errorf("Expected %v, got %v", len(test.output), len(ans))
		}
	}
}
