package arrayshashing

import "testing"

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

func hasDupes(arr []int32) bool {
	// Faster
	values := make(map[int32]bool, len(arr))
	for _, val := range arr {
		if values[val] {
			return true
		} else {
			values[val] = true
		}
	}
	return false
}

func TestArrayDuplicates(t *testing.T) {
	tests := []struct {
		input    []int32
		hasDupes bool
	}{
		{[]int32{1, 3}, false},
		{[]int32{}, false},
		{[]int32{1, 1, 3, 3}, true},
		{[]int32{1, 1, 3, 3}, true},
	}
	for _, test := range tests {
		ans := hasDupes(test.input)
		if ans != test.hasDupes {
			t.Errorf("Expected %v, got %v for %v", test.hasDupes, ans, test.input)
		}
	}

}
