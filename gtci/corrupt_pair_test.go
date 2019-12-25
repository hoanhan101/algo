/*
Problem:
- Given an array containing n+1 numbers taken from the range 1 to n. One
  of the numbers got duplicated which also resulted in one number going
  missing. Find these numbers.

Example:
- Input: []int{3, 1, 2, 5, 2}
  Output: []int{2, 4}

Approach:
- Similar to finding duplicates problem, can place each number on its correct
  index.
- The one is not at its correct index is the duplicate and its index itself
  is the missing number.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindCorruptPair(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{0, 0}},
		{[]int{3, 1, 2, 5, 2}, []int{2, 4}},
		{[]int{3, 1, 2, 3, 6, 4}, []int{3, 5}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findCorruptPair(tt.in),
		)
	}
}

func findCorruptPair(nums []int) []int {
	i := 0

	for i < len(nums) {
		correctIndex := nums[i] - 1
		if nums[i] != nums[correctIndex] {
			common.Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}

	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			return []int{nums[j], j + 1}
		}
	}

	return []int{0, 0}
}
