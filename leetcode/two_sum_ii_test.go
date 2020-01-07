/*
Problem:
- Given a sorted array of integers, return indices of the two numbers such
  that they add up to a specific target.

Example:
- Input: nums = []int{2, 3, 4}, target = 6
  Output: [0, 2] since nums[0] + nums[2] = 2 + 4 = 6

Approach
- Since the array is sorted, can use two-pointer approach that has one point
  to the start of the list while the other point at the end and move the
  toward each other.

Solution:
- Have one pointer point to the start of the list and the other point at the end.
- Iterate through the list and check their sum.
- If it is smaller than the target, increment the start pointer to make the
  sum bigger little by little.
- Otherwise, decrement the end pointer to make the sum smaller.
- When the sum is equal to the target, we found the answer.

Cost:
- O(n) time, O(1) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []int
	}{
		{[]int{}, 6, []int{0, 0}},
		{[]int{1}, 6, []int{0, 0}},
		{[]int{2, 4}, 6, []int{0, 1}},
		{[]int{2, 5}, 6, []int{0, 0}},
		{[]int{2, 3, 4}, 6, []int{0, 2}},
		{[]int{2, 5, 8}, 6, []int{0, 0}},
		{[]int{2, 3, 4, 10}, 6, []int{0, 2}},
	}

	for _, tt := range tests {
		result := twoSumII(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func twoSumII(nums []int, target int) []int {
	// start and end points to the start and end index of the list.
	start := 0
	end := len(nums) - 1
	out := make([]int, 2)

	for start < end {
		sum := nums[start] + nums[end]
		if sum < target {
			start++
		} else if sum > target {
			end--
		} else {
			out[0] = start
			out[1] = end
			break
		}
	}

	return out
}
