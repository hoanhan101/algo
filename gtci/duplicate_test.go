/*
Problem:
- Given an array containing n+1 numbers taken from the range 1 to n. It has
  only one duplicate number but can be repeated over time. Find that one.

Example:
- Input: []int{1, 4, 4, 3, 2}
  Output: 4

Approach:
- Similar to missing number problem, can place each number on its correct
  index.
- If while swapping the number with its index both the numbers being swapped
  are same, we have found the duplicate.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 4, 4, 3, 2}, 4},
		{[]int{2, 4, 1, 4, 4}, 4},
		{[]int{2, 1, 3, 3, 5, 4}, 3},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findDuplicate(tt.in),
		)
	}
}

func findDuplicate(nums []int) int {
	i := 0

	for i < len(nums) {
		if nums[i] != i+1 {
			correctIndex := nums[i] - 1
			if nums[i] != nums[correctIndex] {
				common.Swap(nums, i, correctIndex)
			} else {
				return nums[i]
			}
		} else {
			i++
		}
	}

	return 0
}
