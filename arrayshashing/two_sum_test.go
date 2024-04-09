package arrayshashing

import (
	"testing"
)

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	nums_final_position := make(map[int]int, len(nums))
	for pos1, value := range nums {
		pos0, ok := nums_final_position[target-value]
		if ok {
			return []int{pos0, pos1}
		}
		nums_final_position[value] = pos1
	}
	return nil
}

func TestTwoSums(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, test := range tests {
		ans := twoSum(test.nums, test.target)
		if (ans[0] != test.ans[0]) || (ans[1] != test.ans[1]) {
			t.Errorf("Expected %v, got %v for %v, %v", test.ans, ans, test.nums, test.target)
		}
	}
}
