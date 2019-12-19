/*
Problem:
- Given an array of sorted numbers and a target sum, find a pair in the
  array whose sum is equal to the given target.

Example:
- Input: []int{1, 2, 6, 8, 16, 26}, target=14
  Output: []int{2, 3}
  Explanation: 6 (index 2) + 8 (index 3) = 14

Approach:
- Have one pointer iterate the array and one placing non-duplicate number.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 2}, 2},
		{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, 4},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			removeDuplicates(tt.in),
		)
	}
}

func removeDuplicates(a []int) int {
	if len(a) == 0 {
		return 0
	}

	nonDuplicate := 1

	i := 1
	for i < len(a) {
		// if we see a non-duplicate number, move it next to the last
		// non-duplicate one that we've seen.
		if a[nonDuplicate-1] != a[i] {
			a[nonDuplicate] = a[i]
			nonDuplicate++
		}
		i++
	}

	return nonDuplicate
}
