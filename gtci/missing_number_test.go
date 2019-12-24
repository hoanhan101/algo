/*
Problem:
- Given an array containing n numbers taken from the range 1 to n. It can
  have duplicates. Find all those missing numbers.

Example:
- Input: []int{2, 3, 1, 8, 2, 3, 5, 1}
  Output: []int{4, 6, 7}

Approach:
- Similar to missing number problem, can rearrange the array using cyclic
  sort.
- Those that do not have the correct indices are the missing ones.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMissingNumbers(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 4, 1, 2}, []int{3}},
		{[]int{2, 3, 2, 1}, []int{4}},
		{[]int{2, 3, 1, 8, 2, 3, 5, 1}, []int{4, 6, 7}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			missingNumbers(tt.in),
		)
	}
}

func missingNumbers(nums []int) []int {
	missed := []int{}
	i := 0

	for i < len(nums) {
		correctIndex := nums[i] - 1
		if nums[i] != nums[correctIndex] {
			common.Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}

	// return the missing number that does not have the correct index.
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			missed = append(missed, j+1)
		}
	}

	return missed
}
