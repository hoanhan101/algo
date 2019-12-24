/*
Problem:
- Given an array containing n numbers taken from the range 1 to n. It can
  have some duplicates. Find all those numbers.

Example:
- Input: []int{5, 4, 7, 2, 3, 5, 3}
  Output: []int{3, 5}

Approach:
- Similar to missing number problem, can rearrange the array using cyclic
  sort.
- Those that do not have the correct indices are the duplicate ones.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{3, 4, 4, 5, 5}, []int{5, 4}},
		{[]int{5, 4, 7, 2, 3, 5, 3}, []int{3, 5}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findDuplicates(tt.in),
		)
	}
}

func findDuplicates(nums []int) []int {
	duplicates := []int{}
	i := 0

	for i < len(nums) {
		correctIndex := nums[i] - 1
		if nums[i] != nums[correctIndex] {
			common.Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}

	// return the duplicates that do not have the correct index.
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			duplicates = append(duplicates, nums[j])
		}
	}

	return duplicates
}
