/*
Problem:
- Given an array containing n distinct numbers taken from the range 0 to n.
  Since the array has only n numbers out of the total n+1 numbers, find the
  missing number.

Example:
- Input: []int{4, 0, 3, 1}
  Output: 2

Approach:
- Sort the array using the cyclic sort first.
- The one that does not have the correct index is the missing one.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{4, 0, 3, 1}, 2},
		{[]int{8, 3, 5, 2, 4, 7, 0, 1}, 6},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			missingNumber(tt.in),
		)
	}
}

func missingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		// if the current number is not at the correct index, swap it with the
		// one that is at the correct index.
		// note that we have to also make sure the current value is not larger
		// than the length of the array because it would causes out of index
		// error.
		correctIndex := nums[i]
		if nums[i] < len(nums) && nums[i] != nums[correctIndex] {
			common.Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}

	// return the missing number that does not have the correct index.
	for j := 0; j < len(nums); j++ {
		if nums[j] != j {
			return j
		}
	}

	return len(nums)
}
