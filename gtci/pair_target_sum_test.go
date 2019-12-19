/*
Problem:
- Given an array of sorted numbers and a target sum, find a pair in the
  array whose sum is equal to the given target.

Example:
- Input: []int{1, 2, 6, 8, 16, 26}, target=14
  Output: []int{2, 3}
  Explanation: 6 (index 2) + 8 (index 3) = 14

Approach:
- Have one pointer start at the beginning and one at the end of the array.
- At each step, see if the two pointers add up to the target sum and move
  them toward each other accordingly.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindPairTargetSum(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []int
	}{
		{[]int{}, 6, []int{}},
		{[]int{1, 2, 6, 8, 16, 26}, 14, []int{2, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 6, []int{0, 4}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findPairTargetSum(tt.in1, tt.in2),
		)
	}
}

func findPairTargetSum(a []int, target int) []int {
	out := []int{}

	start := 0
	end := len(a) - 1

	for start < end {
		sum := a[start] + a[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			out = append(out, start, end)
			break
		}
	}

	return out
}
