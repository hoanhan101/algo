/*
Problem:
- Given an array containing 0s, 1s and 2s, sort the array in-place.

Example:
- Input: []int{1, 0, 2, 1, 0}
  Output: []int{0, 0, 1, 1, 2}

Approach:
- Have one pointer start at the beginning and the other at the end
  while iterating through the array.
- We will move all 0s before that start pointer and 2s after the end
  pointer so that all 1s would be between in the end.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestDutchFlagSort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{[]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}},
		{[]int{1, 0, 2, 1, 0}, []int{0, 0, 1, 1, 2}},
	}

	for _, tt := range tests {
		dutchFlagSort(tt.in)
		common.Equal(
			t,
			tt.expected,
			tt.in,
		)
	}
}

func dutchFlagSort(a []int) {
	start := 0
	end := len(a) - 1

	i := 0
	for i <= end {
		if a[i] == 0 {
			// swap a[i] and a[start].
			common.Swap(a, i, start)

			// increment start pointer and i.
			i++
			start++
		} else if a[i] == 1 {
			// increment i pointer only.
			i++
		} else if a[i] == 2 {
			// swap a[i] and a[end].
			common.Swap(a, i, end)

			// decrement end pointer only.
			end--
		}
	}
}
